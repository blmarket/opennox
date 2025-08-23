package opennox

import (
	"encoding/binary"
	"fmt"
	"math"
	"sort"
	"strings"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/ifs"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/audio"
	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

var (
	nox_enable_audio = 1
)

var (
	audioTimer93944 ail.Timer = math.MaxUint32
	audioDev        ail.Driver
)

func sub_43EFD0(a1 unsafe.Pointer) int {
	s := *(**legacy.AudioSample)(unsafe.Add(a1, 272))
	s.Smp.End()
	if s.Flag7 == 0 {
		ptr := s.Field1
		fptr := (*unsafe.Pointer)(unsafe.Add(ptr, 284))
		ccall.CallVoidPtr(*fptr, ptr)
		s.Flag7 = 1
	}
	return 0
}

func sub_43EC30(a1 unsafe.Pointer) int {
	smp := audioDev.AllocateSample()
	s, _ := alloc.New(legacy.AudioSample{})
	*(*unsafe.Pointer)(unsafe.Add(a1, 272)) = unsafe.Pointer(s)
	s.Dev = audioDev
	s.Field1 = a1
	s.Smp = smp
	b1, _ := alloc.Make([]byte{}, 16384)
	b2, _ := alloc.Make([]byte{}, 16384)
	s.Data1 = &b1[0]
	s.Data2 = &b2[0]
	if smp == 0 {
		return -2147221504 // 0x80040000
	}
	smp.SetUserData(s)
	return 0
}

func sub_43ECB0(a1 unsafe.Pointer) int {
	p := *(**legacy.AudioSample)(unsafe.Add(a1, 272))
	if p.Smp != 0 {
		p.Smp.Release()
	}
	if p2 := p.Data1; p2 != nil {
		alloc.Free(p2)
	}
	if p2 := p.Data2; p2 != nil {
		alloc.Free(p2)
	}
	alloc.Free(p)
	return 0
}

func sub_43E940(a1 unsafe.Pointer) int {
	ail.Startup()
	audioTimer93944 = ail.RegisterTimer(func(u uint32) {
		legacy.AudioModule.Sub_486EF0()
		legacy.MusicModule.Sub_43D2D0()
		(*timer.TimerGroup)(legacy.Get_dword_587000_127004()).ClearUpdated()
	})
	if audioTimer93944 == math.MaxUint32 {
		return -2147221504 // 0x80040000
	}
	audioTimer93944.SetFrequency(30)
	audioTimer93944.Start()
	legacy.Sub_42EBB0(1, legacy.Get_sub_43E910(), 0, "Audio")
	legacy.Sub_42EBB0(2, legacy.Get_sub_43E8E0(), 0, "Audio")
	*(*uint32)(unsafe.Add(a1, 20)) = 1
	return 0
}

func sub_43E9F0() {
	if audioTimer93944 != math.MaxUint32 {
		audioTimer93944.Stop()
		audioTimer93944.Release()
		audioTimer93944 = math.MaxUint32
	}
	ail.Shutdown()
}

func noxAudioServeT(a1 int) {
	noxAudioServe()
}

func noxAudioServe() { // sub_43F1A0
	if audioDev != 0 {
		ail.Serve()
	}
}

func sub_43EA20(a1 unsafe.Pointer) int {
	//char v2[16]; // [esp+4h] [ebp-10h]

	v2 := sub_43EA90(unsafe.Add(a1, 60))
	audioDev = sub_43EAD0(v2)
	if audioDev == 0 {
		return -2147221504 // 0x80040000
	}
	*memmap.PtrInt32(0x587000, 93948) = 0
	*(*uint32)(unsafe.Add(a1, 196)) = 24
	return 0
}

func sub_43EA90(a2 unsafe.Pointer) *audio.AudioStruct1 {
	return &audio.AudioStruct1{
		Field0:  1,
		Field2:  *(*uint16)(unsafe.Add(a2, 12)),
		Field4:  *(*uint32)(unsafe.Add(a2, 8)),
		Field8:  *(*uint32)(unsafe.Add(a2, 20)),
		Field12: *(*uint16)(unsafe.Add(a2, 16)) * *(*uint16)(unsafe.Add(a2, 12)),
		Field14: 8 * *(*uint16)(unsafe.Add(a2, 16)),
	}
}

func sub_43EAD0(a1 *audio.AudioStruct1) ail.Driver {
	if dr := ail.WaveOutOpen(); dr != 0 {
		return dr
	}
	_ = ail.LastError()
	return 0
}

func sub_43EC10() int {
	if audioDev != 0 {
		audioDev.Close()
		audioDev = 0
	}
	return 0
}

func sub_43F130() ail.Driver {
	return audioDev
}

func sub_43ED00(a1p unsafe.Pointer) int {
	a1 := unsafe.Slice((*uint32)(a1p), 73)
	s := *(**legacy.AudioSample)(unsafe.Pointer(&a1[68]))
	smp := s.Smp
	smp.Init()
	s.Field4 = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(uintptr(a1[72])), 20))
	if s.Field4 == nil {
		s.Field4 = unsafe.Pointer(uintptr(a1[33]) + 60)
	}
	v3 := int32(legacy.Sub_43F0E0(s.Field4))
	smp.SetType(v3, 0)
	if *(*uint32)(unsafe.Add(s.Field4, 4)) == 2 {
		smp.SetADPCMBlockSize(*(*uint32)(unsafe.Add(s.Field4, 24)))
	}
	sub_43F060(unsafe.Pointer(&a1[0]))
	s.Flag7 = 0
	s.Field3 = 0
	smp.RegisterEOBCallback(func() {
		v := smp.UserData().(*legacy.AudioSample)
		legacy.Sub_43EE00(v)
	})
	smp.RegisterEOSCallback(func() {
		legacy.Sub_43EDB0(smp)
	})
	legacy.Sub_43EE00(s)
	return 0
}

