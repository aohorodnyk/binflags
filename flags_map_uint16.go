package gobitflags

func HasFlagMapUint16(flags map[uint64]uint16, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	flagBits := flags[idx]

	conv := uint16(1 << bit)
	return flagBits&conv == conv
}

func SetFlagMapUint16(flags map[uint64]uint16, flag uint64, set bool) {
	if flags == nil {
		return
	}

	if HasFlagMapUint16(flags, flag) == set {
		return
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	flagBits, ok := flags[idx]

	conv := uint16(1 << bit)
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}
}
