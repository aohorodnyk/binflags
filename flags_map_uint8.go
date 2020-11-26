package gobitflags

func HasFlagMapByte(flags map[uint64]uint8, flag uint64) bool {
	return HasFlagMapUint8(flags, flag)
}

func SetFlagMapByte(flags map[uint64]uint8, flag uint64, val bool) {
	SetFlagMapUint8(flags, flag, val)
}

func HasFlagMapUint8(flags map[uint64]uint8, flag uint64) bool {
	idx, bit := FlagMap(flag, FlagMaxInt8)
	flagBits := flags[idx]

	var conv uint8 = 1 << bit
	return flagBits&conv == conv
}

func SetFlagMapUint8(flags map[uint64]uint8, flag uint64, val bool) {
	if HasFlagMapUint8(flags, flag) == val {
		return
	}

	idx, bit := FlagMap(flag, FlagMaxInt8)
	flagBits, ok := flags[idx]

	var conv uint8 = 1 << bit
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}
}
