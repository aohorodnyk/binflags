package gobitflags

import (
	"errors"
)

func HasFlagInt8(flags int8, flag uint8) (bool, error) {
	if flag > FlagMaxInt8 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := int8(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagInt8(flags int8, flag uint8, set bool) (int8, error) {
	if flag > FlagMaxInt8 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt8(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := int8(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
