package gobitflags

import "errors"

func HasFlagMapUint16(flags map[uint64]uint16, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	return HasFlagUint16(flags[idx], bit)
}

func SetFlagMapUint16(flags map[uint64]uint16, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsMapNil)
	}

	if HasFlagMapUint16(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	bits, ok := flags[idx]

	ret, err := SetFlagUint16(bits, bit, set)
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
