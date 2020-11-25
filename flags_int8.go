package gobitflags

import (
	"errors"
	"math"
)

func HasFlagInt8(flags int8, flag uint8) (bool, error) {
	if flag > FlagMaxInt8 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := flagInt8(flag)
	return flags&conv == conv, nil
}

func SetFlagInt8(flags int8, flag uint8, val bool) (int8, error) {
	if flag > FlagMaxInt8 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt8(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == val {
		return flags, nil
	}

	conv := flagInt8(flag)
	ret := flags ^ conv
	return ret, nil
}

func flagInt8(flag uint8) int8 {
	var ret int8 = math.MinInt8
	if flag < FlagMaxInt8 {
		ret = 1 << flag
	}

	return ret
}
