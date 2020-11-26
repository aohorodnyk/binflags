package gobitflags

import (
	"errors"
)

func HasFlagUint32(flags uint32, flag uint8) bool {
	if flag > FlagMaxInt32 {
		return false
	}

	conv := uint32(1 << flag)
	return flags&conv == conv
}

func SetFlagUint32(flags uint32, flag uint8, set bool) (uint32, error) {
	if flag > FlagMaxInt32 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	if HasFlagUint32(flags, flag) == set {
		return flags, nil
	}

	conv := uint32(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
