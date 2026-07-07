package legacy

/*
#include "GAME1.h"
#include "GAME4_1.h"
extern void* dword_5d4594_251560;
extern uint32_t dword_5d4594_1599656;
*/
import "C"
import (
	"image"
	"math"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	Sub_526CA0                               func(a1 string) int
	Nox_xxx_mapSetWallInGlobalDir0pr1_5004D0 func()
	Nox_xxx_map_5004F0                       func()
	Sub_4FF990                               func(a1 uint32)
	Sub_5000B0                               func(a1 *server.Object) int
)

var _ = [1]struct{}{}[12332-unsafe.Sizeof(server.WallDef{})]

func asWallP(p unsafe.Pointer) *server.Wall {
	return (*server.Wall)(p)
}

//export nox_server_getWallAtGrid_410580
func nox_server_getWallAtGrid_410580(x, y int) unsafe.Pointer {
	return GetServer().S().Walls.GetWallAtGrid(image.Pt(x, y)).C()
}

//export nox_xxx_wall_4105E0
func nox_xxx_wall_4105E0(x, y int) unsafe.Pointer {
	return GetServer().S().Walls.GetWallAtGrid2(image.Pt(x, y)).C()
}

//export nox_xxx_wallCreateAt_410250
func nox_xxx_wallCreateAt_410250(x, y int) unsafe.Pointer {
	return GetServer().S().Walls.CreateAtGrid(image.Pt(x, y)).C()
}

//export nox_xxx_mapDelWallAtPt_410430
func nox_xxx_mapDelWallAtPt_410430(x, y int) {
	GetServer().S().Walls.DeleteAtGrid(image.Pt(x, y))
}

//export sub_4106A0
func sub_4106A0(y int) unsafe.Pointer {
	return GetServer().S().Walls.IndexByY(y).C()
}

//export nox_xxx_wallForeachFn_410640
func nox_xxx_wallForeachFn_410640(cfnc unsafe.Pointer, data unsafe.Pointer) {
	GetServer().S().Walls.EachWallXxx(func(it *server.Wall) bool {
		ccall.CallVoidPtr2(cfnc, it.C(), data)
		return true
	})
}

//export sub_57B500
func sub_57B500(x, y int, flags C.char) C.char {
	return C.char(GetServer().S().Sub_57B500(image.Pt(x, y), byte(int8(flags))))
}

//export sub_4D72C0
func sub_4D72C0() C.int {
	return C.int(bool2int(GetServer().S().Doors.Sub_4D72C0()))
}

//export sub_4D72B0
func sub_4D72B0(v C.int) {
	GetServer().S().Doors.Sub_4D72B0(v != 0)
}

//export nox_xxx_wallFlags
func nox_xxx_wallFlags(ind int) uint32 {
	return GetServer().S().Walls.DefByInd(ind).Flags32
}

//export nox_xxx_getWallSprite_46A3B0
func nox_xxx_getWallSprite_46A3B0(ind int, a2 int, a3 int, a4 int) unsafe.Pointer {
	return GetServer().S().Walls.DefByInd(ind).Sprite(a2, a3, a4)
}

//export nox_xxx_getWallDrawOffset_46A3F0
func nox_xxx_getWallDrawOffset_46A3F0(ind int, a2 int, a3 int, a4 int, px, py *C.int) {
	v := GetServer().S().Walls.DefByInd(ind).DrawOffset(a2, a3, a4)
	*px = C.int(v.X)
	*py = C.int(v.Y)
}

//export nox_xxx_mapWallMaxVariation_410DD0
func nox_xxx_mapWallMaxVariation_410DD0(ind int, a2 int, a3 int) byte {
	return GetServer().S().Walls.DefByInd(ind).Variations(a2, a3)
}

//export nox_xxx_map_410E00
func nox_xxx_map_410E00(ind int) byte {
	return GetServer().S().Walls.DefByInd(ind).Field749
}

//export nox_xxx_mapWallGetHpByTile_410E20
func nox_xxx_mapWallGetHpByTile_410E20(ind int) byte {
	return GetServer().S().Walls.DefByInd(ind).Health41
}

//export nox_xxx_wallFindOpenSound_410EE0
func nox_xxx_wallFindOpenSound_410EE0(ind int) *C.char {
	return internCStr(GetServer().S().Walls.DefByInd(ind).OpenSound())
}

//export nox_xxx_wallFindCloseSound_410F20
func nox_xxx_wallFindCloseSound_410F20(ind int) *C.char {
	return internCStr(GetServer().S().Walls.DefByInd(ind).CloseSound())
}

