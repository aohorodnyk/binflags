package binflags

func HasFlagMapInt16(flags map[uint64]int16, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt16)

	return HasFlagInt16(flags[idx], bit)
}

func SetFlagMapInt16(flags map[uint64]int16, flag uint64, set bool) error {
	if flags == nil {
		return ErrorFlagsMapNil(ErrorMsgFlagsMapNil)
	}

	if HasFlagMapInt16(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	bits, ok := flags[idx]

	ret, err := SetFlagInt16(bits, bit, set)
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
