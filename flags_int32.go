package gobitflags

import (
	"errors"
)

func HasFlagInt32(flags int32, flag uint8) (bool, error) {
	if flag > FlagMaxInt32 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := int32(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagInt32(flags int32, flag uint8, set bool) (int32, error) {
	if flag > FlagMaxInt32 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt32(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := int32(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
