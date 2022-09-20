package opennox

import (
	"github.com/noxworld-dev/opennox-lib/spell"
)

func castDeath(spellID spell.ID, _, a3, _ *Unit, args *spellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, ENCHANT_DEATH, lvl, asUnitC(args.Obj), spellBuffConf{
		Dur: 450, Once: true, Orig: a3, Offensive: true,
	})
}
