package gobitflags

import (
	"errors"
)

func HasFlagInt64(flags int64, flag uint8) (bool, error) {
	if flag > FlagMaxInt64 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := int64(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagInt64(flags int64, flag uint8, set bool) (int64, error) {
	if flag > FlagMaxInt64 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt64(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := int64(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
