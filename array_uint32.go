package binflags

import "errors"

func HasFlagArrayUint32(flags []uint32, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt32)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagUint32(flags[idx], bit)
}

func SetFlagArrayUint32(flags []uint32, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayUint32(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt32)
	if len(flags) <= int(idx) {
		return errors.New(ErrorMsgOutOfRange)
	}

	var err error
	flags[idx], err = SetFlagUint32(flags[idx], bit, set)
	return err
}
