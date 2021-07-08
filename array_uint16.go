package binflags

func HasFlagArrayUint16(flags []uint16, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagUint16(flags[idx], bit)
}

func SetFlagArrayUint16(flags []uint16, flag uint64, set bool) error {
	if flags == nil {
		return ErrorFlagsArrayNil(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayUint16(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	if len(flags) <= int(idx) {
		return ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	var err error
	flags[idx], err = SetFlagUint16(flags[idx], bit, set)

	return err
}
