package opennox

/*
#include "defs.h"
#include "nox_net.h"
#include "common__net_list.h"
#include "server__network__sdecode.h"
#include "GAME5.h"
#include "GAME3_2.h"
#include "GAME4_2.h"
#include "GAME5_2.h"
#ifdef _WIN32
#include <windows.h>
#else
#include "windows_compat.h"
#endif
extern unsigned int dword_5d4594_2660032;
extern unsigned int dword_5d4594_814548;
extern unsigned int dword_5d4594_3843632;
extern unsigned int dword_5d4594_2496472;
extern unsigned int dword_5d4594_2496988;
extern unsigned int dword_5d4594_2495920;
extern unsigned long long qword_5d4594_814956;
extern nox_alloc_class* nox_alloc_gQueue_3844300;
extern nox_socket_t nox_xxx_sockLocalBroadcast_2513920;
extern nox_net_struct_t* nox_net_struct_arr[NOX_NET_STRUCT_MAX];
unsigned int nox_client_getServerAddr_43B300();
int nox_client_getServerPort_43B320();
int nox_client_getClientPort_40A420();
int sub_419E60(nox_object_t* a1);
int sub_43AF90(int a1);
int nox_xxx_netClientSend2_4E53C0(int a1, const void* a2, int a3, int a4, int a5);
int  nox_netlist_addToMsgListCli_40EBC0(int ind1, int ind2, unsigned char* buf, int sz);
void* nox_xxx_spriteGetMB_476F80();
int nox_xxx_netSendPacket_4E5030(int a1, const void* a2, signed int a3, int a4, int a5, char a6);
int  nox_xxx_netSendReadPacket_5528B0(unsigned int a1, char a2);
static int nox_xxx_netSendLineMessage_go(nox_object_t* a1, wchar_t* str) {
	return nox_xxx_netSendLineMessage_4D9EB0(a1, str);
}

int nox_xxx_netHandlerDefXxx_553D60(unsigned int a1, char* a2, int a3, void* a4);
int nox_xxx_netHandlerDefYyy_553D70(unsigned int a1, char* a2, int a3, void* a4);

extern float nox_xxx_warriorMaxHealth_587000_312784;
extern float nox_xxx_warriorMaxMana_587000_312788;

extern float nox_xxx_conjurerMaxHealth_587000_312800;
extern float nox_xxx_conjurerMaxMana_587000_312804;

extern float nox_xxx_wizardMaxHealth_587000_312816;
extern float nox_xxx_wizardMaximumMana_587000_312820;
*/
import "C"
import (
	"context"
	"encoding/binary"
	"errors"
	"image"
	"math"
	"net"
	"time"
	"unsafe"

	"github.com/noxworld-dev/nat"
	"github.com/noxworld-dev/opennox-lib/console"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/serial"
)

const (
	NOX_NET_STRUCT_MAX = C.NOX_NET_STRUCT_MAX
)

func init() {
	configBoolPtr("network.xor", "NOX_NET_XOR", true, &noxNetXor)
	configHiddenBoolPtr("debug.network", "NOX_DEBUG_NET", &debugNet)
}

var (
	noxNetXor           bool
	debugNet            bool
	netLog              = log.New("network")
	dword_5D4594_815700 int
)

var (
	noxMapCRC     = 0
	noxServerHost = "localhost"
)

//export nox_xxx_networkLog_print
func nox_xxx_networkLog_print(cstr *C.char) {
	networkLogPrint(GoString(cstr))
}

func networkLogPrint(str string) {
	if !noxflags.HasGame(noxflags.GameFlag3) {
		return
	}
	netLog.Println(str)
	noxConsole.Print(console.ColorGreen, str)
}

//export nox_xxx_netGet_43C750
func nox_xxx_netGet_43C750() C.int { return C.int(dword_5D4594_815700) }

func newNetStruct() (*netStruct, func()) {
	return alloc.New(netStruct{})
}

func asNetStruct(ptr *C.nox_net_struct_t) *netStruct {
	return (*netStruct)(unsafe.Pointer(ptr))
}

func getNetStructByInd(i int) *netStruct {
	if i < 0 || i >= NOX_NET_STRUCT_MAX {
		return nil
	}
	return asNetStruct(C.nox_net_struct_arr[i])
}

func setNetStructByInd(i int, ns *netStruct) {
	if i < 0 || i >= NOX_NET_STRUCT_MAX {
		panic("out of bounds")
	}
	C.nox_net_struct_arr[i] = ns.C()
}

func getFreeNetStruct() int {
	for i := 0; i < NOX_NET_STRUCT_MAX; i++ {
		if C.nox_net_struct_arr[i] == nil {
			return i
		}
	}
	return -1
}

func nox_xxx_netStructByAddr_551E60(ip net.IP, port int) *netStruct {
	for i := 0; i < NOX_NET_STRUCT_MAX; i++ {
		ns := asNetStruct(C.nox_net_struct_arr[i])
		if ns == nil {
			continue
		}
		ip2, port2 := ns.Addr()
		if port == port2 && ip.Equal(ip2) {
			return ns
		}
	}
	return nil
}

type netStruct C.nox_net_struct_t

func (ns *netStruct) C() *C.nox_net_struct_t {
	return (*C.nox_net_struct_t)(unsafe.Pointer(ns))
}

func (ns *netStruct) FreeXxx() {
	if ns == nil {
		return
	}
	if ns.data_3 != nil {
		alloc.Free(unsafe.Pointer(ns.data_3))
	}
	alloc.Free(unsafe.Pointer(ns.data_1_base))
	alloc.Free(unsafe.Pointer(ns.data_2_base))
	C.CloseHandle(C.HANDLE(ns.mutex_yyy))
	C.CloseHandle(C.HANDLE(ns.mutex_xxx))
	alloc.Free(unsafe.Pointer(ns.C()))
}

func (ns *netStruct) Socket() *Socket {
	if ns == nil {
		return nil
	}
	return getSocket(ns.sock)
}

func (ns *netStruct) SetSocket(s *Socket) {
	ns.sock = newSocketHandle(s)
}

