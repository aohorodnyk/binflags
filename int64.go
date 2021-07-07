package binflags

func HasFlagInt64(flags int64, flag uint8) bool {
	if flag > FlagMaxInt64 {
		return false
	}

	conv := int64(1 << flag)

	return flags&conv == conv
}

func SetFlagInt64(flags int64, flag uint8, set bool) (int64, error) {
	if flag > FlagMaxInt64 {
		return flags, ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	if HasFlagInt64(flags, flag) == set {
		return flags, nil
	}

	conv := int64(1 << flag)
	ret := flags ^ conv

	return ret, nil
}
