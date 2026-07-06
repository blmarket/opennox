package noxrender

import (
	"testing"
)

func TestPixOpOverMultiplyAlpha50(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	dnext, snext := r.pixOpOverMultiplyAlpha50(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOverMultiplyAlpha50 should consume all")
	}
}

func TestPixOpOver4444Multiply(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	dnext, snext := r.pixOpOver4444Multiply(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOver4444Multiply should consume all")
	}
}

func TestPixOpOver4444Alpha(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	dnext, snext := r.pixOpOver4444Alpha(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOver4444Alpha should consume all")
	}
}

func TestPixOpOverAlpha(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	dnext, snext := r.pixOpOverAlpha(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOverAlpha should consume all")
	}
}

func TestPixOpOverMultiplyAlpha(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	dnext, snext := r.pixOpOverMultiplyAlpha(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOverMultiplyAlpha should consume all")
	}
}

func TestPixOpSrcColorize(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	dnext, snext := r.pixOpSrcColorize(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpSrcColorize should consume all")
	}
}

func TestPixBlendPremult(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	dnext, snext := r.pixBlendPremult(dst, src, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixBlendPremult should consume all")
	}
}

func TestPixOpOverAlphaIndexed(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{1, 2, 3, 4}

	dnext, snext := r.pixOpOverAlphaIndexed(dst, src, 0, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOverAlphaIndexed should consume all")
	}
}

func TestPixOpOverMultiplyAlpha50Indexed(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{1, 2, 3, 4}

	dnext, snext := r.pixOpOverMultiplyAlpha50Indexed(dst, src, 0, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOverMultiplyAlpha50Indexed should consume all")
	}
}

func TestPixOpOverMultiplyAlphaIndexed(t *testing.T) {
	r := NewRender(nil)
	d, free := NewRenderData()
	defer free()
	r.p = d
	dst := make([]uint16, 4)
	src := []byte{1, 2, 3, 4}

	dnext, snext := r.pixOpOverMultiplyAlphaIndexed(dst, src, 0, 4)
	if len(dnext) != 0 || len(snext) != 0 {
		t.Error("pixOpOverMultiplyAlphaIndexed should consume all")
	}
}
