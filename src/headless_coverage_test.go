package opennox

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

const headlessCoverageChild = "NOX_HEADLESS_COVERAGE_CHILD"
const headlessClientCoverageChild = "NOX_HEADLESS_CLIENT_COVERAGE_CHILD"

func headlessCoverageData() string {
	if path := os.Getenv("NOX_DATA"); path != "" {
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	path := filepath.Join(home, "proj", "NOX")
	if _, err := os.Stat(filepath.Join(path, "gamedata.bin")); err != nil {
		return ""
	}
	return path
}

func headlessCoverageScenario(t *testing.T, name string) string {
	t.Helper()
	path, err := filepath.Abs(filepath.Join("testdata", "legacy-c-coverage", name))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
	return path
}

// TestZZHeadlessClientCoverage runs the existing deterministic client/server
// script against an in-memory seat. It exercises rendering, GUI, networking,
// and client map loading without creating an SDL or native window.
func TestZZHeadlessClientCoverage(t *testing.T) {
	data := headlessCoverageData()
	if data == "" {
		t.Skip("set NOX_DATA to run the provisioned headless client coverage scenario")
	}
	if os.Getenv(headlessClientCoverageChild) == "1" {
		configReadOnly = true
		noxMapsIgnoreMode = true
		args := []string{
			"opennox",
			"-data", data,
			"-headless",
			"-noaudio",
			"-nolimit",
			"-port", strconv.Itoa(40000 + os.Getpid()%20000),
		}
		if !strings.Contains(filepath.Base(e2ePlay), "menu") {
			args = append(args, "-autosrv")
		}
		err := RunArgs(args)
		if err != nil {
			t.Fatal(err)
		}
		return
	}

	tests := []struct {
		scenario string
		class    string
	}{
		{scenario: "e2e-headless-client.yaml"},
		{scenario: "e2e-headless-combat.yaml"},
		{scenario: "e2e-headless-menu.yaml"},
		{scenario: "e2e-headless-menu-solo.yaml", class: "warrior"},
		{scenario: "e2e-headless-menu-solo.yaml", class: "conjurer"},
		{scenario: "e2e-headless-menu-solo.yaml", class: "wizard"},
		{scenario: "e2e-headless-menu-quest.yaml"},
	}
	for _, tc := range tests {
		name := strings.TrimSuffix(tc.scenario, ".yaml")
		if tc.class != "" {
			name += "-" + tc.class
		}
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestZZHeadlessClientCoverage$", "-test.v")
			cmd.Env = append(os.Environ(),
				headlessClientCoverageChild+"=1",
				"NOX_DATA="+data,
				"NOX_E2E="+headlessCoverageScenario(t, tc.scenario),
				"NOX_E2E_CLASS="+tc.class,
			)
			out, err := cmd.CombinedOutput()
			if ctx.Err() != nil {
				t.Fatalf("headless client coverage scenario timed out: %v\n%s", ctx.Err(), out)
			}
			if err != nil {
				t.Fatalf("headless client coverage scenario failed: %v\n%s", err, out)
			}
		})
	}
}

