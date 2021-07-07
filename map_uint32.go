package binflags

func HasFlagMapUint32(flags map[uint64]uint32, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt32)

	return HasFlagUint32(flags[idx], bit)
}

func SetFlagMapUint32(flags map[uint64]uint32, flag uint64, set bool) error {
	if flags == nil {
		return ErrorFlagsMapNil(ErrorMsgFlagsMapNil)
	}

	if HasFlagMapUint32(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt32)
	bits, ok := flags[idx]

	ret, err := SetFlagUint32(bits, bit, set)
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
