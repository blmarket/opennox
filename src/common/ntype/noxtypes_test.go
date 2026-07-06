package ntype

import (
	"image"
	"testing"
)

func TestPoint32_Point(t *testing.T) {
	p := Point32{X: 10, Y: -5}
	pt := p.Point()
	if pt != (image.Point{X: 10, Y: -5}) {
		t.Fatalf("unexpected point: %v", pt)
	}
}

func TestPlayerInd(t *testing.T) {
	var pi PlayerInd = 3
	if int(pi) != 3 {
		t.Fatalf("unexpected value")
	}
}