func (ns *netStruct) Addr() (net.IP, int) {
	if ns == nil {
		return nil, 0
	}
	return toIPPort(&ns.addr)
}

func (ns *netStruct) SetAddr(ip net.IP, port int) {
	setIPPort(&ns.addr, ip, port)
}

func (ns *netStruct) ID() int {
	if ns == nil {
		return -1
	}
	return int(ns.id)
}

func (ns *netStruct) Data1() []byte {
	if ns == nil {
		return nil
	}
	sz := int(uintptr(unsafe.Pointer(ns.data_1_end)) - uintptr(unsafe.Pointer(ns.data_1_base)))
	if sz < 0 {
		panic("negative size")
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ns.data_1_base)), sz)
}

func (ns *netStruct) Data2() []byte {
	if ns == nil {
		return nil
	}
	sz := int(uintptr(unsafe.Pointer(ns.data_2_end)) - uintptr(unsafe.Pointer(ns.data_2_base)))
	if sz < 0 {
		panic("negative size")
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ns.data_2_base)), sz)
}

func (ns *netStruct) Data2xxx() []byte {
	if ns == nil {
		return nil
	}
	sz := int(uintptr(unsafe.Pointer(ns.data_2_end)) - uintptr(unsafe.Pointer(ns.data_2_xxx)))
	if sz < 0 {
		panic("negative size")
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ns.data_2_xxx)), sz)
}

func clientSetServerHost(host string) {
	netLog.Printf("server host: %s", host)
	noxServerHost = host
}

func clientGetServerHost() string {
	return noxServerHost
}

func clientGetServerPort() int {
	return int(C.nox_client_getServerPort_43B320())
}

func clientGetClientPort() int {
	return int(C.nox_client_getClientPort_40A420())
}

//export nox_client_setServerConnectAddr_435720
func nox_client_setServerConnectAddr_435720(addr *C.char) {
	clientSetServerHost(GoString(addr))
}

//export nox_xxx_cryptXor_56FDD0
func nox_xxx_cryptXor_56FDD0(key C.char, p *C.uchar, n C.int) {
	if p == nil || n == 0 || !noxNetXor {
		return
	}
	buf := unsafe.Slice((*byte)(unsafe.Pointer(p)), int(n))
	netCryptXor(byte(key), buf)
}

//export nox_xxx_cryptXorDst_56FE00
func nox_xxx_cryptXorDst_56FE00(key C.char, src *C.uchar, n C.int, dst *C.uchar) {
	if src == nil || dst == nil || n == 0 {
		return
	}
	sbuf := unsafe.Slice((*byte)(unsafe.Pointer(src)), int(n))
	dbuf := unsafe.Slice((*byte)(unsafe.Pointer(dst)), int(n))
	nox_xxx_cryptXorDst(byte(key), sbuf, dbuf)
}

func nox_xxx_cryptXorDst(key byte, src, dst []byte) {
	if len(src) == 0 || len(dst) == 0 {
		return
	}
	netCryptDst(key, src, dst)
}

func netCryptXor(key byte, p []byte) {
	if len(p) == 0 || !noxNetXor {
		return
	}
	for i := range p {
		p[i] ^= key
	}
}

func netCryptDst(key byte, src, dst []byte) {
	if !noxNetXor {
		copy(dst, src)
		return
	}
	for i := range src {
		dst[i] = key ^ src[i]
	}
}

func nox_xxx_getMapCRC_40A370() int {
	return noxMapCRC
}

//export nox_xxx_setMapCRC_40A360
func nox_xxx_setMapCRC_40A360(crc C.int) {
	if debugNet {
		netLog.Printf("map crc set: %d", int(crc))
	}
	noxMapCRC = int(crc)
}

//export noxOnCliPacketDebug
func noxOnCliPacketDebug(op C.int, data *C.uchar, sz C.int) {
	buf := unsafe.Slice((*byte)(unsafe.Pointer(data)), int(sz))
	if debugNet && sz != 0 {
		op := noxnet.Op(op)
		netLog.Printf("CLIENT: op=%d (%s) [%d:%d]\n%02x %x", int(op), op.String(), int(sz)-1, op.Len(), buf[0], buf[1:])
	}
}

//export noxOnSrvPacketDebug
func noxOnSrvPacketDebug(op C.int, data *C.uchar, sz C.int) {
	buf := unsafe.Slice((*byte)(unsafe.Pointer(data)), int(sz))
	if debugNet && sz != 0 {
		op := noxnet.Op(op)
		netLog.Printf("SERVER: op=%d (%s) [%d:%d]\n%02x %x", int(op), op.String(), int(sz)-1, op.Len(), buf[0], buf[1:])
	}
}

func convSendToServerErr(n int, err error) C.int {
	if err == errLobbyNoSocket {
		return -17
	} else if err != nil {
		return -1
	}
	return C.int(n)
}

func sub_43AF90(v int) {
	C.dword_5d4594_814548 = C.uint(v)
}

func nox_client_createSockAndJoin_43B440() error {
	if dword_587000_87404 == 1 {
		if err := nox_xxx_createSocketLocal(0); err != nil {
			return err
		}
	}
	return nox_client_joinGame()
}

//export nox_client_joinGame_438A90
func nox_client_joinGame_438A90() C.int {
	if err := nox_client_joinGame(); err != nil {
		return convSendToServerErr(0, err)
	}
	return 1
}

func nox_client_joinGame() error {
	endianess := binary.LittleEndian
	buf, freeBuf := alloc.Make([]byte{}, 100)
	defer freeBuf()
	if s, ok := serial.Serial(); ok {
		copy(buf[56:], s)
	}
	wstr := memmap.PtrOff(0x85B3FC, 12204)
	if n := alloc.StrLen((*C.wchar_t)(wstr)); n != 0 {
		copy(buf[4:54], unsafe.Slice((*byte)(wstr), n*2))
	}
	buf[54] = memmap.Uint8(0x85B3FC, 12254)
	buf[55] = memmap.Uint8(0x85B3FC, 12256)
	endianess.PutUint32(buf[80:], NOX_CLIENT_VERS_CODE)
	endianess.PutUint32(buf[84:], uint32(C.dword_5d4594_2660032))

	copy(buf[88:98], GoStringP(memmap.PtrOff(0x85B3FC, 10344)))
	buf[98] = byte(bool2int(!nox_xxx_checkHasSoloMaps()))

	sub_43AF90(3)
	C.qword_5d4594_814956 = C.ulonglong(platformTicks() + 20000)

	addr := nox_client_getServerAddr_43B300()
	port := clientGetServerPort()
	netLog.Printf("join server: %s:%d", addr.String(), port)
	_, err := sendJoinGame(addr, port, buf)
	return err
}

