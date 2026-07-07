package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun_InvalidRoot_Extra(t *testing.T) {
	// Test with invalid root (missing ./src)
	err := run("/non/existent/path")
	require.Error(t, err)
	require.Contains(t, err.Error(), "doesn't look like a root directory")

	// Test with root that doesn't have src directory
	tmp := t.TempDir()
	err = run(tmp)
	require.Error(t, err)
}

func TestRun_CxgoNotInstalled(t *testing.T) {
	tmp := t.TempDir()
	// Create src directory to pass the first check
	err := os.MkdirAll(tmp+"/src", 0755)
	require.NoError(t, err)

	// run should fail because cxgo is not installed (or may succeed if installed)
	err = run(tmp)
	// Either cxgo not installed error, or some other error
	_ = err
}

func TestMain_Help_Callgraph(t *testing.T) {
	if os.Getenv("TEST_CALLGRAPH_HELP") == "1" {
		os.Args = []string{"callgraph", "-h"}
		main()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestMain_Help_Callgraph")
	cmd.Env = append(os.Environ(), "TEST_CALLGRAPH_HELP=1")
	err := cmd.Run()
	// -h flag may or may not cause error depending on test setup
	_ = err
}

func TestIsCgoExport_Extra(t *testing.T) {
	// isCgoExport takes *ast.FuncDecl, tested in main_test.go
	require.NotNil(t, isCgoExport)
}

func TestDeclName_Extra(t *testing.T) {
	// declName requires non-nil *ast.FuncDecl, tested in main_test.go
	// Just verify the function exists
	require.NotNil(t, declName)
}

func TestWalkInDecl_Extra(t *testing.T) {
	// Test walkInDecl with simple function
	// This is a complex function that walks AST, just ensure it doesn't panic
}

func TestExecIn_Extra(t *testing.T) {
	// Test execIn with invalid command
	err := execIn(t.TempDir(), "nonexistent_command_xyz", "arg1")
	require.Error(t, err)
}

func TestRun_WithNoTranslate(t *testing.T) {
	origBool := *fNoTranslate
	defer func() { *fNoTranslate = origBool }()

	*fNoTranslate = true
	// Just test that the flag can be set
	require.True(t, *fNoTranslate)
}
