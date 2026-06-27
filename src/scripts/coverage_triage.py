#!/usr/bin/env python3
"""Rank decompiled legacy C functions by how easy they are to unit-test.

Used by the coverage loop (see COVERAGE_LOOP.md) to pick the next batch of
targets in GAME*.c (and other legacy/*.c). It:

  * parses banner-delimited functions out of the .c files,
  * classifies each easy / medium / hard by a token + size heuristic,
  * skips functions that already have a C_* wrapper in a *_test_helper.go,
  * (optionally) skips functions whose definition line is already covered,
  * prints a ranked, file-grouped worklist.

Usage:
  scripts/coverage_triage.py                      # all easy, unwrapped candidates
  scripts/coverage_triage.py --file GAME5.c       # focus one file
  scripts/coverage_triage.py --tier easy --max 15 # cap output
  scripts/coverage_triage.py --json               # machine-readable

Run from the repo's src/ directory (where legacy/ lives).
"""
import argparse
import glob
import json
import os
import re
import sys

BANNER = re.compile(r"^//----- \(([0-9A-Fa-f]{8})\)")
SIG = re.compile(r"^[A-Za-z_].*\b(?P<name>(sub_|nox_|nullsub_)[A-Za-z0-9_]+)\s*\([^;{]*\)\s*\{")

# Tokens that make a function genuinely hard/unpredictable to unit-test:
# external I/O, allocation, GUI/windows, server singletons. (Prefix matches —
# no trailing \b, since real tokens look like nox_fs_open, AIL_sample_*, etc.)
HARD = re.compile(
    r"\bGetServer\b|\bFILE\b|\bnox_fs_|\bcalloc\b|\bmalloc\b|\bfree\s*\(|"
    r"\bfopen\b|\bsocket\b|\brecv\b|\bsend\b|\bnox_window|\bnox_xxx_wnd|"
    r"\bAIL_|\bnox_swprintf\b|\bnox_sprintf\b|\bnox_alloc|\bnox_new_window"
)
# Memory-blob / decompiled-global access => "medium": deterministic once you
# snapshot/seed the blob (0x5D4594 game state, 0x581450 / 0x587000 data).
MEM = re.compile(r"\bgetMem|\b(dword|qword|byte|word|flt|off)_[0-9a-fA-F]")


def legacy_dir():
    for cand in ("legacy", "."):
        if os.path.isdir(os.path.join(cand, "common", "ccall")) or glob.glob(
            os.path.join(cand, "GAME*.c")
        ):
            return cand
    return "legacy"


def wrapped_names(ld):
    """Names already exposed via a C_* wrapper in any *_test_helper.go."""
    names = set()
    for f in glob.glob(os.path.join(ld, "*_test_helper.go")):
        txt = open(f, encoding="utf-8", errors="replace").read()
        for m in re.finditer(r"\bC\.([A-Za-z0-9_]+)\(", txt):
            names.add(m.group(1))
    return names


def covered_lines(report_path):
    """file -> set(uncovered line numbers); used to skip covered funcs."""
    out = {}
    if not report_path or not os.path.exists(report_path):
        return out
    rep = json.load(open(report_path))
    for f in rep.get("files", []):
        out[f["file"]] = set(f.get("uncovered_lines", []))
    return out


def parse_functions(path):
    lines = open(path, encoding="utf-8", errors="replace").read().split("\n")
    idx = [i for i, l in enumerate(lines) if BANNER.match(l)]
    idx.append(len(lines))
    funcs = []
    for a, b in zip(idx, idx[1:]):
        block = lines[a:b]
        sig = name = None
        defline = None
        for j, l in enumerate(block):
            m = SIG.match(l.strip())
            if m:
                sig, name, defline = l.strip(), m.group("name"), a + j + 1
                break
        if not name:
            continue
        body = "\n".join(block)
        code = [l for l in block if l.strip() and not l.strip().startswith("//")]
        funcs.append(
            {"name": name, "sig": sig, "line": defline, "size": len(code), "body": body}
        )
    return funcs


def classify(fn):
    if HARD.search(fn["body"]):
        return "hard"
    if MEM.search(fn["body"]):
        return "medium"
    return "easy"


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--file", help="restrict to one .c file (e.g. GAME5.c)")
    ap.add_argument("--tier", default="easy", choices=["easy", "medium", "hard", "all"])
    ap.add_argument("--max", type=int, default=40)
    ap.add_argument("--max-size", type=int, default=14, help="max body lines for easy")
    ap.add_argument("--report", default="/tmp/opennox-ccover.json")
    ap.add_argument("--json", action="store_true")
    args = ap.parse_args()

    ld = legacy_dir()
    pat = args.file if args.file else "GAME*.c"
    files = sorted(glob.glob(os.path.join(ld, pat)))
    if not files:
        print(f"no source files matched {pat!r} under {ld}/", file=sys.stderr)
        return 2

    wrapped = wrapped_names(ld)
    uncovered = covered_lines(args.report)

    results = []
    for path in files:
        base = os.path.basename(path)
        unc = uncovered.get(base)
        for fn in parse_functions(path):
            if fn["name"] in wrapped:
                continue
            tier = classify(fn)
            if args.tier != "all" and tier != args.tier:
                continue
            if tier == "easy" and fn["size"] > args.max_size:
                continue
            # if we have a report, prefer functions known to be uncovered
            covered = unc is not None and fn["line"] not in unc
            if covered:
                continue
            results.append(
                {
                    "file": base,
                    "name": fn["name"],
                    "line": fn["line"],
                    "size": fn["size"],
                    "tier": tier,
                    "sig": fn["sig"][:110],
                }
            )

    # rank: smaller body first (cheaper, more predictable), then by file
    results.sort(key=lambda r: (r["size"], r["file"], r["line"]))
    results = results[: args.max]

    if args.json:
        print(json.dumps(results, indent=2))
        return 0

    if not results:
        print("no candidates (everything matching is already wrapped/covered)")
        return 0

    cur = None
    for r in results:
        if r["file"] != cur:
            cur = r["file"]
            print(f"\n===== {cur} =====")
        print(f"  L{r['line']:<6} [{r['tier']:6}] ({r['size']:2}) {r['sig']}")
    print(f"\n{len(results)} candidate(s). Pick from ONE file, implement, verify.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
