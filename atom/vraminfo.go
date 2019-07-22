package atom

type AtomVRAMInfoV21 struct {
	Header               AtomCommonTableHeader
	MemAdjustTblOffset   uint16 // offset of ATOM_INIT_REG_BLOCK structure for memory vendor specific MC adjust setting
	MemClkPatchTblOffset uint16 // offset of ATOM_INIT_REG_BLOCK structure for memory clock specific MC setting
	PerBytePresetOffset  uint16 // offset of ATOM_INIT_REG_BLOCK structure for Per Byte Offset Preset Settings
	usReserved           [3]uint16
	NumOfVRAMModule      byte // indicate number of VRAM module
	MemoryClkPatchTblVer byte // version of memory AC timing register list
	VramModuleVer        byte // indicate ATOM_VRAM_MODUE version
	_                    byte
	VramInfo             []AtomVRAMModuleV7 // just for allocation, real number of blocks is in ucNumOfVRAMModule;
}

type AtomVRAMModuleV7 struct {
}

type AtomVRAMInfoV22 struct {
	Header                   AtomCommonTableHeader
	MemAdjustTblOffset       uint16
	MemClkPatchTblOffset     uint16
	McAdjustPerTileTblOffset uint16
	McPhyInitTableOffset     uint16
	DramDataRemapTblOffset   uint16
	_                        uint16
	NumOfVRAMModule          byte
	MemoryClkPatchTblVer     byte
	VramModuleVer            byte
	McPhyTileNum             byte
	VramInfo                 []AtomVRAMModuleV8
}

type AtomVRAMModuleV8 struct {
	ChannelMapCfg     uint32
	ModuleSize        uint16
	McRAMCfg          uint16
	EnableChannels    uint16
	ExtMemoryID       byte
	MemoryType        byte
	ChannelNum        byte
	ChannelWidth      byte
	Density           byte
	BankCol           byte
	Misc              byte
	VREFI             byte
	_                 uint16
	MemorySize        uint16
	McTunningSetID    byte
	RowNum            byte
	EMRS2Value        uint16
	EMRS3Value        uint16
	MemoryVenderID    byte
	RefreshRateFactor byte
	FIFODepth         byte
	CDRBandwidth      byte
	ChannelMapCfg1    uint32
	BankMapCfg        uint32
	_                 uint32
	MemPNString       string `struct:"[20]byte"`
}

type AtomVRAMInfoV23 struct {
	Header                   AtomCommonTableHeader
	MemAdjustTblOffset       uint16
	MemClkPatchTblOffset     uint16
	McAdjustPerTileTblOffset uint16
	McPhyInitTableOffset     uint16
	DramDataRemapTblOffset   uint16
	TmrsSeqOffset            uint16 // offset of HBM tmrs
	PostUcodeInitOffset      uint16 // offset of atom_umc_init_reg_block structure for MC phy init after MC uCode complete umc init
	VramRsd2                 uint16
	NumOfVRAMModule          byte // indicate number of VRAM module
	_                        [2]byte
	McPhyTileNum             byte
	VramInfo                 []AtomVRAMModuleV9 // just for allocation, real number of blocks is in ucNumOfVRAMModule;
}

type AtomVRAMModuleV9 struct {
	MemorySize        uint32 // Total memory size in unit of MB for CONFIG_MEMSIZE zeros
	EnableChannels    uint32 // bit vector, each bit indicate specific channel enable or not
	MaxMemoryClock    uint32 // max memory clock of this memory in unit of 10kHz, =0 means it is not defined
	_                 [3]uint16
	MemoryVoltage     uint16 // mem_voltage
	ModuleSize        uint16 // Size of atom_vram_module_v9
	ExtMemoryID       byte   // Current memory module ID
	MemoryType        byte   // enum of atom_dgpu_vram_type
	ChannelNum        byte   // Number of mem. channels supported in this module
	ChannelWidth      byte   // CHANNEL_16BIT/CHANNEL_32BIT/CHANNEL_64BIT
	Density           byte   // _8Mx32, _16Mx32, _16Mx16, _32Mx16
	McTunningSetID    byte   // MC phy registers set per.
	VendorRevID       byte   // [7:4] Revision, [3:0] Vendor code
	RefreshRateFactor byte   // [1:0]=RefreshFactor (00=8ms, 01=16ms, 10=32ms,11=64ms)
	HBMVendorRevID    byte   // hbm_ven_rev_id
	_                 byte   // reserved
	MemPNString       string `struct:"[20]byte"` // part number end with '0'.
}
