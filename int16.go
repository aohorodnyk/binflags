package binflags

func HasFlagInt16(flags int16, flag uint8) bool {
	if flag > FlagMaxInt16 {
		return false
	}

	conv := int16(1 << flag)

	return flags&conv == conv
}

func SetFlagInt16(flags int16, flag uint8, set bool) (int16, error) {
	if flag > FlagMaxInt16 {
		return flags, ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	if HasFlagInt16(flags, flag) == set {
		return flags, nil
	}

	conv := int16(1 << flag)
	ret := flags ^ conv

	return ret, nil
}