func sub_43F060(a1p unsafe.Pointer) int {
	a1 := unsafe.Slice((*uint32)(unsafe.Pointer(a1p)), 69)
	s := *(**legacy.AudioSample)(unsafe.Pointer(&a1[68]))
	smp := s.Smp
	smp.SetVolume(int((127 * (a1[45] >> 16)) >> 14))
	smp.SetPan(int((127 * (a1[61] >> 16)) >> 14))
	p2 := s.Field4
	v2 := legacy.Sub_486640(unsafe.Pointer(&a1[44]), int(*(*uint32)(unsafe.Add(p2, 8))))
	smp.SetPlaybackRate(v2)
	return 0
}

func nox_xxx_parseSoundSetBin_424170(path string) error {
	if legacy.Nox_xxx_parseSoundSetBin_424170(path) == 0 {
		return fmt.Errorf("failed to load sounds")
	}
	return nil
}

func nox_audio_initall(a3 int) int {
	if nox_enable_audio == 0 {
		return 1
	}
	if a3 != 0 {
		legacy.AudioModule.Sub_486F30()
		if sub_4311F0() != 0 {
			legacy.Set_dword_587000_81128(&legacy.Get_dword_5d4594_805984().TimerGroup_22)
			legacy.AudioModule.Externs.Dword_5d4594_805980 = sub_4866F0("audio", "audio")
		}
	}
	(*timer.TimerGroup)(memmap.PtrOff(0x5D4594, 805884)).Init()
	(*timer.TimerGroup)(legacy.Get_dword_587000_93164()).Init()
	(*timer.TimerGroup)(legacy.Get_dword_587000_122852()).Init()
	(*timer.TimerGroup)(legacy.Get_dword_587000_127004()).Init()
	legacy.Dialogs.Nox_xxx_WorkerHurt_44D810()
	legacy.MusicModule.Init()
	legacy.AudioModule.Sub_451850(legacy.Get_dword_5d4594_805984(), legacy.AudioModule.Externs.Dword_5d4594_805980)
	v1 := configGetVolume(VolumeMusic)
	if v1 == 0 {
		legacy.Sub_43DC00()
	}
	(*timer.Timer)(legacy.Get_dword_587000_93164()).SetRaw(uint32(v1))
	v2 := configGetVolume(VolumeDialog)
	if v2 == 0 {
		legacy.Sub_44D960()
	}
	(*timer.Timer)(legacy.Get_dword_587000_122852()).SetRaw(uint32(v2))
	v3 := configGetVolume(VolumeFX)
	if v3 == 0 {
		legacy.Sub_453050()
	}
	(*timer.Timer)(legacy.Get_dword_587000_127004()).SetRaw(uint32(v3))
	return 1
}

