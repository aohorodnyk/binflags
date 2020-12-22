package binflags

import (
	"errors"
)

func HasFlagUint16(flags uint16, flag uint8) bool {
	if flag > FlagMaxInt16 {
		return false
	}

	conv := uint16(1 << flag)
	return flags&conv == conv
}

func SetFlagUint16(flags uint16, flag uint8, set bool) (uint16, error) {
	if flag > FlagMaxInt16 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	if HasFlagUint16(flags, flag) == set {
		return flags, nil
	}

	conv := uint16(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