//export sub_5550D0
func sub_5550D0(addr C.int, port C.uint16_t, cdata *C.char) C.int {
	buf := unsafe.Slice((*byte)(unsafe.Pointer(cdata)), 22)
	n, err := sendXXX_5550D0(int2ip(uint32(addr)), int(port), buf)
	return convSendToServerErr(n, err)
}

func sendJoinGame(addr net.IP, port int, data []byte) (int, error) {
	data[0] = 0
	data[1] = 0
	data[2] = 14 // 0x0E
	return sendToServer(addr, port, data)
}

func sendXXX_5550D0(addr net.IP, port int, data []byte) (int, error) {
	data[0] = 0
	data[1] = 0
	data[2] = 17 // 0x11
	return sendToServer(addr, port, data)
}

func (s *Server) nox_xxx_netSendPacket_4E5030(a1 int, buf []byte, a4, a5, a6 int) int {
	b, free := alloc.CloneSlice(buf)
	defer free()
	return int(C.nox_xxx_netSendPacket_4E5030(C.int(a1), unsafe.Pointer(&b[0]), C.int(len(b)), C.int(a4), C.int(a5), C.char(a6)))
}

func (s *Server) nox_xxx_netSendPacket0_4E5420(a1 int, buf []byte, a4, a5 int) int {
	return s.nox_xxx_netSendPacket_4E5030(a1, buf, a4, a5, 0)
}
func (s *Server) nox_xxx_netSendPacket1_4E5390(a1 int, buf []byte, a4, a5 int) int {
	return s.nox_xxx_netSendPacket_4E5030(a1, buf, a4, a5, 1)
}

func (s *Server) nox_xxx_netMsgFadeBegin_4D9800(a1, a2 bool) int {
	var p [3]byte
	p[0] = byte(noxnet.MSG_FADE_BEGIN)
	p[1] = byte(bool2int(a1))
	p[2] = byte(bool2int(a2))
	return s.nox_xxx_netSendPacket1_4E5390(255, p[:], 0, 1)
}

func nox_client_getServerAddr_43B300() net.IP {
	return int2ip(uint32(C.nox_client_getServerAddr_43B300()))
}

func nox_xxx_netClientSendSocial(a1 int, emote byte, a4, a5 int) {
	var buf [2]byte
	buf[0] = byte(noxnet.MSG_SOCIAL)
	buf[1] = emote
	nox_xxx_netClientSend2_4E53C0(a1, buf[:], a4, a5)
}

func nox_xxx_netClientSend2_4E53C0(a1 int, buf []byte, a4, a5 int) {
	p, free := alloc.CloneSlice(buf)
	defer free()
	C.nox_xxx_netClientSend2_4E53C0(C.int(a1), unsafe.Pointer(&p[0]), C.int(len(buf)), C.int(a4), C.int(a5))
}

func (c *Client) clientSendInput(pli int) bool {
	nbuf := ctrlEvent.netBuf
	if len(nbuf) == 0 {
		return true
	}
	var buf [2]byte
	buf[0] = byte(noxnet.MSG_PLAYER_INPUT)
	buf[1] = byte(len(nbuf))
	if !nox_netlist_addToMsgListCli(pli, 0, buf[:2]) {
		return false
	}
	if !nox_netlist_addToMsgListCli(pli, 0, nbuf) {
		return false
	}
	return true
}

func (c *Client) clientSendInputMouse(pli int, mp image.Point) bool {
	sp := nox_xxx_spriteGetMB_476F80()
	if sp != nil {
		mp.Y = sp.Pos().Y
	}
	if mp == c.netPrevMouse {
		return true
	}
	c.netPrevMouse = mp

	var buf [5]byte
	buf[0] = byte(noxnet.MSG_MOUSE)
	binary.LittleEndian.PutUint16(buf[1:], uint16(mp.X))
	binary.LittleEndian.PutUint16(buf[3:], uint16(mp.Y))
	return nox_netlist_addToMsgListCli(pli, 0, buf[:5])
}

type netStructOpt struct {
	field0    uint32
	field1    uint32
	port      int
	data3size int
	field4    int
	datasize  int
	field6    uint32
	field7    uint32
	funcxxx   unsafe.Pointer
	funcyyy   unsafe.Pointer
}

func (s *Server) nox_xxx_netAddPlayerHandler_4DEBC0(port int) (ind, cport int, _ error) {
	narg := &netStructOpt{
		port:      port,
		data3size: 0,
		field4:    s.getServerMaxPlayers(),
		datasize:  2048,
		funcyyy:   C.nox_xxx_netlist_ServRecv_4DEC30,
		funcxxx:   C.nox_xxx_netFn_UpdateStream_4DF630,
	}
	C.nox_xxx_allocNetGQueue_5520B0(200, 1024)
	ind, err := nox_xxx_netInit_554380(narg)
	if err != nil {
		return ind, 0, err
	}
	*memmap.PtrInt32(0x5D4594, 1563148) = int32(ind)
	return ind, narg.port, err
}

func nox_xxx_netPreStructToFull(narg *netStructOpt) (ind int, _ error) {
	if narg == nil {
		return -2, errors.New("empty options")
	}
	if narg.field0 != 0 {
		return -5, errors.New("not empty")
	}
	ind = getFreeNetStruct()
	if ind < 0 {
		return -8, errors.New("no more slots for net structs")
	}
	ns := nox_xxx_makeNewNetStruct(narg)
	setNetStructByInd(ind, ns)
	return ind, nil
}

var netSomePort uint16

//export sub_5545A0
func sub_5545A0() C.short {
	return C.short(netSomePort)
}

