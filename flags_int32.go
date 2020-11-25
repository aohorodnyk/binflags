package gobitflags

import (
	"errors"
	"math"
)

func HasFlagInt32(flags int32, flag uint8) (bool, error) {
	if flag > FlagMaxInt32 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := flagInt32(flag)
	return flags&conv == conv, nil
}

func SetFlagInt32(flags int32, flag uint8, val bool) (int32, error) {
	if flag > FlagMaxInt32 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt32(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == val {
		return flags, nil
	}

	conv := flagInt32(flag)
	ret := flags ^ conv
	return ret, nil
}

func flagInt32(flag uint8) int32 {
	var ret int32 = math.MinInt32
	if flag < FlagMaxInt32 {
		ret = 1 << flag
	}

	return ret
}
