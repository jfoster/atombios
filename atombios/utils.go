package atombios

import (
	"encoding/binary"
	"fmt"

	"github.com/go-restruct/restruct"
)

func Unpack(bytes []byte, offset uint16, object interface{}) {
	err := restruct.Unpack(bytes[offset:], binary.LittleEndian, object)
	if err != nil {
		fmt.Println(err)
	}
}

func CalcOffset(bytes []byte) byte {
	var o byte
	for i := 0; i < len(bytes); i++ {
		o += bytes[i]
	}
	return o
}

func FixChecksum(bytes []byte) {
	size := int(bytes[AtomRomHeaderSizeOffset]) * 512
	bytes[AtomRomChecksumOffset] -= CalcOffset(bytes[:size])
}
