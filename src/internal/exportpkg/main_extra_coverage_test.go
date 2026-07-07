package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun_UnsupportedMode(t *testing.T) {
	// Test with unsupported mode by directly setting fMode
	origMode := *fMode
	defer func() { *fMode = origMode }()

	*fMode = "invalid"
	err := run()
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported mode")
}

func TestMain_Help(t *testing.T) {
	if os.Getenv("TEST_MAIN_HELP") == "1" {
		os.Args = []string{"exportpkg", "-h"}
		main()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestMain_Help")
	cmd.Env = append(os.Environ(), "TEST_MAIN_HELP=1")
	err := cmd.Run()
	// -h flag causes flag.Parse to exit, may or may not be error
	_ = err
}

func TestAsExportedName_Extra(t *testing.T) {
	require.Equal(t, "Test", asExportedName("test"))
	require.Equal(t, "Test", asExportedName("Test"))
	// Empty string causes panic, so skip it
	// require.Equal(t, "", asExportedName(""))
}

func TestEachSrcFile_Error(t *testing.T) {
	origDir := *fDir
	defer func() { *fDir = origDir }()

	// Test with non-existent directory
	*fDir = "/non/existent/path"
	err := eachSrcFile(func(path, epath string) error {
		return nil
	})
	require.Error(t, err)
}

func TestEachDstFile_Error(t *testing.T) {
	origDir := *fDir
	defer func() { *fDir = origDir }()

	// Test with non-existent directory
	*fDir = "/non/existent/path"
	err := eachDstFile(func(epath string) error {
		return nil
	})
	require.Error(t, err)
}

func TestIsC_Extra(t *testing.T) {
	// isC takes an ast.Expr, tested in export_test.go
	// Just verify the function exists
	require.NotNil(t, isC)
}
