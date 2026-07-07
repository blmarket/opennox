package opennox

import (
	"testing"

	"github.com/spf13/viper"
)

func TestConfigHiddenBoolPtrCallback(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()

	var val bool
	viper.Set("test.hidden.callback2", true)
	configHiddenBoolPtr("test.hidden.callback2", "", &val)

	if !val {
		t.Error("configHiddenBoolPtr should set val to true initially")
	}

	// Change the value in viper
	viper.Set("test.hidden.callback2", false)

	// Trigger onConfigRead callbacks
	for _, fnc := range onConfigRead {
		fnc()
	}

	// val should be updated to false
	if val {
		t.Error("configHiddenBoolPtr callback should update val to false")
	}
}

func TestConfigHiddenBoolPtrWithEnv(t *testing.T) {
	viper.Reset()
	defer viper.Reset()
	old := onConfigRead
	onConfigRead = nil
	defer func() { onConfigRead = old }()

	var val bool
	// Test with env var - just ensure it doesn't panic
	configHiddenBoolPtr("test.hidden.env", "TEST_HIDDEN_ENV_VAR", &val)
}
