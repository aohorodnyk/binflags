package gobitflags

func FlagMap(flag uint64, bytes uint8) (uint64, uint8) {
	bytes64 := uint64(bytes)
	if flag <= bytes64 {
		return 0, uint8(flag)
	}

	bytes64Div := bytes64 + 1
	key := flag / bytes64Div
	val := flag % bytes64Div

	return key, uint8(val)
}
