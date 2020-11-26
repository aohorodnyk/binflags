package gobitflags

func HasFlagMapInt16(flags map[uint64]int16, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	flagBits := flags[idx]

	conv := int16(1 << bit)
	return flagBits&conv == conv
}

func SetFlagMapInt16(flags map[uint64]int16, flag uint64, set bool) {
	if flags == nil {
		return
	}

	if HasFlagMapInt16(flags, flag) == set {
		return
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	flagBits, ok := flags[idx]

	conv := int16(1 << bit)
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}
}
