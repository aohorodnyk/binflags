package gobitflags

import (
	"errors"
)

func HasFlagUint32(flags uint32, flag uint8) (bool, error) {
	if flag > FlagMaxInt32 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := uint32(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagUint32(flags uint32, flag uint8, set bool) (uint32, error) {
	if flag > FlagMaxInt32 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagUint32(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := uint32(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
