package legacy

import (
	"testing"
)

func TestNetworkLogPrintf(t *testing.T) {
	// Hook NetworkLogPrint
	oldLog := NetworkLogPrint
	defer func() { NetworkLogPrint = oldLog }()

	var received string
	NetworkLogPrint = func(str string) {
		received = str
	}

	C_nox_xxx_networkLog_printf_s("User logged in: %s", "JohnDoe")
	expected := "User logged in: JohnDoe"
	if received != expected {
		t.Errorf("Expected network log print %q, got %q", expected, received)
	}

	C_nox_xxx_networkLog_printf_d("Value code: %d", 12345)
	expected2 := "Value code: 12345"
	if received != expected2 {
		t.Errorf("Expected network log print %q, got %q", expected2, received)
	}
}
