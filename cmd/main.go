package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jfoster/atombiostool/atombios"
	"github.com/ttacon/chalk"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	if len(data)%0x100 != 0 {
		fmt.Println("Invalid file!")
		return
	}

	fmt.Printf("0x%X\n", data[atombios.AtomRomChecksumOffset])
	atombios.FixChecksum(data)
	fmt.Printf("0x%X\n", data[atombios.AtomRomChecksumOffset])

	offset := binary.LittleEndian.Uint16(data[atombios.AtomRomHeaderPointerOffset:])
	header := atombios.AtomRomHeader{}
	atombios.Unpack(data, offset, &header)
	test(header, offset)

	if string(header.FirmWareSignature[:]) != "ATOM" {
		fmt.Println("Not an ATOM file!")
		return
	}

	offset = header.MasterCommandTableOffset
	commandtable := atombios.AtomMasterCommandTable{}
	atombios.Unpack(data, offset, &commandtable)
	test(commandtable, offset)

	offset = header.MasterDataTableOffset
	datatable := atombios.AtomMasterDataTable{}
	atombios.Unpack(data, offset, &datatable)
	test(datatable, offset)

	offset = datatable.PowerPlayInfo
	powerplaytable := atombios.AtomTongaPowerPlayTable{}
	atombios.Unpack(data, offset, &powerplaytable)
	test(powerplaytable, offset)

	offset = datatable.PowerPlayInfo + powerplaytable.PowerTuneTableOffset
	powertunetable := atombios.AtomFijiPowerTuneTable{}
	atombios.Unpack(data, offset, &powertunetable)
	test(powertunetable, offset)

}

func test(object interface{}, offset uint16) {
	fmt.Printf("%s%T @ 0x%X: %s%+v%s\n", chalk.Red, object, offset, chalk.Cyan, object, chalk.Reset)
}