func nox_xxx_netInit_554380(narg *netStructOpt) (ind int, _ error) {
	if narg == nil {
		return -2, errors.New("empty options")
	}
	if narg.field0 != 0 {
		return -5, errors.New("not empty")
	}
	if narg.field4 > 128 {
		return -2, errors.New("max limit reached")
	}
	*memmap.PtrUint8(0x973F18, 44216) = 0
	*memmap.PtrUint8(0x973F18, 44232) = 0
	v2 := getFreeNetStruct()
	if v2 < 0 {
		return -8, errors.New("no more slots for net structs")
	}
	ns := nox_xxx_makeNewNetStruct(narg)
	setNetStructByInd(v2, ns)
	ns.Data2()[0] = byte(v2)
	ns.id = -1
	nox_net_init()
	sock := newSocketUDP()
	ns.SetSocket(sock)

	if narg.port < 1024 || narg.port > 0x10000 {
		narg.port = 18590
	}

	netSomePort = uint16(narg.port)
	for {
		err := sock.Bind(nil, narg.port)
		if err == nil {
			break
		} else if !ErrIsInUse(err) {
			return 0, err
		}
		narg.port++
	}
	if ip, err := nat.ExternalIP(context.Background()); err == nil {
		C.dword_5d4594_3843632 = C.uint(ip2int(ip))
		StrCopyP(memmap.PtrOff(0x973F18, 44216), 16, ip.String())
	} else if ips, err := nat.InternalIPs(context.Background()); err == nil && len(ips) != 0 {
		ip = ips[0].IP
		C.dword_5d4594_3843632 = C.uint(ip2int(ip))
		StrCopyP(memmap.PtrOff(0x973F18, 44216), 16, ip.String())
	}
	return v2, nil
}

func toNetStructOpt(arg *C.nox_net_struct_arg_t) *netStructOpt {
	return &netStructOpt{
		field0:    uint32(arg.field_0),
		field1:    uint32(arg.field_1),
		port:      int(arg.port),
		data3size: int(arg.data_3_size),
		field4:    int(arg.field_4),
		datasize:  int(arg.data_size),
		field6:    uint32(arg.field_6),
		field7:    uint32(arg.field_7),
		funcxxx:   unsafe.Pointer(arg.func_xxx),
		funcyyy:   unsafe.Pointer(arg.func_yyy),
	}
}

//export nox_xxx_makeNewNetStruct_553000
func nox_xxx_makeNewNetStruct_553000(arg *C.nox_net_struct_arg_t) *C.nox_net_struct_t {
	narg := toNetStructOpt(arg)
	return nox_xxx_makeNewNetStruct(narg).C()
}

var zeroHandle C.HANDLE

func nox_xxx_makeNewNetStruct(arg *netStructOpt) *netStruct {
	ns, _ := newNetStruct()

	my := C.CreateMutexA(nil, 0, nil)
	if my == zeroHandle {
		panic("cannot create mutex")
	}
	ns.mutex_yyy = unsafe.Pointer(my)

	mx := C.CreateMutexA(nil, 0, nil)
	if mx == zeroHandle {
		panic("cannot create mutex")
	}
	ns.mutex_xxx = unsafe.Pointer(mx)
	if arg.data3size > 0 {
		p, _ := alloc.Make([]byte{}, arg.data3size)
		ns.data_3 = unsafe.Pointer(&p[0])
	}
	if dsz := arg.datasize; dsz > 0 {
		dsz -= dsz % 4
		arg.datasize = dsz
	} else {
		arg.datasize = 1024
	}
	data1, _ := alloc.Make([]byte{}, arg.datasize+2)
	ns.data_1_base = (*C.char)(unsafe.Pointer(&data1[0]))
	ns.data_1_xxx = (*C.char)(unsafe.Pointer(&data1[0]))
	ns.data_1_yyy = (*C.char)(unsafe.Pointer(&data1[0]))
	ns.data_1_end = (*C.char)(unsafe.Add(unsafe.Pointer(&data1[0]), len(data1)))

	data2, _ := alloc.Make([]byte{}, arg.datasize+2)
	data2[0] = 0xff
	ns.data_2_base = (*C.char)(unsafe.Pointer(&data2[0]))
	ns.data_2_xxx = (*C.char)(unsafe.Pointer(&data2[2]))
	ns.data_2_yyy = (*C.char)(unsafe.Pointer(&data2[2]))
	ns.data_2_end = (*C.char)(unsafe.Add(unsafe.Pointer(&data2[0]), len(data2)))

	ns.field_20 = C.uint(arg.field4)
	if arg.funcxxx != nil {
		ns.func_xxx = (*[0]byte)(arg.funcxxx)
	} else {
		ns.func_xxx = (*[0]byte)(C.nox_xxx_netHandlerDefXxx_553D60)
	}
	if arg.funcyyy != nil {
		ns.func_yyy = (*[0]byte)(arg.funcyyy)
	} else {
		ns.func_yyy = (*[0]byte)(C.nox_xxx_netHandlerDefYyy_553D70)
	}
	ns.field_28_1 = -1
	ns.xor_key = 0
	return ns
}

func (s *Server) nox_server_netClose_5546A0(i int) {
	if ns := getNetStructByInd(i); ns != nil {
		_ = ns.Socket().Close()
		ns.SetSocket(nil)
		ns.FreeXxx()
		setNetStructByInd(i, nil)
	}
}

//export nox_xxx_netStructFree_5531C0
func nox_xxx_netStructFree_5531C0(ns *C.nox_net_struct_t) {
	asNetStruct(ns).FreeXxx()
}

//export nox_xxx_netStructReadPackets_5545B0
func nox_xxx_netStructReadPackets_5545B0(ind C.uint) C.int {
	return C.int(noxServer.nox_xxx_netStructReadPackets(int(ind)))
}

