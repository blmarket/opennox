#!/usr/bin/env bash
#
# Drive the coverage loop (COVERAGE_LOOP.md) for N headless iterations.
# Each iteration is a fresh, non-interactive Claude Code run that adds tests for
# a batch of EASY decompiled C functions, verifies with ./t1, and commits.
#
# Usage:
#   scripts/coverage_loop_run.sh [N]      # default N=10
#
# Env overrides:
#   CLAUDE_BIN    claude executable (default: claude)
#   CLAUDE_FLAGS  extra flags (default: --dangerously-skip-permissions)
#   STALL_LIMIT   stop after this many consecutive no-progress iters (default: 2)
#
# Logs: /tmp/coverage-loop-logs/iter-<n>.log
set -uo pipefail

REPO_SRC="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_SRC"

N="${1:-10}"
CLAUDE_BIN="${CLAUDE_BIN:-metacode}"
CLAUDE_FLAGS="${CLAUDE_FLAGS:---yolo}"
STALL_LIMIT="${STALL_LIMIT:-2}"
LOGDIR="/tmp/coverage-loop-logs"
REPORT="/tmp/opennox-ccover.json"
mkdir -p "$LOGDIR"

PROMPT='Perform exactly ONE iteration of the coverage loop described in ./COVERAGE_LOOP.md (read that file and follow it precisely). Work ONLY on EASY (pure) functions selected via scripts/coverage_triage.py --tier easy. Add the test files, ensure ./t1 is green and function coverage strictly increased, then commit per the file. If you cannot make green progress, revert your changes and stop. Do not work on more than one source file.'

if ! command -v "$CLAUDE_BIN" >/dev/null 2>&1; then
	echo "error: '$CLAUDE_BIN' not found on PATH (set CLAUDE_BIN)" >&2
	exit 1
fi

funcs() {
	python3 -c "import json;print(json.load(open('$REPORT'))['functions']['covered'])" 2>/dev/null || echo 0
}

# Baseline coverage.
if [[ ! -f "$REPORT" ]]; then
	echo "[runner] no report yet, generating baseline with ./t1 ..."
	./t1 >/dev/null 2>&1
fi
prev="$(funcs)"
start="$prev"
start_ref="$(git rev-parse HEAD 2>/dev/null || echo '')"
echo "[runner] baseline functions covered: $prev"

stall=0
for ((i = 1; i <= N; i++)); do
	echo "============================================================"
	echo "[runner] iteration $i/$N  (functions covered: $prev)  $(date)"
	echo "============================================================"

	"$CLAUDE_BIN" $CLAUDE_FLAGS run "$PROMPT" 2>&1 | tee "$LOGDIR/iter-$i.log"

	cur="$(funcs)"
	echo "[runner] iteration $i result: functions $prev -> $cur"

	if [[ "$cur" -le "$prev" ]]; then
		stall=$((stall + 1))
		echo "[runner] no progress (stall $stall/$STALL_LIMIT)"
		if [[ "$stall" -ge "$STALL_LIMIT" ]]; then
			echo "[runner] stopping early: $STALL_LIMIT consecutive no-progress iterations"
			break
		fi
	else
		stall=0
	fi
	prev="$cur"
done

echo "============================================================"
echo "[runner] done. functions covered: $start -> $(funcs)  (+$(($(funcs) - start)))"
echo "[runner] per-iteration logs in $LOGDIR/"
echo "[runner] commits this run:"
if [[ -n "$start_ref" ]]; then
	git --no-pager log --oneline "${start_ref}..HEAD" 2>/dev/null || true
fi
