package binflags

func HasFlagUint64(flags uint64, flag uint8) bool {
	if flag > FlagMaxInt64 {
		return false
	}

	conv := uint64(1 << flag)

	return flags&conv == conv
}

func SetFlagUint64(flags uint64, flag uint8, set bool) (uint64, error) {
	if flag > FlagMaxInt64 {
		return flags, ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	if HasFlagUint64(flags, flag) == set {
		return flags, nil
	}

	conv := uint64(1 << flag)
	ret := flags ^ conv

	return ret, nil
}
