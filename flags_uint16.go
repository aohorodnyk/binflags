package gobitflags

import (
	"errors"
)

func HasFlagUint16(flags uint16, flag uint8) (bool, error) {
	if flag > FlagMaxInt16 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	var conv uint16 = 1 << flag
	return flags&conv == conv, nil
}

func SetFlagUint16(flags uint16, flag uint8, val bool) (uint16, error) {
	if flag > FlagMaxInt16 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagUint16(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == val {
		return flags, nil
	}

	var conv uint16 = 1 << flag
	ret := flags ^ conv
	return ret, nil
}
