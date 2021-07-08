package binflags

func HasFlagArrayUint64(flags []uint64, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagUint64(flags[idx], bit)
}

func SetFlagArrayUint64(flags []uint64, flag uint64, set bool) error {
	if flags == nil {
		return ErrorFlagsArrayNil(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayUint64(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	if len(flags) <= int(idx) {
		return ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	var err error
	flags[idx], err = SetFlagUint64(flags[idx], bit, set)

	return err
}
