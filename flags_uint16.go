package gobitflags

import (
	"errors"
)

func HasFlagUint16(flags uint16, flag uint8) (bool, error) {
	if flag > FlagMaxInt16 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := uint16(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagUint16(flags uint16, flag uint8, set bool) (uint16, error) {
	if flag > FlagMaxInt16 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagUint16(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := uint16(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
