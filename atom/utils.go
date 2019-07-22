package atom

func FixChecksum(bytes []byte) {
	bytes[AtomRomChecksumOffset] = CalcChecksum(bytes)
}

func CalcChecksum(bytes []byte) byte {
	size := int(bytes[AtomRomHeaderSizeOffset]) * 512
	checksum := bytes[AtomRomChecksumOffset] - calcOffset(bytes[:size])
	return checksum
}

func calcOffset(bytes []byte) byte {
	var o byte
	for i := 0; i < len(bytes); i++ {
		o += bytes[i]
	}
	return o
}
