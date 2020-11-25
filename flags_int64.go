package gobitflags

import (
	"errors"
	"math"
)

func HasFlagInt64(flags int64, flag uint8) (bool, error) {
	if flag > FlagMaxInt64 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := flagInt64(flag)
	return flags&conv == conv, nil
}

func SetFlagInt64(flags int64, flag uint8, val bool) (int64, error) {
	if flag > FlagMaxInt64 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagInt64(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == val {
		return flags, nil
	}

	conv := flagInt64(flag)
	ret := flags ^ conv
	return ret, nil
}

func flagInt64(flag uint8) int64 {
	var ret int64 = math.MinInt64
	if flag < FlagMaxInt64 {
		ret = 1 << flag
	}

	return ret
}