// TestZZHeadlessServerCoverage runs last and uses a subprocess because the
// legacy application lifecycle owns process-global state. -serveronly skips
// SDL seat creation, and -noDraw keeps the client renderer disabled while the
// server loads and advances several representative maps.
func TestZZHeadlessServerCoverage(t *testing.T) {
	data := headlessCoverageData()
	if data == "" {
		t.Skip("set NOX_DATA to run the provisioned headless server coverage scenario")
	}
	if os.Getenv(headlessCoverageChild) == "1" {
		runHeadlessCoverageChild(t, data)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestZZHeadlessServerCoverage$", "-test.v")
	cmd.Env = append(os.Environ(), headlessCoverageChild+"=1", "NOX_DATA="+data)
	out, err := cmd.CombinedOutput()
	if ctx.Err() != nil {
		t.Fatalf("headless server coverage scenario timed out: %v\n%s", ctx.Err(), out)
	}
	if err != nil {
		t.Fatalf("headless server coverage scenario failed: %v\n%s", err, out)
	}
}

func runHeadlessCoverageChild(t *testing.T, data string) {
	configReadOnly = true
	noxMapsIgnoreMode = true
	type mapRun struct {
		name   string
		frames int
	}
	maps := []mapRun{
		{"Estate", 120},  // arena
		{"CapFlag", 60},  // capture the flag
		{"FlagBall", 60}, // flagball
		{"Kingdoms", 60}, // king of the realm
		{"FreezOut", 60}, // elimination
		{"Bunker", 30},   // dense multiplayer objects
		{"ManaMine", 30}, // multiplayer objects and elevators
		{"LostTomb", 30}, // multiplayer traps and doors
		{"So_Dun", 60},   // solo/cooperative objects
		{"So_Beach", 30},
		{"So_FOV", 30},
		{"So_Grok", 30},
		{"So_Mines", 30},
		{"So_Swamp", 30},
		{"war01a", 60}, // warrior campaign object families
		{"war03d", 30},
		{"war05c", 30},
		{"war07h", 30},
		{"war10d", 30},
		{"con01a", 60}, // conjurer campaign object families
		{"con03b", 30},
		{"con05c", 30},
		{"con07h", 30},
		{"con10d", 30},
		{"G_Castle", 30}, // quest themes and monster sets
		{"G_Crypts", 30},
		{"G_Forest", 30},
		{"G_Lava", 30},
		{"G_Mines", 30},
		{"G_Swamp", 30},
		{"G_Temple", 30},
		{"G_LOTDD", 30},
	}
	seen := make(map[string]bool, len(maps))
	for _, run := range maps {
		seen[strings.ToLower(run.name)] = true
	}
	// These king-of-the-realm maps dereference a missing crown when loaded in
	// this server-only lifecycle. Keep the coverage scenario headless instead
	// of starting an interactive client window for them.
	unsafeHeadless := map[string]bool{
		"inferno":  true,
		"library":  true,
		"minimine": true,
		"mnavault": true,
		"oasis":    true,
	}
	entries, err := os.ReadDir(filepath.Join(data, "maps"))
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		name := strings.ToLower(entry.Name())
		if !entry.IsDir() || seen[name] || unsafeHeadless[name] {
			continue
		}
		maps = append(maps, mapRun{name: entry.Name(), frames: 5})
	}

	target := 0
	frames := 0
	var scenarioErr error
	stop := func() {
		cleanup()
		nox_exit(0)
	}
	mainloopHook = func() {
		if noxServer == nil || !noxflags.HasGame(noxflags.GameFlag29) {
			return
		}
		loaded := strings.TrimSuffix(noxServer.nox_server_currentMapGetFilename_409B30(), ".map")
		if !strings.EqualFold(loaded, maps[target].name) {
			return
		}
		frames++
		if frames < maps[target].frames {
			return
		}
		if target == len(maps)-1 {
			stop()
			return
		}
		target++
		frames = 0
		if err := noxLoadMap(maps[target].name, mapLoadOptions{Force: true}); err != nil {
			scenarioErr = fmt.Errorf("load %s: %w", maps[target].name, err)
			stop()
		}
	}
	t.Cleanup(func() { mainloopHook = nil })

	err = RunArgs([]string{
		"opennox",
		"-data", data,
		"-config", filepath.Join(t.TempDir(), "opennox.yml"),
		"-serveronly",
		"-noDraw",
		"-noaudio",
		"-nolimit",
		"-autosrv",
		"-port", strconv.Itoa(40000 + os.Getpid()%20000),
		"-autoexec", "load " + maps[0].name,
		"-logs", t.TempDir(),
	})
	if scenarioErr != nil {
		t.Fatal(scenarioErr)
	}
	if err != nil {
		t.Fatal(err)
	}
	if target != len(maps)-1 || frames < maps[target].frames {
		t.Fatalf("scenario stopped early at map %q frame %d", maps[target].name, frames)
	}
}
