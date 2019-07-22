package atom

type AtomRomHeader struct {
	Header                    AtomCommonTableHeader
	FirmWareSignature         string `struct:"[4]byte"`
	BiosRuntimeSegmentAddress uint16
	ProtectedModeInfoOffset   uint16
	ConfigFilenameOffset      uint16
	CRCBlockOffset            uint16
	BIOSBootupMessageOffset   uint16
	Int10Offset               uint16
	PciBDevInitCode           uint16
	IoBaseAddress             uint16
	SubsystemVendorID         uint16
	SubsystemID               uint16
	PCIInfoOffset             uint16
	MasterCommandTableOffset  uint16
	MasterDataTableOffset     uint16
	ExtendedFunctionCode      byte
	_                         byte
}
