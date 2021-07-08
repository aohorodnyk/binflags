package binflags

// Max values for bits by int size.
const (
	FlagMaxInt8  uint8 = 7
	FlagMaxInt16 uint8 = 15
	FlagMaxInt32 uint8 = 31
	FlagMaxInt64 uint8 = 63
)

// List of error texts used in the library.
const (
	ErrorMsgOutOfRange    = "flag is out of range"
	ErrorMsgFlagsArrayNil = "flags array is nil"
	ErrorMsgFlagsMapNil   = "flags map is nil"
)
