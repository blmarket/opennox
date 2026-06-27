#!/usr/bin/env bash
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
src_root="$repo_root/src"
minimum=${CCOVER_MIN:-100}
report=${1:-"$repo_root/legacy-c-coverage.json"}

if [[ -z ${NOX_DATA:-} ]]; then
	echo "coverage-legacy-c: NOX_DATA must point at a provisioned Nox data directory" >&2
	exit 2
fi
if [[ ! -d $NOX_DATA ]]; then
	echo "coverage-legacy-c: NOX_DATA is not a directory: $NOX_DATA" >&2
	exit 2
fi

tmp=$(mktemp -d "${TMPDIR:-/tmp}/opennox-ccover.XXXXXX")
cleanup() {
	if [[ ${CCOVER_KEEP_WORK:-0} == 1 ]]; then
		echo "coverage-legacy-c: retained build data at $tmp" >&2
	else
		rm -rf "$tmp"
	fi
}
trap cleanup EXIT
mkdir -p "$tmp/work" "$tmp/cache"

cd "$src_root"
env \
	GOARCH=386 \
	CGO_ENABLED=1 \
	GOTMPDIR="$tmp/work" \
	GOCACHE="$tmp/cache" \
	CGO_CFLAGS='-O0 -g --coverage -fprofile-abs-path' \
	CGO_LDFLAGS='--coverage' \
	'CGO_CFLAGS_ALLOW=(-fshort-wchar)|(-fno-strict-aliasing)|(-fno-strict-overflow)' \
	go test -tags=ccover -count=1 -work ./legacy/... 2>&1 | tee "$tmp/test.log"

test_work=$(sed -n 's/^WORK=//p' "$tmp/test.log" | head -n 1)
if [[ -z $test_work ]]; then
	echo "coverage-legacy-c: go test did not report its work directory" >&2
	exit 1
fi

env \
	GOARCH=386 \
	CGO_ENABLED=1 \
	GOTMPDIR="$tmp/work" \
	GOCACHE="$tmp/cache" \
	CGO_CFLAGS='-O0 -g --coverage -fprofile-abs-path' \
	CGO_LDFLAGS='--coverage' \
	'CGO_CFLAGS_ALLOW=(-fshort-wchar)|(-fno-strict-aliasing)|(-fno-strict-overflow)' \
	go test -tags='ccover safe' -count=1 \
	-run 'TestLegacySettingsDisconnected|TestNoxThing(Skip|ReadImage|ReadAbility)' \
	-work ./legacy \
	> "$tmp/test-safe.log" 2>&1

if [[ ${CCOVER_E2E:-1} == 1 ]]; then
	env \
		GOARCH=386 \
		CGO_ENABLED=1 \
		GOTMPDIR="$tmp/work" \
		GOCACHE="$tmp/cache" \
		CGO_CFLAGS='-O0 -g --coverage -fprofile-abs-path' \
		CGO_LDFLAGS='--coverage' \
		'CGO_CFLAGS_ALLOW=(-fshort-wchar)|(-fno-strict-aliasing)|(-fno-strict-overflow)' \
		go build -tags=ccover -work -o "$tmp/opennox" ./cmd/opennox > "$tmp/build.log" 2>&1

	maps=(BankShot BluDeath Bunker CapFlag EndGame Estate FlagBall FlagWar FortNox Fortress FreezOut G_CastlD G_Castle G_CryptD G_Crypts G_ForesD G_Forest G_LOTD G_LOTDD G_Lava G_Mines G_Swamp G_TemplD G_Temple Inferno Kingdoms Library LostTomb ManaMine MiniMine MnaVault Oasis So_Beach So_Brin So_Druid So_Dun So_FOV So_Galav So_Grok So_Ix So_LOD So_Mines So_Mount So_Open So_Swamp So_Waste So_Woods SpyFort TreeHaus TriLevel con01a con02a con03a con03b con04a con04b con04c con05a con05b con05c con06a con06b con07a con07b con07c con07d con07e con07f con07g con07h con08a con08b con08c con08d con08e con09a con09b con09c con09d con10a con10b con10c con10d con11a war01a war02a war02b war03a war03b war03c war03d war04a war04b war04c war05a war05b war05c war06a war06b war07a war07b war07c war07d war07e war07f war07g war07h war08a war08b war08c war08d war08e war09a war09b war09c war09d war10a war10b war10c war10d war11a wiz01a wiz02a wiz02b wiz02c wiz03a wiz03b wiz03c wiz04a wiz04b wiz04c wiz05a wiz05b wiz05c wiz06a wiz06b wiz06c wiz07a wiz07b wiz07c wiz07d wiz07e wiz07f wiz08a wiz08b wiz08c wiz08d wiz08e wiz09a wiz09b wiz09c wiz09d wiz10a wiz10b wiz10c wiz10d wiz11a)
	concurrency=12
	count=0
	for map in "${maps[@]}"; do
		(
			port=$((18590 + count))
			clientport=$((19590 + count))
			pprof=$((6060 + count))
			if ! env \
				NOX_E2E="$src_root/testdata/legacy-c-coverage/e2e.yaml" \
				xvfb-run -a "$tmp/opennox" \
				-noaudio \
				-autosrv \
				-port $port \
				-clientport $clientport \
				-pprof "127.0.0.1:$pprof" \
				-autoexec "set maps allow.all true; load $map" \
				-config "$tmp/opennox-$map.yml" > "$tmp/e2e-$map.log" 2>&1; then
				echo "Warning: E2E map failed: $map" >&2
				tail -n 80 "$tmp/e2e-$map.log" >&2
			fi
		) &
		count=$((count + 1))
		if [[ $((count % concurrency)) -eq 0 ]]; then
			wait
		fi
	done
	wait
fi

go run ./internal/ccoverreport \
	-root "$src_root/legacy" \
	-work "$tmp/work" \
	-out "$report" \
	-fail-under "$minimum"
