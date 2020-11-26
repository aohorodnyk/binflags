package gobitflags

import (
	"errors"
)

func HasFlagInt16(flags int16, flag uint8) (bool, error) {
	if flag > FlagMaxInt16 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := int16(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagInt16(flags int16, flag uint8, set bool) (int16, error) {
	if flag > FlagMaxInt16 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt16(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := int16(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