//export nox_xxx_wallTileByName_410D60
func nox_xxx_wallTileByName_410D60(name *C.char) int {
	return GetServer().S().Walls.DefIndByName(GoString(name))
}

//export sub_526CA0
func sub_526CA0(a1 *C.char) int {
	return Sub_526CA0(GoString(a1))
}

//export nox_xxx_mapSetWallInGlobalDir0pr1_5004D0
func nox_xxx_mapSetWallInGlobalDir0pr1_5004D0() {
	Nox_xxx_mapSetWallInGlobalDir0pr1_5004D0()
}

//export nox_xxx_map_5004F0
func nox_xxx_map_5004F0() {
	Nox_xxx_map_5004F0()
}

//export sub_4FF990
func sub_4FF990(a1 C.uint) {
	Sub_4FF990(uint32(a1))
}

//export sub_5000B0
func sub_5000B0(a1 *nox_object_t) int {
	return Sub_5000B0(asObjectS(a1))
}

//export nox_xxx_mapDamageToWalls_534FC0
func nox_xxx_mapDamageToWalls_534FC0(a1 *C.int4, a2 unsafe.Pointer, a3 C.float, a4, a5 int, a6 unsafe.Pointer) C.bool {
	rect := image.Rect(int(a1.field_0), int(a1.field_4), int(a1.field_8), int(a1.field_C))
	return C.bool(GetServer().Nox_xxx_mapDamageToWalls_534FC0(rect, *(*types.Pointf)(a2), float32(a3), a4, object.DamageType(a5), AsObjectP(a6)))
}

//export nox_xxx_damageToMap_534BC0
func nox_xxx_damageToMap_534BC0(a1, a2, a3, a4 int, a5 *nox_object_t) int {
	return GetServer().Nox_xxx_damageToMap_534BC0(a1, a2, a3, object.DamageType(a4), asObjectS(a5))
}

//export nox_xxx_wallBreackableListAdd_410840
func nox_xxx_wallBreackableListAdd_410840(a1 unsafe.Pointer) {
	GetServer().S().Walls.AddBreakable(asWallP(a1))
}

//export nox_xxx_wallSecretNext_410790
func nox_xxx_wallSecretNext_410790(a1 *C.int) C.int {
	if a1 != nil {
		return *a1
	}
	return 0
}

//export nox_xxx_wall_4DF1E0
func nox_xxx_wall_4DF1E0(a1 int) {
	GetServer().Nox_xxx_wall_4DF1E0(a1)
}

func Sub_5071C0() bool {
	return C.dword_5d4594_1599656 != 0
}

//export nox_xxx_math_509ED0
func nox_xxx_math_509ED0(a1 *C.float2) C.int {
	x := float64(a1.field_0)
	y := float64(a1.field_4)
	v2 := (math.Atan2(y, x)+6.2831855)*40.743664 + 0.5
	result := int(v2)
	if result < 0 {
		result += int((uint32(255-result) >> 8) << 8)
	}
	if result >= 256 {
		result += -256 * int(uint32(result)>>8)
	}
	return C.int(result)
}

func Nox_xxx_math_509ED0(pos types.Pointf) int {
	cpos, free := alloc.New(types.Pointf{})
	defer free()
	*cpos = pos
	return int(C.nox_xxx_math_509ED0((*C.float2)(unsafe.Pointer(cpos))))
}

func Nox_xxx_math_509EA0(a1 int) int {
	return int(C.nox_xxx_math_509EA0(C.int(a1)))
}

//export nox_xxx_utilNormalizeVector_509F20
func nox_xxx_utilNormalizeVector_509F20(a1 *C.float2) {
	x := float64(a1.field_0)
	y := float64(a1.field_4)
	v1 := math.Sqrt(x*x + y*y)
	a1.field_0 = C.float(x / v1)
	a1.field_4 = C.float(y / v1)
}

//export nox_xxx_monsterActionIsCondition_50A010
func nox_xxx_monsterActionIsCondition_50A010(a1 C.int) C.int {
	if a1 > 39 {
		return 1
	}
	return 0
}

//export nox_xxx_mapGenCheckRoomType_5238F0
func nox_xxx_mapGenCheckRoomType_5238F0(a1 *C.int) C.int {
	v := *a1
	if v == 2 || v == 3 || v == 4 || v == 5 {
		return 1
	}
	return 0
}
