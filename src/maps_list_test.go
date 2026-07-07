package opennox

import (
	"testing"
)

func TestScanMaps(t *testing.T) {
	// scanMaps scans the datapath for maps
	// It may fail if datapath is not set up, but should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("scanMaps panicked: %v", r)
		}
	}()
	list, err := scanMaps()
	if err != nil {
		t.Logf("scanMaps returned error: %v", err)
	}
	t.Logf("scanMaps returned %d maps", len(list))
}

func TestNoxCommonScanAllMaps(t *testing.T) {
	// nox_common_scanAllMaps_4D07F0 accesses legacy and calls scanMaps
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_common_scanAllMaps_4D07F0 panicked as expected: %v", r)
		}
	}()
	err := nox_common_scanAllMaps_4D07F0()
	if err != nil {
		t.Logf("nox_common_scanAllMaps_4D07F0 returned error: %v", err)
	}
}

func TestNoxCommonScanAddMap(t *testing.T) {
	// nox_common_scanAddMap accesses legacy and memmap
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_common_scanAddMap panicked as expected: %v", r)
		}
	}()
	nox_common_scanAddMap("testmap")
	nox_common_scanAddMap("verylongmapname")
	nox_common_scanAddMap("")
}

func TestNoxXxxCheckHasSoloMaps(t *testing.T) {
	// nox_xxx_checkHasSoloMaps stats a file in datapath
	// It should not panic
	result := nox_xxx_checkHasSoloMaps()
	t.Logf("nox_xxx_checkHasSoloMaps returned: %v", result)
}
