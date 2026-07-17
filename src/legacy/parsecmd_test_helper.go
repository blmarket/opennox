package legacy

/*
#include <stdint.h>

int nox_cmd_set_sysop(int, int, unsigned short**);
int nox_cmd_set_cycle(int, int, unsigned short**);
int nox_cmd_set_weapons(int, int, unsigned short**);
int nox_cmd_set_staffs(int, int, unsigned short**);
int nox_cmd_set_name(int, int, unsigned short**);
int nox_cmd_set_mnstrs(int, int, unsigned short**);
int nox_cmd_set_spell(int, int, unsigned short**);
int nox_cmd_ban(int, int, unsigned short**);
int nox_cmd_kick(int, int, unsigned short**);
int nox_cmd_set_players(int, int, unsigned short**);
int nox_cmd_unmute(int, int, unsigned short**);
int nox_cmd_mute(int, int, unsigned short**);
int nox_cmd_exec_rul(int, int, unsigned short**);
int nox_cmd_offonly1(int, int, unsigned short**);
int nox_cmd_offonly2(int, int, unsigned short**);
int nox_cmd_show_motd(int, int, unsigned short**);
int nox_cmd_show_seq(int, int, unsigned short**);
int nox_cmd_cheat_level(int, int, unsigned short**);
int nox_cmd_set_time(int, int, unsigned short**);
int nox_cmd_set_lessons(int, int, unsigned short**);
int nox_cmd_allow_user(int, int, unsigned short**);
int nox_cmd_allow_ip(int, int, unsigned short**);
int nox_cmd_set_spellpts(int, int, unsigned short**);
int nox_cmd_set_net_debug(int, int, unsigned short**);
int nox_cmd_unset_net_debug(int, int, unsigned short**);
int nox_cmd_show_rank(int, int, unsigned short**);
int nox_cmd_cheat_ability(int, int, unsigned short**);
int nox_cmd_window(int, int, unsigned short**);
int nox_cmd_menu_options(int, int, unsigned short**);
int nox_cmd_reenter(int, int, unsigned short**);
int sub_57A0F0(unsigned short*);
int sub_57A130(unsigned short*);
int sub_57A080(unsigned short*);
int sub_57A0C0(unsigned short*);
int nox_xxx_serverHandleClientConsole_443E90(void*, char, unsigned short*);

static int parsecmd_invalid_count(int which) {
	switch (which) {
	case 0: return nox_cmd_set_sysop(0, 0, 0);
	case 1: return nox_cmd_set_cycle(0, 0, 0);
	case 2: return nox_cmd_set_weapons(0, 0, 0);
	case 3: return nox_cmd_set_staffs(0, 0, 0);
	case 4: return nox_cmd_set_name(0, 2, 0);
	case 5: return nox_cmd_set_mnstrs(0, 0, 0);
	case 6: return nox_cmd_set_spell(0, 0, 0);
	case 7: return nox_cmd_ban(0, 0, 0);
	case 8: return nox_cmd_kick(0, 0, 0);
	case 9: return nox_cmd_set_players(0, 0, 0);
	case 10: return nox_cmd_unmute(0, 0, 0);
	case 11: return nox_cmd_mute(0, 0, 0);
	case 12: return nox_cmd_exec_rul(0, 0, 0);
	case 13: return nox_cmd_offonly1(0, 0, 0);
	case 14: return nox_cmd_offonly2(0, 0, 0);
	case 15: return nox_cmd_show_motd(0, 0, 0);
	case 16: return nox_cmd_show_seq(0, 0, 0);
	case 17: return nox_cmd_set_time(0, 0, 0);
	case 18: return nox_cmd_set_lessons(0, 0, 0);
	default: return -1;
	}
}

static int parsecmd_valid_case(int which) {
	unsigned short* tokens[4] = {
		(unsigned short*)L"set", (unsigned short*)L"missing-player",
		(unsigned short*)L"bogus", (unsigned short*)L"on",
	};
	switch (which) {
	case 0: return nox_cmd_set_cycle(2, 3, tokens);
	case 1: return nox_cmd_set_weapons(2, 3, tokens);
	case 2: return nox_cmd_set_staffs(2, 3, tokens);
	case 3: return nox_cmd_set_spell(2, 4, tokens);
	case 4: return nox_cmd_kick(1, 2, tokens);
	case 5: return nox_cmd_mute(1, 2, tokens);
	case 6: return nox_cmd_unmute(1, 2, tokens);
	case 7: return nox_cmd_exec_rul(1, 2, tokens);
	case 8:
		tokens[1] = (unsigned short*)L"320";
		return nox_cmd_window(1, 2, tokens);
	case 9:
		tokens[1] = (unsigned short*)L"-15";
		return nox_cmd_window(1, 2, tokens);
	case 10:
		tokens[2] = (unsigned short*)L"12";
		return nox_cmd_set_time(2, 3, tokens);
	case 11:
		tokens[2] = (unsigned short*)L"7";
		return nox_cmd_set_lessons(2, 3, tokens);
	case 12:
		tokens[2] = (unsigned short*)L"secret";
		return nox_cmd_set_sysop(2, 3, tokens);
	case 13: return nox_cmd_allow_user(0, 0, tokens);
	case 14: return nox_cmd_allow_ip(0, 0, tokens);
	case 15: return nox_cmd_set_spellpts(0, 0, tokens);
	case 16:
		tokens[2] = (unsigned short*)L"0";
		return nox_cmd_set_players(2, 3, tokens);
	default: return -1;
	}
}
*/
import "C"

func C_parsecmdInvalidCounts() []int {
	out := make([]int, 19)
	for i := range out {
		out[i] = int(C.parsecmd_invalid_count(C.int(i)))
	}
	return out
}

func C_parsecmdNameWithoutArguments() int {
	return int(C.nox_cmd_set_name(3, 3, nil))
}

func C_parsecmdNilPlayerHelpers() [7]int {
	return [7]int{
		int(C.sub_57A0F0(nil)),
		int(C.sub_57A130(nil)),
		int(C.sub_57A080(nil)),
		int(C.sub_57A0C0(nil)),
		int(C.nox_xxx_serverHandleClientConsole_443E90(nil, 0, nil)),
		int(C.sub_57A0F0(nil)),
		int(C.sub_57A080(nil)),
	}
}

func C_parsecmdEngineDebug() (set, unset int) {
	set = int(C.nox_cmd_set_net_debug(0, 2, nil))
	unset = int(C.nox_cmd_unset_net_debug(0, 2, nil))
	return
}

func C_parsecmdShowRank() int {
	return int(C.nox_cmd_show_rank(0, 0, nil))
}

func C_parsecmdSafeOnlineCommands() [5]int {
	return [5]int{
		int(C.nox_cmd_cheat_ability(0, 0, nil)),
		int(C.nox_cmd_cheat_level(0, 0, nil)),
		int(C.nox_cmd_window(0, 0, nil)),
		int(C.nox_cmd_menu_options(0, 0, nil)),
		int(C.nox_cmd_reenter(0, 0, nil)),
	}
}

func C_parsecmdCheatLevelOffline() int {
	return int(C.nox_cmd_cheat_level(0, 0, nil))
}

func C_parsecmdValidCase(which int) int {
	return int(C.parsecmd_valid_case(C.int(which)))
}
