package atom

func FixChecksum(bytes *[]byte) {
	b := *bytes
	b[AtomRomChecksumOffset] = CalcChecksum(b)
	bytes = &b
}

func CalcChecksum(bytes []byte) (b byte) {
	if len(bytes) > 0 {
		size := int(bytes[AtomRomHeaderSizeOffset]) * 512
		b = bytes[AtomRomChecksumOffset] - calcOffset(bytes[:size])
	}
	return b
}

func calcOffset(bytes []byte) (b byte) {
	for _, v := range bytes {
		b += v
	}
	return b
}
