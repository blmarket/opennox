package opennox

import (
	"github.com/noxworld-dev/opennox-lib/spell"
)

func castDetectMagic(spellID spell.ID, _, _, _ *Unit, args *spellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, ENCHANT_DETECTING, lvl, asUnitC(args.Obj), spellBuffConf{
		Dur: 60, DurFPSMul: true,
	})
}
