package binflags

func HasFlagInt32(flags int32, flag uint8) bool {
	if flag > FlagMaxInt32 {
		return false
	}

	conv := int32(1 << flag)

	return flags&conv == conv
}

func SetFlagInt32(flags int32, flag uint8, set bool) (int32, error) {
	if flag > FlagMaxInt32 {
		return flags, ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	if HasFlagInt32(flags, flag) == set {
		return flags, nil
	}

	conv := int32(1 << flag)
	ret := flags ^ conv

	return ret, nil
}
