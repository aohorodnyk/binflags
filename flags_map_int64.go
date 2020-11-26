package gobitflags

func HasFlagMapInt64(flags map[uint64]int64, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	flagBits := flags[idx]

	conv := int64(1 << bit)
	return flagBits&conv == conv
}

func SetFlagMapInt64(flags map[uint64]int64, flag uint64, set bool) {
	if flags == nil {
		return
	}

	if HasFlagMapInt64(flags, flag) == set {
		return
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	flagBits, ok := flags[idx]

	conv := int64(1 << bit)
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}
}
