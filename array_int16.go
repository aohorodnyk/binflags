package binflags

func HasFlagArrayInt16(flags []int16, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagInt16(flags[idx], bit)
}

func SetFlagArrayInt16(flags []int16, flag uint64, set bool) error {
	if flags == nil {
		return ErrorFlagsArrayNil(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayInt16(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	if len(flags) <= int(idx) {
		return ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	var err error
	flags[idx], err = SetFlagInt16(flags[idx], bit, set)

	return err
}
