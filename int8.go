package gobitflags

import (
	"errors"
)

func HasFlagInt8(flags int8, flag uint8) bool {
	if flag > FlagMaxInt8 {
		return false
	}

	conv := int8(1 << flag)
	return flags&conv == conv
}

func SetFlagInt8(flags int8, flag uint8, set bool) (int8, error) {
	if flag > FlagMaxInt8 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	if HasFlagInt8(flags, flag) == set {
		return flags, nil
	}

	conv := int8(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
