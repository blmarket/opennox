package opennox

/*
#include "defs.h"
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/things"

	"github.com/noxworld-dev/opennox/v1/common/unit/ai"
)

func (u *Unit) monsterCast(spellInd spell.ID, target *Object) {
	ud := u.updateDataMonster()
	u.monsterPushAction(ai.DEPENDENCY_UNINTERRUPTABLE)
	sp := u.getServer().SpellDefByInd(spellInd)
	if sp.Def.Flags.Has(things.SpellDuration) {
		ts := gameFrame() + uint32(noxRndCounter1.IntClamp(int(gameFPS()/2), int(2*gameFPS())))
		u.monsterPushAction(ai.DEPENDENCY_TIME, ts)
		u.monsterPushAction(ai.ACTION_CAST_DURATION_SPELL, uint32(spellInd), 0, target)
	} else {
		u.monsterPushAction(ai.ACTION_CAST_SPELL_ON_OBJECT, uint32(spellInd), 0, target)
	}
	if target.CObj() != u.CObj() && !u.monsterActionIsScheduled(ai.ACTION_FLEE) {
		if !sp.Def.Flags.Has(things.SpellTargeted) { // TODO: looks like the flag name is incorrect
			u.monsterPushAction(ai.ACTION_FACE_OBJECT, target)
			u.monsterPushAction(ai.DEPENDENCY_BLOCKED_LINE_OF_FIRE, target)
		}

		v9 := *(*uint32)(unsafe.Add(ud.field_121, 212))
		u.monsterPushAction(ai.DEPENDENCY_OBJECT_FARTHER_THAN, v9, 0, target)
		u.monsterPushAction(ai.DEPENDENCY_OR)
		pos2 := target.Pos()
		u.monsterPushAction(ai.ACTION_MOVE_TO, pos2.X, pos2.Y, target)
	}
}

//export nox_xxx_monsterCast_540A30
func nox_xxx_monsterCast_540A30(cu *C.nox_object_t, spellInd C.int, a3p *C.nox_object_t) {
	asUnitC(cu).monsterCast(spell.ID(spellInd), asObjectC(a3p))
}
