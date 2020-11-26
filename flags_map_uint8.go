package gobitflags

func HasFlagMapByte(flags map[uint64]uint8, flag uint64) bool {
	return HasFlagMapUint8(flags, flag)
}

func SetFlagMapByte(flags map[uint64]uint8, flag uint64, val bool) {
	SetFlagMapUint8(flags, flag, val)
}

func HasFlagMapUint8(flags map[uint64]uint8, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	flagBits := flags[idx]

	conv := uint8(1 << bit)
	return flagBits&conv == conv
}

func SetFlagMapUint8(flags map[uint64]uint8, flag uint64, set bool) {
	if flags == nil {
		return
	}

	if HasFlagMapUint8(flags, flag) == set {
		return
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	flagBits, ok := flags[idx]

	conv := uint8(1 << bit)
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}
}