func (s *Server) nox_xxx_netStructReadPackets(ind int) int {
	if ind < 0 || ind >= NOX_NET_STRUCT_MAX {
		return -3
	}
	ns := getNetStructByInd(ind)
	if ns == nil {
		return 0
	}
	v4 := ns.ID()
	v1 := ind
	var si, ei int
	if v4 == -1 {
		si, ei = 0, NOX_NET_STRUCT_MAX
		v4 = v1
	} else {
		si, ei = ind, ind+1
		ns2 := getNetStructByInd(v4)
		if ns2 == nil || ns2.id != -1 {
			ns.FreeXxx()
			setNetStructByInd(v1, nil)
			return 0
		}
	}
	for i := si; i < ei; i++ {
		ns2 := getNetStructByInd(i)
		if ns2 == nil || ns2.ID() != v4 {
			continue
		}
		C.nox_xxx_netSendReadPacket_5528B0(C.uint(i), 1)
		var buf [1]byte
		buf[0] = 11
		nox_xxx_netSendSock_552640(i, buf[:], 0)
		C.nox_xxx_netSendReadPacket_5528B0(C.uint(i), 1)
		getNetStructByInd(v4).field_21--
		C.sub_555360(C.uint(v1), 0, 2)
		ns2.FreeXxx()
		setNetStructByInd(i, nil)
	}
	return 0
}

func (s *Server) nox_xxx_netStructReadPackets2_4DEC50(a1 int) int {
	return s.nox_xxx_netStructReadPackets(a1 + 1)
}

//export nox_xxx_netlist_ServRecv_4DEC30
func nox_xxx_netlist_ServRecv_4DEC30(a1 C.int, a2 *C.uchar, a3 C.int, a4 unsafe.Pointer) C.int {
	// should pass the pointer unchanged, otherwise expect bugs!
	nox_xxx_netOnPacketRecvServ_51BAD0_net_sdecode_raw(int(a1-1), unsafe.Slice((*byte)(a2), int(a3)))
	return 1
}

const (
	NOX_NET_SEND_NO_LOCK = 0x1
	NOX_NET_SEND_FLAG2   = 0x2
)

func nox_xxx_netSendSock_552640(id int, buf []byte, flags int) (int, error) {
	ns := getNetStructByInd(id)
	if ns == nil {
		return -3, errors.New("no net struct")
	}
	if len(buf) == 0 {
		return -2, errors.New("empty buffer")
	}
	var (
		idd    int
		ei, si int
	)
	if ns.id == -1 {
		ei = NOX_NET_STRUCT_MAX
		si = 0
		idd = id
	} else {
		si = id
		ei = id + 1
		idd = ns.ID()
	}
	if flags&NOX_NET_SEND_NO_LOCK != 0 {
		n := len(buf)
		for i := si; i < ei; i++ {
			ns2 := getNetStructByInd(i)
			if ns2 != nil && ns2.ID() == idd {
				v12, err := sub_555130(i, buf)
				if v12 == -1 {
					return -1, err
				}
				n = v12
				if flags&NOX_NET_SEND_FLAG2 != 0 {
					nox_xxx_netSend_5552D0(i, byte(v12), true)
				}
			}
		}
		return n, nil
	}
	n := len(buf)
	for i := si; i < ei; i++ {
		ns2 := getNetStructByInd(i)
		if ns2 == nil {
			continue
		}
		if ns2.ID() != idd {
			continue
		}
		d2b := ns2.Data2()
		d2x := ns2.Data2xxx()
		if n+1 > len(d2x) {
			return -7, errors.New("buffer too short")
		}
		v14 := int32(C.WaitForSingleObject(C.HANDLE(ns2.mutex_yyy), 0x3E8))
		if v14 == -1 || v14 == 258 {
			return -16, errors.New("cannot wait for object")
		}
		if flags&NOX_NET_SEND_FLAG2 != 0 {
			copy(d2x[:2], d2b[:2])
			copy(d2x[2:2+n], buf)
			ip, port := ns2.Addr()
			n2, err := nox_xxx_sendto_551F90(ns2.Socket(), d2x[:n+2], ip, port)
			if n2 == -1 {
				return -1, err
			}
			sub_553F40(n+2, 1)
			nox_xxx_netCountData_554030(n+2, i)
			C.ReleaseMutex(C.HANDLE(ns2.mutex_yyy))
			return n2, nil
		}
		copy(d2x[:n], buf)
		ns2.data_2_xxx = (*C.char)(unsafe.Pointer(&d2x[n]))
		if C.ReleaseMutex(C.HANDLE(ns2.mutex_yyy)) == 0 {
			C.ReleaseMutex(C.HANDLE(ns2.mutex_yyy))
		}
	}
	return n, nil
}

func nox_xxx_netCountData_554030(n int, ind int) {
	*memmap.PtrUint32(0x5D4594, 2498024+4*uintptr(ind)) += uint32(n)
}

func sub_553F40(a1, a2 int) {
	*memmap.PtrUint32(0x5D4594, 2495952) += uint32(a1)
	*memmap.PtrUint32(0x5D4594, 2495956) += uint32(a2)
	i := memmap.Uint32(0x5D4594, 2497504)
	j := memmap.Uint32(0x5D4594, 2498020)
	*memmap.PtrUint32(0x5D4594, 2496992+4*uintptr(i)) = uint32(a1)
	*memmap.PtrUint32(0x5D4594, 2497508+4*uintptr(j)) = uint32(a2)
	*memmap.PtrUint32(0x5D4594, 2497504) = uint32(C.dword_5d4594_2496472+1) % 128
	*memmap.PtrUint32(0x5D4594, 2498020) = uint32(C.dword_5d4594_2496988+1) % 128
}

func sub_555130(a1 int, buf []byte) (int, error) {
	if len(buf) > int(memmap.Int32(0x5D4594, 2512884)) {
		return -1, errors.New("buffer too large")
	}
	if len(buf) == 0 {
		return -1, errors.New("empty buffer")
	}
	ns := getNetStructByInd(a1)
	if ns == nil {
		return -3, errors.New("no net struct")
	}
	v5p := alloc.AsClass(unsafe.Pointer(C.nox_alloc_gQueue_3844300)).NewObject()
	if v5p == nil {
		return -1, errors.New("cannot alloc gqueue")
	}
	v5 := unsafe.Slice((*uint32)(v5p), 5)
	v5b := unsafe.Slice((*byte)(v5p), 22+len(buf))
	v5[0] = uint32(uintptr(ns.field_29))
	ns.field_29 = v5p

	v5[3] = 1
	v5[4] = uint32(len(buf) + 2)
	v5b[20] = ns.Data2()[0] | 0x80
	v5b[21] = byte(ns.field_28_0)
	ns.field_28_0++
	copy(v5b[22:], buf)
	return int(v5b[21]), nil
}

