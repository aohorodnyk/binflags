package gobitflags

import (
	"errors"
	"math"
)

func HasFlagInt16(flags int16, flag uint8) (bool, error) {
	if flag > FlagMaxInt16 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := flagInt16(flag)
	return flags&conv == conv, nil
}

func SetFlagInt16(flags int16, flag uint8, val bool) (int16, error) {
	if flag > FlagMaxInt16 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt16(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == val {
		return flags, nil
	}

	conv := flagInt16(flag)
	ret := flags ^ conv
	return ret, nil
}

func flagInt16(flag uint8) int16 {
	var ret int16 = math.MinInt16
	if flag < FlagMaxInt16 {
		ret = 1 << flag
	}

	return ret
}
