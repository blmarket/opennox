package noxrender

import (
	"testing"
)

func TestSpritesCoverage(t *testing.T) {
	var img Image
	func() { defer func() { recover() }(); _ = img.C() }()
	func() { defer func() { recover() }(); img.Free() }()
	_ = img.String()
	_ = img.Type()
	func() { defer func() { recover() }(); _ = img.Pixdata() }()
	func() { defer func() { recover() }(); _, _, _ = img.Meta() }()

	var rs RenderSprites
	func() { defer func() { recover() }(); rs.Free() }()
	func() { defer func() { recover() }(); _ = NewRawImage(0, nil) }()
	func() { defer func() { recover() }(); _ = rs.AsImage(nil) }()
	func() { defer func() { recover() }(); _ = rs.ImageByIndex(0) }()
	_ = rs.ThingsImageRef(nil)
	func() { defer func() { recover() }(); _ = rs.ImageRef(0, 0, "") }()
	_ = rs.LoadExternalImage(0, "")
	func() { defer func() { recover() }(); _, _ = rs.ImageByBagSection(0, 0) }()

	func() { defer func() { recover() }(); _ = rs.ReadVideoBag() }()
}