var sendXorBuf [4096]byte

func nox_xxx_sendto_551F90(s *Socket, buf []byte, ip net.IP, port int) (int, error) {
	ns := nox_xxx_netStructByAddr_551E60(ip, port)
	if ns == nil {
		return s.SendTo(buf, ip, port)
	}
	if ns.xor_key == 0 {
		return s.SendTo(buf, ip, port)
	}
	dst := sendXorBuf[:len(buf)]
	nox_xxx_cryptXorDst(byte(ns.xor_key), buf, dst)
	return s.SendTo(dst, ip, port)
}

func nox_xxx_netSend_5552D0(ind int, a2 byte, a3 bool) int {
	ns := getNetStructByInd(ind)
	if ns == nil {
		return -3
	}
	for it := unsafe.Pointer(ns.field_29); it != nil; it = *(*unsafe.Pointer)(it) {
		gb := unsafe.Slice((*byte)(it), 22)
		gi := unsafe.Slice((*uint32)(it), 5)
		if a3 {
			if gb[21] == a2 {
				sz := int(gi[4])
				gi[3] = 0
				gi[1] = uint32(C.dword_5d4594_2495920) + 2000
				gb := unsafe.Slice((*byte)(it), 22+sz)
				ip, port := ns.Addr()
				if _, err := nox_xxx_sendto_551F90(ns.Socket(), gb[20:20+sz], ip, port); err != nil {
					netLog.Println(err)
					return 0
				}
			}
		} else if gi[3] != 0 {
			sz := int(gi[4])
			gi[3] = 0
			gi[1] = uint32(C.dword_5d4594_2495920) + 2000
			gb := unsafe.Slice((*byte)(it), 22+sz)
			ip, port := ns.Addr()
			if _, err := nox_xxx_sendto_551F90(ns.Socket(), gb[20:20+sz], ip, port); err != nil {
				netLog.Println(err)
				return 0
			}
		}
	}
	return 0
}

func nox_xxx_netSendClientReady_43C9F0() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_CLIENT_READY)
	nox_xxx_netSendSock_552640(dword_5D4594_815700, data[:], NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2)
	return 1
}

func nox_xxx_netKeepAliveSocket_43CA20() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_KEEP_ALIVE)
	nox_xxx_netSendSock_552640(dword_5D4594_815700, data[:], NOX_NET_SEND_FLAG2)
	return 1
}

func nox_xxx_netRequestMap_43CA50() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_REQUEST_MAP)
	nox_xxx_netSendSock_552640(dword_5D4594_815700, data[:], NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2)
	return 1
}

func nox_xxx_netMapReceived_43CA80() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_RECEIVED_MAP)
	nox_xxx_netSendSock_552640(dword_5D4594_815700, data[:], NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2)
	return 1
}

func nox_xxx_cliSendCancelMap_43CAB0() int {
	id := dword_5D4594_815700
	var data [1]byte
	data[0] = byte(noxnet.MSG_CANCEL_MAP)
	v0, _ := nox_xxx_netSendSock_552640(id, data[:], NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2)
	if nox_xxx_cliWaitServerResponse_5525B0(id, v0, 20, 6) != 0 {
		return 0
	}
	nox_netlist_resetByInd_40ED10(noxMaxPlayers-1, 0)
	return 1
}

func nox_xxx_netSendIncomingClient_43CB00() int {
	id := dword_5D4594_815700
	var data [1]byte
	data[0] = byte(noxnet.MSG_INCOMING_CLIENT)
	v0, _ := nox_xxx_netSendSock_552640(id, data[:], NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2)
	if nox_xxx_cliWaitServerResponse_5525B0(id, v0, 20, 6) != 0 {
		return 0
	}
	nox_netlist_resetByInd_40ED10(noxMaxPlayers-1, 0)
	return 1
}

func nox_xxx_cliSendOutgoingClient_43CB50() int {
	id := dword_5D4594_815700
	var data [1]byte
	data[0] = byte(noxnet.MSG_OUTGOING_CLIENT)
	v0, _ := nox_xxx_netSendSock_552640(id, data[:], NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2)
	if nox_xxx_cliWaitServerResponse_5525B0(id, v0, 20, 6) != 0 {
		return 0
	}
	C.nox_xxx_servNetInitialPackets_552A80(C.uint(id), 3)
	nox_netlist_resetByInd_40ED10(noxMaxPlayers-1, 0)
	return 1
}

func nox_xxx_cliWaitServerResponse_5525B0(a1 int, a2 int, a3 int, a4 byte) int {
	if debugNet {
		netLog.Printf("nox_xxx_cliWaitServerResponse_5525B0: %d, %d, %d, %d\n", a1, a2, a3, a4)
	}
	if a1 >= NOX_NET_STRUCT_MAX {
		return -3
	}
	ns := getNetStructByInd(a1)
	if ns == nil {
		return -3
	}

	if int(ns.field_28_1) >= a2 {
		return 0
	}
	for v6 := 0; v6 <= 20*a3; v6++ {
		platform.Sleep(50 * time.Millisecond)
		C.nox_xxx_servNetInitialPackets_552A80(C.uint(a1), C.char(a4|1))
		C.nox_xxx_netMaybeSendAll_552460()
		if int(ns.field_28_1) >= a2 {
			return 0
		}
		// FIXME(awesie)
		return 0
	}
	return -23
}

func nox_xxx_netInformTextMsg_4DA0F0(pid int, code byte, ind int) bool {
	if pid < 0 {
		return false
	}
	var buf [6]byte
	buf[0] = byte(noxnet.MSG_INFORM)
	buf[1] = code
	switch code {
	case 0, 1, 2, 12, 13, 16, 20, 21:
		binary.LittleEndian.PutUint32(buf[2:], uint32(ind))
		return nox_netlist_addToMsgListCli(pid, 1, buf[:6])
	case 17:
		return nox_netlist_addToMsgListCli(pid, 1, buf[:2])
	default:
		return true
	}
}

