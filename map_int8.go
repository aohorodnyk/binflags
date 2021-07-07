package binflags

func HasFlagMapInt8(flags map[uint64]int8, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt8)

	return HasFlagInt8(flags[idx], bit)
}

func SetFlagMapInt8(flags map[uint64]int8, flag uint64, set bool) error {
	if flags == nil {
		return ErrorFlagsMapNil(ErrorMsgFlagsMapNil)
	}

	if HasFlagMapInt8(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	bits, ok := flags[idx]

	ret, err := SetFlagInt8(bits, bit, set)
	if err != nil {
		return err
	}

	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}

	return nil
}
