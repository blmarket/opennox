package noxrender

import (
	"testing"
)

func TestCopy16b(t *testing.T) {
	dst := make([]uint16, 4)
	src := []byte{0x01, 0x00, 0x02, 0x00, 0x03, 0x00, 0x04, 0x00} // little endian 1,2,3,4
	n := copy16b(dst, src)
	if n != 4 {
		t.Errorf("copy16b n = %d, want 4", n)
	}
	if dst[0] != 1 || dst[1] != 2 || dst[2] != 3 || dst[3] != 4 {
		t.Error("copy16b values mismatch")
	}
	// Test odd length src
	dst2 := make([]uint16, 2)
	src2 := []byte{0x01, 0x00, 0x02} // odd, should only copy 1
	n = copy16b(dst2, src2)
	if n != 1 {
		t.Errorf("copy16b odd n = %d, want 1", n)
	}
}

func TestSkipPixdata(t *testing.T) {
	// Simple pixdata: op=0, val=2 (skip 2), then op=0, val=1
	pix := []byte{
		0, 2, // op 0, val 2
		0, 1, // op 0, val 1
		0, 1, // another row
	}
	result := skipPixdata(pix, 3, 1) // width 3, skip 1 row
	// After skipping 1 row of width 3: first op val=2 (covers 2), then op val=1 (covers 1) = total 3
	// Should be at position after first row
	if len(result) == 0 {
		t.Error("skipPixdata should not return empty")
	}
}

func TestPixOpSrc(t *testing.T) {
	dst := make([]uint16, 4)
	src := []byte{0x01, 0x00, 0x02, 0x00, 0x03, 0x00, 0x04, 0x00}
	dnext, snext := pixOpSrc(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpSrc should consume all")
	}
	if dst[0] != 1 || dst[3] != 4 {
		t.Error("pixOpSrc values mismatch")
	}
}

func TestDrawOpU16(t *testing.T) {
	r := &NoxRender{}
	dst := make([]uint16, 4)
	src := []byte{0x01, 0x00, 0x02, 0x00, 0x03, 0x00, 0x04, 0x00}
	dnext, snext := r.drawOpU16(dst, src, 4, func(old uint16, src uint16) uint16 {
		return old + src
	})
	if len(dnext) != 0 {
		t.Error("drawOpU16 dnext should be empty")
	}
	if len(snext) != 0 {
		t.Error("drawOpU16 snext should be empty")
	}
	if dst[0] != 1 || dst[1] != 2 {
		t.Error("drawOpU16 values mismatch")
	}
}

func TestDrawOpU8(t *testing.T) {
	r := &NoxRender{}
	dst := make([]uint16, 4)
	src := []byte{1, 2, 3, 4}
	dnext, snext := r.drawOpU8(dst, src, 4, func(old uint16, src byte) uint16 {
		return old + uint16(src)
	})
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("drawOpU8 should consume all")
	}
	if dst[0] != 1 || dst[3] != 4 {
		t.Error("drawOpU8 values mismatch")
	}
}

func TestRawImage(t *testing.T) {
	data := make([]byte, 10*20*2)
	img := NewRawImage16(2, data)
	if img == nil {
		t.Fatal("NewRawImage16 should not return nil")
	}
	if img.Type() != 2 {
		t.Errorf("Type = %d, want 2", img.Type())
	}
	pix := img.Pixdata()
	if len(pix) != 10*20*2 {
		t.Errorf("Pixdata len = %d, want %d", len(pix), 10*20*2)
	}
}

func TestPixOpFunctions(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	// Test pixOpOverAlpha50
	dnext, snext := r.pixOpOverAlpha50(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOverAlpha50 should consume all")
	}

	// Test pixOpSrcMultiply
	dst = make([]uint16, 4)
	dnext, snext = r.pixOpSrcMultiply(dst, src, 4)
	if len(dnext) != 0 {
		t.Error("pixOpSrcMultiply dnext should be empty")
	}

	// Test pixOpOver4444
	dst = make([]uint16, 4)
	dnext, snext = r.pixOpOver4444(dst, src, 4)
	if len(dnext) != 0 {
		t.Error("pixOpOver4444 dnext should be empty")
	}
}

func TestPixOpIndexedFunctions(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{1, 2, 3, 4}

	dnext, snext := r.pixOpSrcIndexed(dst, src, 0, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpSrcIndexed should consume all")
	}

	dst = make([]uint16, 4)
	dnext, snext = r.pixOpSrcMultiplyIndexed(dst, src, 0, 4)
	if len(dnext) != 0 {
		t.Error("pixOpSrcMultiplyIndexed dnext should be empty")
	}

	dst = make([]uint16, 4)
	dnext, snext = r.pixOpOverAlpha50Indexed(dst, src, 0, 4)
	if len(dnext) != 0 {
		t.Error("pixOpOverAlpha50Indexed dnext should be empty")
	}
}
