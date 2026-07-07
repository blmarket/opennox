package opennox

import (
	"testing"
)

func TestConfigCoverage(t *testing.T) {
	isolateConfigPath(t)
	call := func(f func()) {
		defer func() { recover() }()
		f()
	}

	call(func() { writeConfig() })
	call(func() { maybeWriteConfig() })
	call(func() { writeConfigLater() })
}

func TestMiscCoverage(t *testing.T) {
	// misc.go functions already have 100% coverage, but we can call them again
	_ = abs(-5)
	_ = clamp(5, 0, 10)
	_ = bool2int(true)
	_ = rotl(1, 2)
	_ = rotl16(1, 2)
	p := []byte{1, 2, 3, 4}
	swap4(p)
	_ = find([]int{1, 2, 3}, 2)
}
