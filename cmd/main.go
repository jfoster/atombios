package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/iancoleman/orderedmap"
	"github.com/jfoster/atombiostool/atom"
	"github.com/jfoster/atombiostool/atom/vega10"
	"github.com/jfoster/atombiostool/bios"

	"github.com/go-restruct/restruct"
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

	expectedchecksum := atom.CalcChecksum(data)
	currentchecksum := data[atom.AtomRomChecksumOffset]
	var color = chalk.Green
	if expectedchecksum != currentchecksum {
		color = chalk.Red
	}
	fmt.Printf("Checksum: %s0x%X%s\n", color, currentchecksum, chalk.Reset)

	b := bios.New()

	var header atom.AtomRomHeader
	headerOffset := binary.LittleEndian.Uint16(data[atom.AtomRomHeaderPointerOffset:])
	unpack(data, headerOffset, &header)
	b.SetO(header, headerOffset)

	if string(header.FirmWareSignature[:]) != "ATOM" {
		fmt.Println("Not an ATOM file!")
		return
	}

	var pciInfo atom.AtomRomPCIInfo
	pciInfoOffset := header.PCIInfoOffset
	unpack(data, pciInfoOffset, &pciInfo)
	b.SetO(pciInfo, pciInfoOffset)

	if pciInfo.Signature != "PCIR" {
		return
	}

	if pciInfo.VendorID != 0x1002 {
		return
	}

	fmt.Printf("DeviceID: 0x%X\n", pciInfo.DeviceID)

	var commandTable atom.AtomMasterCommandTable
	commandTableOffset := header.MasterCommandTableOffset
	unpack(data, commandTableOffset, &commandTable)
	b.SetO(commandTable, commandTableOffset)

	var dataTable atom.AtomMasterDataTable
	dataTableOffset := header.MasterDataTableOffset
	unpack(data, dataTableOffset, &dataTable)
	b.SetO(dataTable, dataTableOffset)

	var powerPlayTable vega10.AtomPowerPlayTable
	powerPlayTableOffset := dataTable.PowerPlayInfo
	unpack(data, powerPlayTableOffset, &powerPlayTable)
	b.SetO(powerPlayTable, powerPlayTableOffset)

	var powerTuneTable vega10.AtomPowerTuneTable
	powerTuneTableOffset := powerPlayTableOffset + powerPlayTable.PowerTuneTableOffset
	unpack(data, powerTuneTableOffset, &powerTuneTable)
	b.SetO(powerTuneTable, powerTuneTableOffset)

	var mClkTable vega10.AtomMClkTable
	mClkTableOffset := powerPlayTableOffset + powerPlayTable.MclkDependencyTableOffset
	unpack(data, mClkTableOffset, &mClkTable)
	b.SetO(mClkTable, mClkTableOffset)

	var gfxClkTable vega10.AtomGFXClkTableV2
	gfxClkTableOffset := powerPlayTableOffset + powerPlayTable.GfxclkDependencyTableOffset
	unpack(data, gfxClkTableOffset, &gfxClkTable)
	b.SetO(gfxClkTable, gfxClkTableOffset)

	var vramInfo atom.AtomVRAMInfoV23
	vramInfoOffset := dataTable.VRAMInfo
	unpack(data, vramInfoOffset, &vramInfo)
	vramEntries := make([]atom.AtomVRAMModuleV9, vramInfo.NumOfVRAMModule)
	size, _ := restruct.SizeOf(vramInfo)
	vramEntryOffset := vramInfoOffset + uint16(size)
	for i := range vramEntries {
		unpack(data, vramEntryOffset, &vramEntries[i])
		vramEntryOffset += vramEntries[i].ModuleSize
	}
	vramInfo.VramInfo = vramEntries
	b.SetO(vramInfo, vramInfoOffset)

	var voltageInfo atom.AtomVoltageInfo
	voltageInfoOffset := dataTable.VoltageObjectInfo
	unpack(data, voltageInfoOffset, &voltageInfo)
	voltageEntries := make([]byte, uint16(voltageInfo.NumOfVoltageEntries)*uint16(voltageInfo.BytesPerVoltageEntry))
	size, _ = restruct.SizeOf(voltageInfo)
	voltageEntryOffset := voltageInfoOffset + uint16(size)
	for i := range voltageEntries {
		unpack(data, voltageEntryOffset, &voltageEntries[i])
		voltageEntryOffset++
	}
	voltageInfo.VoltageEntries = voltageEntries
	b.SetO(voltageInfo, voltageInfoOffset)

	var hbmVoltageTable vega10.AtomVoltageTable
	hbmVoltageTableOffset := powerPlayTableOffset + powerPlayTable.VddmemLookupTableOffset
	unpack(data, hbmVoltageTableOffset, &hbmVoltageTable)
	b.SetO(hbmVoltageTable, hbmVoltageTableOffset)

	var hardLimitTable vega10.AtomHardLimitTable
	hardLimitTableOffset := powerPlayTableOffset + powerPlayTable.HardLimitTableOffset
	unpack(data, hardLimitTableOffset, &hardLimitTable)
	b.SetO(hardLimitTable, hardLimitTableOffset)

	asicInfoItem := bios.Item{Offset: dataTable.ASICProfilingInfo, Table: atom.AtomAsicProfilingInfoV35{}}
	unpack(data, asicInfoItem.Offset, &asicInfoItem.Table)
	b.SetItem2(asicInfoItem)

	/* */

	b.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(bios.Item).Offset < b.Value().(bios.Item).Offset
	})
	for _, k := range b.Keys() {
		v, _ := b.Get(k)
		item := v.(bios.Item)
		fmt.Printf("%s0x%X: %s: %s%+v%s\n", chalk.Red, item.Offset, k, chalk.Cyan, item.Table, chalk.Reset)
	}
}

func unpackItem(bytes []byte, item *bios.Item) {
	unpack(bytes, item.Offset, &item.Table)
}

func unpack(bytes []byte, offset uint16, object interface{}) {
	err := restruct.Unpack(bytes[offset:], binary.LittleEndian, object)
	if err != nil {
		fmt.Println(err)
	}
}
