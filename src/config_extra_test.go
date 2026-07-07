package opennox

import (
	"testing"

	"github.com/spf13/viper"
)

func TestConfigStrPtrExtra(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	var s string
	configStrPtr("test.key2", "", "defaultVal2", &s)
	if s != "defaultVal2" {
		t.Errorf("configStrPtr default = %q, want defaultVal2", s)
	}
}

func TestConfigBoolPtrExtra(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	var b bool
	configBoolPtr("test.bool2", "", false, &b)
	if b {
		t.Error("configBoolPtr default should be false")
	}
}

func TestConfigHiddenBoolPtrExtra(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()
	var b bool
	viper.Set("test.hidden2", false)
	configHiddenBoolPtr("test.hidden2", "", &b)
	if b {
		t.Error("configHiddenBoolPtr should read false")
	}
}

func TestRegisterOnConfigReadExtra(t *testing.T) {
	called := false
	fn := func() {
		called = true
	}
	registerOnConfigRead(fn)
	// The function should be registered without panic
	_ = called
}

func TestWriteConfigLaterExtra(t *testing.T) {
	// Just ensure it doesn't panic
	writeConfigLater()
}

func TestMaybeWriteConfig(t *testing.T) {
	// Just ensure it doesn't panic
	isolateConfigPath(t)
	configDirty = true
	maybeWriteConfig()
}
