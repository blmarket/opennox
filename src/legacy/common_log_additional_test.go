package legacy

import (
	"testing"
)

func TestCommonLogAdditional(t *testing.T) {
	// Hook NetworkLogPrint to avoid nil pointer
	oldLog := NetworkLogPrint
	defer func() { NetworkLogPrint = oldLog }()
	NetworkLogPrint = func(str string) {
		// Just consume the log message
		_ = str
	}

	// Test NetworkLogPrintf with various messages
	t.Run("NetworkLogPrintf with empty string", func(t *testing.T) {
		C_nox_xxx_networkLog_printf_s("%s", "")
	})

	t.Run("NetworkLogPrintf with long string", func(t *testing.T) {
		longStr := ""
		for i := 0; i < 100; i++ {
			longStr += "This is a long log message. "
		}
		C_nox_xxx_networkLog_printf_s("%s", longStr)
	})

	t.Run("NetworkLogPrintf with special characters", func(t *testing.T) {
		messages := []string{
			"Simple message",
			"Message with numbers 12345",
			"Message with symbols !@#$%^&*()",
			"Message with newline and tab characters",
			"",
			" ",
		}
		for _, msg := range messages {
			C_nox_xxx_networkLog_printf_s("%s", msg)
		}
	})

	t.Run("NetworkLogPrintf multiple calls", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			C_nox_xxx_networkLog_printf_d("Log message %d", i)
		}
	})

	t.Run("NetworkLogPrintf with format strings", func(t *testing.T) {
		// Test with format-like strings (should be treated as literal)
		C_nox_xxx_networkLog_printf_s("%s", "Test %s %d")
		C_nox_xxx_networkLog_printf_s("%s", "100% complete")
	})

	t.Run("NetworkLogPrintf with integers", func(t *testing.T) {
		C_nox_xxx_networkLog_printf_d("Value: %d", 0)
		C_nox_xxx_networkLog_printf_d("Value: %d", -1)
		C_nox_xxx_networkLog_printf_d("Value: %d", 2147483647)
		C_nox_xxx_networkLog_printf_d("Value: %d", -2147483648)
	})
}
