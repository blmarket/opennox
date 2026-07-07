package noxrender

import (
	"testing"
)

func TestColor16_Extra2(t *testing.T) {
	var c Color16
	_ = c
}

func TestRenderData_Extra2(t *testing.T) {
	rd, free := NewRenderData()
	if rd == nil {
		t.Fatal("NewRenderData should not return nil")
	}
	if free != nil {
		free()
	}
}

func TestFadeKey_Extra2(t *testing.T) {
	var fk FadeKey
	_ = fk
}
