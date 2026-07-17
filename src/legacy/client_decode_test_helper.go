package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int nox_xxx_netOnPacketRecvCli_48EA70_switch(int a1, int op, unsigned char* data, int sz);
*/
import "C"

func C_clientDecodeZeroPacket(op byte) int {
	const packetSize = 512
	data := C.calloc(1, packetSize)
	defer C.free(data)
	*(*byte)(data) = op
	return int(C.nox_xxx_netOnPacketRecvCli_48EA70_switch(
		0, C.int(op), (*C.uchar)(data), packetSize,
	))
}