func sub_4311F0() int {
	legacy.AudioModule.Sub_486FA0(*memmap.PtrT[*audio.Struct587000_94032](0x587000, 94032))
	v2a, free := alloc.Make([][7]uint32{}, 1)
	defer free()
	v2 := v2a[0]
	v2[2] = 22050
	v2[1] = 0
	v2[3] = 2
	v2[4] = 2
	v2[0] = 4
	legacy.AudioModule.Sub_487D00(&v2)
	v0 := legacy.AudioModule.Sub_487150(int32(-1), &v2)
	legacy.Set_dword_5d4594_805984(v0)
	return bool2int(v0 != nil && legacy.AudioModule.Sub_487790(v0, 16) == 16)
}

func freeAudioStructXxx(p *audio.AudioStructXxx) {
	p.Free(legacy.AudioModule)
}

func sub_4866F0(path1 string, path2 string) *audio.AudioStructXxx {
	idxPath := path1
	if i := strings.LastIndexByte(idxPath, '.'); i >= 0 {
		idxPath = idxPath[:i]
	}
	idxPath += ".idx"
	fi, err := ifs.Open(idxPath)
	if err != nil {
		return nil
	}
	bf1 := binfile.NewFile(fi)
	cf1 := legacy.NewFileHandle(bf1)
	defer legacy.Nox_fs_close(cf1)

	bagPath := path1
	if i := strings.LastIndexByte(bagPath, '.'); i >= 0 {
		bagPath = bagPath[:i]
	}
	bagPath += ".bag"
	fb, err := ifs.Open(bagPath)
	if err != nil {
		return nil
	}

	p, _ := alloc.New(audio.AudioStructXxx{})
	p.Bagfile268 = audio.FILE(legacy.NewFileHandle(binfile.NewFile(fb)))
	if p.Bagfile268 == nil {
		freeAudioStructXxx(p)
		return nil
	}
	var hdr [12]byte
	if _, err := bf1.Read(hdr[:12]); err != nil {
		freeAudioStructXxx(p)
		return nil
	}
	_ = binary.LittleEndian.Uint32(hdr[0:])
	vers := binary.LittleEndian.Uint32(hdr[4:])
	p.Size4 = uint(binary.LittleEndian.Uint32(hdr[8:]))
	var arr []audio.AudioStructYyy
	if p.Size4 > 0 {
		arr, _ = alloc.Make([]audio.AudioStructYyy{}, p.Size4)
		p.Arr0 = &arr[0]
		var buf [36]byte
		if vers != 1 {
			for i := uint(0); i < p.Size4; i++ {
				if _, err := bf1.Read(buf[:36]); err != nil {
					freeAudioStructXxx(p)
					return nil
				}
				copy(arr[i].Field0[:], buf[:32])
				arr[i].Field32 = binary.LittleEndian.Uint32(buf[32:])
			}
		} else {
			for i := uint(0); i < p.Size4; i++ {
				if _, err := bf1.Read(buf[:32]); err != nil {
					freeAudioStructXxx(p)
					return nil
				}
				copy(arr[i].Field0[:], buf[:32])
				arr[i].Field32 = 0
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		s1, s2 := alloc.GoStringS(arr[i].Field0[:]), alloc.GoStringS(arr[j].Field0[:])
		s1, s2 = strings.ToLower(s1), strings.ToLower(s2)
		return s1 < s2
	})
	p.Field276 = 0
	if path2 != "" {
		alloc.StrCopyZero(p.Path2_8[:], path2)
		if i := alloc.StrLenS(p.Path2_8[:]); p.Path2_8[i] == '\\' {
			p.Path2_8[i] = 0
		}
		var find legacy.WIN32_FIND_DATAA

		if h := legacy.FindFirstFileA(&p.Path2_8[0], &find); int(h) != -1 {
			if find.FileAttributes&0x10 != 0 {
				p.Field276 = 1
			} else {
				for legacy.FindNextFileA(h, &find) != 0 {
					if find.FileAttributes&0x10 != 0 {
						p.Field276 = 1
						break
					}
				}
			}
			legacy.FindClose(h)
		}
		// TODO: strlen()-1 ?
		if i := alloc.StrLenS(p.Path2_8[:]); p.Path2_8[i] != '\\' {
			p.Path2_8[i] = '\\'
			p.Path2_8[i+1] = 0
		}
	}
	return p
}