func nox_xxx_netReportSpellStat_4D9630(a1 int, a2 spell.ID, a3 byte) bool {
	var buf [6]byte
	buf[0] = byte(noxnet.MSG_REPORT_SPELL_STAT)
	binary.LittleEndian.PutUint32(buf[1:], uint32(a2))
	buf[5] = a3
	return noxServer.nox_xxx_netSendPacket0_4E5420(a1, buf[:], 0, 1) != 0
}

func nox_xxx_netSendLineMessage_4D9EB0(u *Unit, s string) bool {
	_ = noxnet.MSG_TEXT_MESSAGE
	cstr, free := CWString(s)
	defer free()
	return C.nox_xxx_netSendLineMessage_go(u.CObj(), cstr) != 0
}

func nox_xxx_netSendPointFx_522FF0(fx noxnet.Op, pos types.Pointf) bool {
	var buf [5]byte
	buf[0] = byte(fx)
	binary.LittleEndian.PutUint16(buf[1:], uint16(int(pos.X)))
	binary.LittleEndian.PutUint16(buf[3:], uint16(int(pos.Y)))
	return nox_xxx_netSendFxAllCli_523030(pos, buf[:5])
}

func nox_xxx_netSendRayFx_5232F0(fx noxnet.Op, p1, p2 image.Point) bool {
	var buf [9]byte
	buf[0] = byte(fx)
	binary.LittleEndian.PutUint16(buf[1:], uint16(p1.X))
	binary.LittleEndian.PutUint16(buf[3:], uint16(p1.Y))
	binary.LittleEndian.PutUint16(buf[5:], uint16(p2.X))
	binary.LittleEndian.PutUint16(buf[7:], uint16(p2.Y))
	return nox_xxx_servCode_523340(p1, p2, buf[:9])
}

func nox_xxx_netSparkExplosionFx_5231B0(pos types.Pointf, a2 byte) bool {
	var buf [6]byte
	buf[0] = byte(noxnet.MSG_FX_SPARK_EXPLOSION)
	binary.LittleEndian.PutUint16(buf[1:], uint16(pos.X))
	binary.LittleEndian.PutUint16(buf[3:], uint16(pos.Y))
	buf[5] = a2
	return nox_xxx_netSendFxAllCli_523030(pos, buf[:6])
}

func nox_xxx_earthquakeSend_4D9110(pos types.Pointf, a2 int) {
	cpos, pfree := alloc.Make([]float32{}, 2)
	defer pfree()
	cpos[0] = pos.X
	cpos[1] = pos.Y

	C.nox_xxx_earthquakeSend_4D9110((*C.float)(unsafe.Pointer(&cpos[0])), C.int(a2))
}

func nox_xxx_netSendFxGreenBolt_523790(p1, p2 image.Point, a2 int) bool {
	var buf [11]byte
	buf[0] = byte(noxnet.MSG_FX_GREEN_BOLT)
	binary.LittleEndian.PutUint16(buf[1:], uint16(p1.X))
	binary.LittleEndian.PutUint16(buf[3:], uint16(p1.Y))
	binary.LittleEndian.PutUint16(buf[5:], uint16(p2.X))
	binary.LittleEndian.PutUint16(buf[7:], uint16(p2.Y))
	binary.LittleEndian.PutUint16(buf[9:], uint16(a2))
	pos := types.Pointf{
		X: float32(p1.X) + float32(p2.X-p1.X)*0.5,
		Y: float32(p1.Y) + float32(p2.Y-p1.Y)*0.5,
	}
	return nox_xxx_netSendFxAllCli_523030(pos, buf[:11])
}

func nox_xxx_netSendVampFx_523270(fx noxnet.Op, p1, p2 image.Point, a3 int) bool {
	var buf [11]byte
	buf[0] = byte(fx)
	binary.LittleEndian.PutUint16(buf[1:], uint16(p1.X))
	binary.LittleEndian.PutUint16(buf[3:], uint16(p1.Y))
	binary.LittleEndian.PutUint16(buf[5:], uint16(p2.X))
	binary.LittleEndian.PutUint16(buf[7:], uint16(p2.Y))
	binary.LittleEndian.PutUint16(buf[9:], uint16(a3))
	pos := types.Pointf{
		X: float32(p2.X),
		Y: float32(p2.Y),
	}
	return nox_xxx_netSendFxAllCli_523030(pos, buf[:11])
}

func nox_xxx_netSendFxAllCli_523030(pos types.Pointf, data []byte) bool {
	cdata, dfree := alloc.Make([]byte{}, len(data))
	defer dfree()
	copy(cdata, data)

	cpos, pfree := alloc.Make([]float32{}, 2)
	defer pfree()
	cpos[0] = pos.X
	cpos[1] = pos.Y

	return C.nox_xxx_netSendFxAllCli_523030((*C.float2)(unsafe.Pointer(&cpos[0])), unsafe.Pointer(&cdata[0]), C.int(len(data))) != 0
}

func nox_xxx_servCode_523340(p1, p2 image.Point, data []byte) bool {
	cdata, dfree := alloc.Make([]byte{}, len(data))
	defer dfree()
	copy(cdata, data)

	cpos, pfree := alloc.Make([]int32{}, 4)
	defer pfree()
	cpos[0] = int32(p1.X)
	cpos[1] = int32(p1.Y)
	cpos[2] = int32(p2.X)
	cpos[3] = int32(p2.Y)

	return C.nox_xxx_servCode_523340((*C.int)(unsafe.Pointer(&cpos[0])), unsafe.Pointer(&cdata[0]), C.int(len(data))) != 0
}

