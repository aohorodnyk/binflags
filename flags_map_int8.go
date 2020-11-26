package gobitflags

func HasFlagMapInt8(flags map[uint64]int8, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	flagBits := flags[idx]

	conv := int8(1 << bit)
	return flagBits&conv == conv
}

func SetFlagMapInt8(flags map[uint64]int8, flag uint64, set bool) {
	if flags == nil {
		return
	}

	if HasFlagMapInt8(flags, flag) == set {
		return
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	flagBits, ok := flags[idx]

	conv := int8(1 << bit)
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}
}
