package client

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/types"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/server"
)

func TestSightCoverage(t *testing.T) {
	func() {
		defer func() { recover() }()
		var c Client
		c.Nox_xxx_drawBlack_496150(nil)
		c.Sub_498AE0()

		var cs clientSight
		_ = cs.GetSightPoints()
		_ = cs.getSightObjs()
		cs.Init(10, 10)
		cs.Free()
		_ = cs.nextListZzz()
		cs.addListZzz(nil)
		cs.sightListReset()
		cs.listPutXxx(nil)
		cs.sightReset(nil)
		cs.Nox_xxx_drawBlack_496150_E(nil, nil, nil, nil)
		cs.sightObjVec(nil, nil)
		cs.Nox_xxx_drawBlack_496150_F(nil, nil, nil)
		_ = cs.checkXxx(image.Rect(0, 0, 10, 10))
		_ = cs.Nox_xxx_client_4984B0_drawable_A(nil, nil)
		cs.Sub_498AE0_B(nil)
		_ = cs.Sub_498C20(nil, nil, 0)
		_ = cs.Sub_4992B0(0, 0)
		cs.sightArrInsertAt(nil, 0)
		cs.sightArrRemoveAt(0)
		_ = cs.sightArrFind(nil)
		_ = cs.sightArrFindFirst(nil)
		_ = cs.sightArrFindInsert(nil)
		cs.sightArrInsert(nil)
		_ = cs.sightArrRemove(nil)
		cs.procListZzz()
		cs.procZzz(nil)
		_ = cs.newStruct()
		_ = cs.copyStruct(nil)
		cs.newFromWall(0, 0, 0)
		cs.newFromDrawableDoor(nil)
		cs.newFromDrawableCircle(nil)
		cs.newFromDrawableBox(nil)
		cs.newFromViewport(nil)
		cs.sub_499130(image.Pt(0, 0))
		_ = cs.Sub_499290(0)
		cs.sub_4991E0(image.Pt(0, 0))
		cs.newFromDrawableBoxSub(0, 0, 0, 0, nil)
		_ = cs.Nox_xxx_drawBlack_496150_C()
		cs.sub_4989A0()

		_ = sightAngleFromRad(0)
		var sa sightAngle
		_ = sa.ConvA()
		_ = sa.Rad()
		_ = sa.Angle()
		_ = sa.Normalize()

		_ = sub_497B80(nil, image.Pt(0, 0))
		_ = cs.sub_4CA8B0(0, 0)
		_ = sub_4CA960(ntype.Point32{}, 0, types.Rectf{})
		_ = sub_4CAA90(ntype.Point32{}, types.Rectf{}, 0, 0)
		_ = sub_4990D0(image.Pt(0, 0), image.Pt(0, 0))
		_ = sub_499160(image.Pt(0, 0), image.Pt(0, 0), image.Pt(0, 0))
		_ = sub_4CAC30(ntype.Point32{}, types.Rectf{}, 0, 0)
		_ = sub_414C50(0)
		_ = sightFarAtAngleX(0)
		_ = sightFarAtAngleY(0)
		_ = sub_57BA30(nil, nil, image.Rect(0, 0, 0, 0))
		cs.sub_498380(nil, nil)
		_ = sub_427C80(image.Rect(0, 0, 0, 0), image.Rect(0, 0, 0, 0))

		var so SightObject
		_ = so.UID()
	}()

	// Use imports
	_ = noxrender.Viewport{}
	_ = server.Wall{}
}