func nox_xxx_netReportLesson_4D8EF0(u *Unit) {
	var buf [11]byte
	buf[0] = byte(noxnet.MSG_REPORT_LESSON)
	pl := u.ControllingPlayer()
	binary.LittleEndian.PutUint16(buf[1:], uint16(u.net_code))
	binary.LittleEndian.PutUint32(buf[3:], uint32(pl.lessons))
	binary.LittleEndian.PutUint32(buf[7:], uint32(pl.field_2140))
	noxServer.nox_xxx_netSendPacket1_4E5390(255, buf[:11], 0, 1)
}

func nox_xxx_netScriptMessageKill_4D9760(u *Unit) {
	if !u.Class().Has(object.ClassPlayer) {
		return
	}
	pl := u.ControllingPlayer()
	var buf [1]byte
	buf[0] = byte(noxnet.MSG_MESSAGES_KILL)
	noxServer.nox_xxx_netSendPacket0_4E5420(pl.Index(), buf[:1], 0, 1)
}

func nox_xxx_netKillChat_528D00(u *Unit) {
	var buf [3]byte
	buf[0] = byte(noxnet.MSG_CHAT_KILL)
	binary.LittleEndian.PutUint16(buf[1:], uint16(noxServer.getUnitNetCode(u)))
	for _, pl := range noxServer.getPlayers() {
		u := pl.UnitC()
		if u == nil {
			continue
		}
		nox_netlist_addToMsgListCli(pl.Index(), 1, buf[:3])
	}
}

func netSendGauntlet() {
	var buf [2]byte
	buf[0] = byte(noxnet.MSG_GAUNTLET)
	buf[1] = 27
	nox_xxx_netClientSend2_4E53C0(noxMaxPlayers-1, buf[:2], 0, 0)
}

func nox_xxx_sendGauntlet_4DCF80(ind int, v byte) {
	var buf [3]byte
	buf[0] = byte(noxnet.MSG_GAUNTLET)
	buf[1] = 28
	buf[2] = v
	noxServer.nox_xxx_netSendPacket1_4E5390(ind, buf[:3], 0, 0)
}

//export nox_xxx_netStatsMultiplier_4D9C20
func nox_xxx_netStatsMultiplier_4D9C20(a1p *nox_object_t) C.int {
	u := asUnitC(a1p)
	if u == nil {
		return 0
	}
	pl := u.ControllingPlayer()
	var buf [17]byte
	buf[0] = byte(noxnet.MSG_STAT_MULTIPLIERS)
	switch pl.PlayerClass() {
	default:
		return 0
	case player.Warrior:
		binary.LittleEndian.PutUint32(buf[1:], math.Float32bits(float32(C.nox_xxx_warriorMaxHealth_587000_312784)))
		binary.LittleEndian.PutUint32(buf[5:], math.Float32bits(float32(C.nox_xxx_warriorMaxMana_587000_312788)))
		binary.LittleEndian.PutUint32(buf[9:], math.Float32bits(noxServer.players.mult.warrior.strength))
		binary.LittleEndian.PutUint32(buf[13:], math.Float32bits(noxServer.players.mult.warrior.speed))
	case player.Wizard:
		binary.LittleEndian.PutUint32(buf[1:], math.Float32bits(float32(C.nox_xxx_wizardMaxHealth_587000_312816)))
		binary.LittleEndian.PutUint32(buf[5:], math.Float32bits(float32(C.nox_xxx_wizardMaximumMana_587000_312820)))
		binary.LittleEndian.PutUint32(buf[9:], math.Float32bits(noxServer.players.mult.wizard.strength))
		binary.LittleEndian.PutUint32(buf[13:], math.Float32bits(noxServer.players.mult.wizard.speed))
	case player.Conjurer:
		binary.LittleEndian.PutUint32(buf[1:], math.Float32bits(float32(C.nox_xxx_conjurerMaxHealth_587000_312800)))
		binary.LittleEndian.PutUint32(buf[5:], math.Float32bits(float32(C.nox_xxx_conjurerMaxMana_587000_312804)))
		binary.LittleEndian.PutUint32(buf[9:], math.Float32bits(noxServer.players.mult.conjurer.strength))
		binary.LittleEndian.PutUint32(buf[13:], math.Float32bits(noxServer.players.mult.conjurer.speed))
	}
	return C.int(noxServer.nox_xxx_netSendPacket0_4E5420(pl.Index(), buf[:17], 0, 1))
}

func netSendServerQuit() {
	var buf [1]byte
	buf[0] = byte(noxnet.MSG_SERVER_QUIT)
	noxServer.nox_xxx_netSendPacket0_4E5420(159, buf[:1], 0, 1)
}

func nox_xxx_netSendBallStatus_4D95F0(a1 int, a2 byte, a3 uint16) int {
	var buf [4]byte
	buf[0] = byte(noxnet.MSG_REPORT_BALL_STATUS)
	buf[1] = a2
	binary.LittleEndian.PutUint16(buf[2:], a3)
	return noxServer.nox_xxx_netSendPacket1_4E5390(a1, buf[:4], 0, 1)
}

func (s *Server) netPrintLineToAll(id strman.ID) { // nox_xxx_netPrintLineToAll_4DA390
	for _, u := range s.getPlayerUnits() {
		nox_xxx_netPriMsgToPlayer_4DA2C0(u, id, 0)
	}
}

func nox_xxx_netPriMsgToPlayer_4DA2C0(u *Unit, id strman.ID, a3 byte) {
	var buf [52]byte
	if u == nil || !u.Class().Has(object.ClassPlayer) || id == "" || len(id) > len(buf)-4 || C.sub_419E60(u.CObj()) != 0 {
		return
	}
	buf[0] = byte(noxnet.MSG_INFORM)
	buf[1] = 15
	buf[2] = a3
	n := copy(buf[3:len(buf)-1], string(id))
	buf[3+n] = 0
	nox_netlist_addToMsgListCli(u.ControllingPlayer().Index(), 1, buf[:n+4])
}

func nox_xxx_netSendBySock_40EE10(a1 int, a2 int, a3 int) {
	v3 := nox_netlist_copyPacketList(a2, a3)
	if len(v3) != 0 {
		nox_xxx_netSendSock_552640(a1, v3, 0)
		C.nox_xxx_netSendReadPacket_5528B0(C.uint(a1), 1)
	}
}
