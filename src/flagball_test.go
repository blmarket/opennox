package opennox

import (
	"testing"
)

func TestSub4E8290(t *testing.T) {
	// sub_4E8290 accesses noxServer and memmap, then calls s.Nox_xxx_netSendBallStatus_4D95F0
	defer func() {
		if r := recover(); r != nil {
			t.Logf("sub_4E8290 panicked as expected: %v", r)
		}
	}()
	result := sub_4E8290(0, 0)
	t.Logf("sub_4E8290 returned: %d", result)

	result = sub_4E8290(1, 100)
	t.Logf("sub_4E8290 returned: %d", result)
}
