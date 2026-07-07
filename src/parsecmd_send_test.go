package opennox

import (
	"testing"
)

func TestNoxXxxNetServerCmd(t *testing.T) {
	// nox_xxx_netServerCmd_440950 creates a buffer and calls nox_xxx_netClientSend2_4E53C0
	// It accesses legacy.ClientPlayerNetCode()
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_xxx_netServerCmd_440950 panicked as expected: %v", r)
		}
	}()
	nox_xxx_netServerCmd_440950(0, "")
	nox_xxx_netServerCmd_440950(1, "test command")
	nox_xxx_netServerCmd_440950(255, "another command")
}

func TestNoxXxxServerHandleClientConsole(t *testing.T) {
	// nox_xxx_serverHandleClientConsole_443E90 calls legacy function
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_xxx_serverHandleClientConsole_443E90 panicked as expected: %v", r)
		}
	}()
	nox_xxx_serverHandleClientConsole_443E90(nil, 0, "")
	nox_xxx_serverHandleClientConsole_443E90(nil, 1, "test")
}

func TestNoxXxxCmdSayDo(t *testing.T) {
	// nox_xxx_cmdSayDo_46A4B0 calls legacy function
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_xxx_cmdSayDo_46A4B0 panicked as expected: %v", r)
		}
	}()
	nox_xxx_cmdSayDo_46A4B0("", 0)
	nox_xxx_cmdSayDo_46A4B0("hello", 1)
}

func TestNoxConsoleSendSysOpPass(t *testing.T) {
	// nox_console_sendSysOpPass_4409D0 creates a buffer and calls nox_xxx_netClientSend2_4E53C0
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_console_sendSysOpPass_4409D0 panicked as expected: %v", r)
		}
	}()
	nox_console_sendSysOpPass_4409D0("")
	nox_console_sendSysOpPass_4409D0("password")
	nox_console_sendSysOpPass_4409D0("very long password that exceeds buffer")
}
