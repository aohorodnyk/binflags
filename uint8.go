package binflags

func HasFlagByte(flags byte, flag uint8) bool {
	return HasFlagUint8(flags, flag)
}

func SetFlagByte(flags, flag uint8, set bool) (uint8, error) {
	return SetFlagUint8(flags, flag, set)
}

func HasFlagUint8(flags, flag uint8) bool {
	if flag > FlagMaxInt8 {
		return false
	}

	conv := uint8(1 << flag)

	return flags&conv == conv
}

func SetFlagUint8(flags, flag uint8, set bool) (uint8, error) {
	if flag > FlagMaxInt8 {
		return flags, ErrorOutOfRange(ErrorMsgOutOfRange)
	}

	if HasFlagUint8(flags, flag) == set {
		return flags, nil
	}

	conv := uint8(1 << flag)
	ret := flags ^ conv

	return ret, nil
}
