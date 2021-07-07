package binflags

func flagExt(flag uint64, size uint8) (idx uint64, bit uint8) {
	bytes64 := uint64(size)
	if flag <= bytes64 {
		return 0, uint8(flag)
	}

	bytes64Div := bytes64 + 1
	key := flag / bytes64Div
	val := flag % bytes64Div

	return key, uint8(val)
}
