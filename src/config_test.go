package opennox

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestConfigStrPtr(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	var s string
	configStrPtr("test.key", "", "defaultVal", &s)
	if s != "defaultVal" {
		t.Errorf("configStrPtr default = %q, want defaultVal", s)
	}
	viper.Set("test.key", "newVal")
	// trigger onConfigRead callbacks
	for _, fnc := range onConfigRead {
		fnc()
	}
	if s != "newVal" {
		t.Errorf("configStrPtr after read = %q, want newVal", s)
	}
}

func TestConfigBoolPtr(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	var b bool
	configBoolPtr("test.bool", "", true, &b)
	if !b {
		t.Error("configBoolPtr default should be true")
	}
	viper.Set("test.bool", false)
	for _, fnc := range onConfigRead {
		fnc()
	}
	if b {
		t.Error("configBoolPtr after set should be false")
	}
}

func TestConfigHiddenBoolPtr(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	var b bool
	viper.Set("test.hidden", true)
	configHiddenBoolPtr("test.hidden", "", &b)
	if !b {
		t.Error("configHiddenBoolPtr should read true")
	}
}

func TestReadConfig(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "opennox.yml")
	if err := os.WriteFile(cfgPath, []byte("test_key: test_val\n"), 0644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	err := readConfig(cfgPath)
	if err != nil {
		t.Errorf("readConfig failed: %v", err)
	}
	if viper.GetString("test_key") != "test_val" {
		t.Errorf("config value = %q, want test_val", viper.GetString("test_key"))
	}
}

func TestReadConfigNotFound(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	configDirty = false
	err := readConfig("/nonexistent/path/opennox.yml")
	if err != nil {
		t.Errorf("readConfig nonexistent should not error, got %v", err)
	}
	if !configDirty {
		t.Error("configDirty should be true after missing config")
	}
	configDirty = false
}
