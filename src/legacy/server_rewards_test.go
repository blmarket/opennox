package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerRewardReportsIgnoreNonPlayerObjects(t *testing.T) {
	require.NotZero(t, C_nox_xxx_netAbilityReport_4D8060_nonPlayer())
	require.Equal(t, 0, C_nox_xxx_abilityRewardServ_4FB9C0_nonPlayer())
	require.NotZero(t, C_nox_xxx_netReportGuideAward_4D8000_nonPlayer())
	require.Equal(t, 0, C_nox_xxx_awardBeastGuide_4FAE80_nonPlayer())
	require.NotZero(t, C_nox_xxx_netSendSpellAward_4D7F90_nonPlayer())
	require.Equal(t, 0, C_nox_xxx_spellGrantToPlayer_4FB550_nonPlayer())
}

func TestExpLevelProtectionBelowProtectedRange(t *testing.T) {
	require.Equal(t, uintptr(12345), C_sub_56F980_belowProtectedRange(12345, 7))
}
