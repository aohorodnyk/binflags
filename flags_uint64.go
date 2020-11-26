package gobitflags

import (
	"errors"
)

func HasFlagUint64(flags uint64, flag uint8) (bool, error) {
	if flag > FlagMaxInt64 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := uint64(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagUint64(flags uint64, flag uint8, set bool) (uint64, error) {
	if flag > FlagMaxInt64 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagUint64(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := uint64(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
