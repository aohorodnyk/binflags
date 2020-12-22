package binflags

import "errors"

func HasFlagMapInt64(flags map[uint64]int64, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	return HasFlagInt64(flags[idx], bit)
}

func SetFlagMapInt64(flags map[uint64]int64, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsMapNil)
	}

	if HasFlagMapInt64(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	bits, ok := flags[idx]

	ret, err := SetFlagInt64(bits, bit, set)
	if err != nil {
		return err
	}

	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}

	return nil
}
