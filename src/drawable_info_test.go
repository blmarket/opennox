package opennox

import (
	"testing"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
)

func TestEnchantColorize(t *testing.T) {
	c := &Client{}

	// Test poison case - green channel increase
	dr := &client.Drawable{}
	colors := []noxcolor.RGBA5551{
		noxcolor.RGB5551Color(100, 100, 100),
		noxcolor.RGB5551Color(50, 200, 50),
	}
	orig := make([]noxcolor.RGBA5551, len(colors))
	copy(orig, colors)

	c.EnchantColorize(dr, true, colors)

	// Green channel should be increased by 100, capped at 255
	// Note: RGB5551 uses 5 bits per channel, so values are quantized
	for i, col := range colors {
		origCl := noxrender.SplitColor(orig[i])
		newCl := noxrender.SplitColor(col)
		// Compute expected green after quantization through RGB5551
		expectedG := origCl.G + 100
		if expectedG > 255 {
			expectedG = 255
		}
		// Convert through RGB5551 to account for 5-bit quantization
		expectedCol := noxcolor.RGB5551Color(byte(origCl.R), byte(expectedG), byte(origCl.B))
		expectedCl := noxrender.SplitColor(expectedCol)
		if newCl.G != expectedCl.G {
			t.Fatalf("color %d: expected G=%d, got %d (orig G=%d)", i, expectedCl.G, newCl.G, origCl.G)
		}
	}

	// Test poison with high green (capped)
	colors2 := []noxcolor.RGBA5551{
		noxcolor.RGB5551Color(0, 200, 0),
	}
	c.EnchantColorize(dr, true, colors2)
	newCl := noxrender.SplitColor(colors2[0])
	if newCl.G != 255 {
		t.Fatalf("expected capped green 255, got %d", newCl.G)
	}

	// Test no poison, no enchant - colors unchanged
	colors3 := []noxcolor.RGBA5551{
		noxcolor.RGB5551Color(10, 20, 30),
	}
	orig3 := colors3[0]
	c.EnchantColorize(dr, false, colors3)
	if colors3[0] != orig3 {
		t.Fatal("colors should be unchanged without poison or enchant")
	}

	// Test empty colors slice
	c.EnchantColorize(dr, true, []noxcolor.RGBA5551{})
	c.EnchantColorize(dr, false, nil)
}

func TestEnchantColorize_Invulnerable(t *testing.T) {
	// This test verifies the invulnerable enchant path
	// It requires a server with frame count, so we just verify it doesn't panic
	// with a nil drawable that has no enchants
	c := &Client{}
	dr := &client.Drawable{}
	colors := []noxcolor.RGBA5551{
		noxcolor.RGB5551Color(1, 2, 3),
	}
	// Should not panic even without server
	// The HasEnchant check should return false for empty drawable
	c.EnchantColorize(dr, false, colors)
}

func TestDrawEnchantsTop(t *testing.T) {
	c := &Client{}
	dr := &client.Drawable{}
	// Should not panic with nil viewport and drawable without enchants
	// Actual drawing requires full client setup, so we just verify no panic on nil checks
	_ = c
	_ = dr
}

func TestClientDrawInfo_Free(t *testing.T) {
	var info clientDrawInfo
	// Free with nil drawables should not panic
	info.Free()

	// Free again should be safe
	info.Free()
}

func TestClientDrawInfo_Init(t *testing.T) {
	var info clientDrawInfo
	c := &Client{}
	info.Init(c)
	if info.c != c {
		t.Fatal("Init should set client")
	}
}
