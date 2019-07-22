package atom

type AtomSClkEntry struct {
	VddInd                 byte
	VddcOffset             uint16
	Sclk                   uint32
	EdcCurrent             uint16
	ReliabilityTemperature byte
	CKSVOffsetandDisable   byte
	SclkOffset             uint32
}

type AtomSClkTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomSClkEntry
}

type AtomVoltageEntry struct {
	Vdd    uint16
	CACLow uint16
	CACMid uint16
	CACHig uint16
}

type AtomVoltageTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomVoltageEntry
}

type AtomVRAMTimingEntry struct {
	ClkRang uint32
	Latency [0x3]byte
}

type EFUSE_LINEAR_FUNC_PARAM struct {
	EfuseIndex      uint16 // Efuse Index in DWORD address, for example Index 911, usEuseIndex=112
	EfuseBitLSB     byte   // Efuse bit LSB in DWORD address, for example Index 911, usEfuseBitLSB= 911-112*8=15
	EfuseLength     byte   // Efuse bits length,
	EfuseEncodeRang uint32 // Range = Max - Min, bit31 indicate the efuse is negative number
	EfuseMin        uint32 // Min
}

// for  Polaris10/Polaris11 speed EVV algorithm
type AtomAsicProfilingInfoV35 struct {
	Header                AtomCommonTableHeader
	MaxVdd                uint32                  //Maximum voltage for all parts, in unit of 0.01mv
	MinVddc               uint32                  //Minimum voltage for all parts, in unit of 0.01mv
	LkgEuseIndex          uint16                  //Efuse Lkg_FT address ( BYTE address )
	LkgEfuseBitLSB        byte                    //Efuse Lkg_FT bit shift in 32bit DWORD
	LkgEfuseLength        byte                    //Efuse Lkg_FT length
	LkgEncodeLn_MaxDivMin uint32                  //value of ln(Max_Lkg_Ft/Min_Lkg_Ft ) in unit of 0.00001 ( unit=100000 )
	LkgEncodeMax          uint32                  //Maximum Lkg_Ft measured value ( or efuse decode value ), in unit of 0.00001 ( unit=100000 )
	LkgEncodeMin          uint32                  //Minimum Lkg_Ft measured value ( or efuse decode value ), in unit of 0.00001 ( unit=100000 )
	RoFuse                EFUSE_LINEAR_FUNC_PARAM //Efuse RO info: DWORD address, bit shift, length, max/min measure value. in unit of 1.
	EvvDefaultVddc        uint32                  //def="EVV_DEFAULT_VDDC" descr="return default VDDC(v) when Efuse not cut" unit="100000"/>
	EvvNoCalcVddc         uint32                  //def="EVV_NOCALC_VDDC" descr="return VDDC(v) when Calculation is bad" unit="100000"/>
	Speed_Model           uint32                  //def="EVV_SPEED_MODEL" descr="0 = Greek model, 1 = multivariate model" unit="1"/>
	SM_A0                 uint32                  //def="EVV_SM_A0" descr="Leakage coeff(Multivariant Mode)." unit="100000"/>
	SM_A1                 uint32                  //def="EVV_SM_A1" descr="Leakage/SCLK coeff(Multivariant Mode)." unit="1000000"/>
	SM_A2                 uint32                  //def="EVV_SM_A2" descr="Alpha( Greek Mode ) or VDDC/SCLK coeff(Multivariant Mode)." unit="100000"/>
	SM_A3                 uint32                  //def="EVV_SM_A3" descr="Beta( Greek Mode ) or SCLK coeff(Multivariant Mode)." unit="100000"/>
	SM_A4                 uint32                  //def="EVV_SM_A4" descr="VDDC^2/SCLK coeff(Multivariant Mode)." unit="100000"/>
	SM_A5                 uint32                  //def="EVV_SM_A5" descr="VDDC^2 coeff(Multivariant Mode)." unit="100000"/>
	SM_A6                 uint32                  //def="EVV_SM_A6" descr="Gamma( Greek Mode ) or VDDC coeff(Multivariant Mode)." unit="100000"/>
	SM_A7                 uint32                  //def="EVV_SM_A7" descr="Epsilon( Greek Mode ) or constant(Multivariant Mode)." unit="100000"/>
	SM_A0_sign            byte                    //def="EVV_SM_A0_SIGN" descr="=0 SM_A0 is postive. =1: SM_A0 is negative" unit="1"/>
	SM_A1_sign            byte                    //def="EVV_SM_A1_SIGN" descr="=0 SM_A1 is postive. =1: SM_A1 is negative" unit="1"/>
	SM_A2_sign            byte                    //def="EVV_SM_A2_SIGN" descr="=0 SM_A2 is postive. =1: SM_A2 is negative" unit="1"/>
	SM_A3_sign            byte                    //def="EVV_SM_A3_SIGN" descr="=0 SM_A3 is postive. =1: SM_A3 is negative" unit="1"/>
	SM_A4_sign            byte                    //def="EVV_SM_A4_SIGN" descr="=0 SM_A4 is postive. =1: SM_A4 is negative" unit="1"/>
	SM_A5_sign            byte                    //def="EVV_SM_A5_SIGN" descr="=0 SM_A5 is postive. =1: SM_A5 is negative" unit="1"/>
	SM_A6_sign            byte                    //def="EVV_SM_A6_SIGN" descr="=0 SM_A6 is postive. =1: SM_A6 is negative" unit="1"/>
	SM_A7_sign            byte                    //def="EVV_SM_A7_SIGN" descr="=0 SM_A7 is postive. =1: SM_A7 is negative" unit="1"/>
	Margin_RO_a           uint32                  //def="EVV_MARGIN_RO_A" descr="A Term to represent RO equation in Ax2+Bx+C, unit=1"
	Margin_RO_b           uint32                  //def="EVV_MARGIN_RO_B" descr="B Term to represent RO equation in Ax2+Bx+C, unit=1"
	Margin_RO_c           uint32                  //def="EVV_MARGIN_RO_C" descr="C Term to represent RO equation in Ax2+Bx+C, unit=1"
	Margin_fixed          uint32                  //def="EVV_MARGIN_FIXED" descr="Fixed MHz to add to SCLK margin, unit=1" unit="1"/>
	Margin_Fmax_mean      uint32                  //def="EVV_MARGIN_FMAX_MEAN" descr="Percentage to add for Fmas mean margin unit=10000" unit="10000"/>
	Margin_plat_mean      uint32                  //def="EVV_MARGIN_PLAT_MEAN" descr="Percentage to add for platform mean margin unit=10000" unit="10000"/>
	Margin_Fmax_sigma     uint32                  //def="EVV_MARGIN_FMAX_SIGMA" descr="Percentage to add for Fmax sigma margin unit=10000" unit="10000"/>
	Margin_plat_sigma     uint32                  //def="EVV_MARGIN_PLAT_SIGMA" descr="Percentage to add for platform sigma margin unit=10000" unit="10000"/>
	Margin_DC_sigma       uint32                  //def="EVV_MARGIN_DC_SIGMA" descr="Regulator DC tolerance margin (mV) unit=100" unit="100"/>
	Reserved              [1]uint32
}

/* for Polars10/11 AVFS parameters */
type AtomAsicProfilingInfoV36 struct {
	Header                          AtomCommonTableHeader
	MaxVddc                         uint32
	MinVddc                         uint32
	LkgEuseIndex                    uint16
	LkgEfuseBitLSB                  byte
	LkgEfuseLength                  byte
	LkgEncodeLn_MaxDivMin           uint32
	LkgEncodeMax                    uint32
	LkgEncodeMin                    uint32
	RoFuse                          EFUSE_LINEAR_FUNC_PARAM
	EvvDefaultVddc                  uint32
	EvvNoCalcVddc                   uint32
	Speed_Model                     uint32
	SM_A0                           uint32
	SM_A1                           uint32
	SM_A2                           uint32
	SM_A3                           uint32
	SM_A4                           uint32
	SM_A5                           uint32
	SM_A6                           uint32
	SM_A7                           uint32
	SM_A0_sign                      byte
	SM_A1_sign                      byte
	SM_A2_sign                      byte
	SM_A3_sign                      byte
	SM_A4_sign                      byte
	SM_A5_sign                      byte
	SM_A6_sign                      byte
	SM_A7_sign                      byte
	Margin_RO_a                     uint32
	Margin_RO_b                     uint32
	Margin_RO_c                     uint32
	Margin_fixed                    uint32
	Margin_Fmax_mean                uint32
	Margin_plat_mean                uint32
	Margin_Fmax_sigma               uint32
	Margin_plat_sigma               uint32
	Margin_DC_sigma                 uint32
	LoadLineSlop                    uint32
	aTDClimitPerDPM                 [8]uint32
	aNoCalcVddcPerDPM               [8]uint32
	AVFS_meanNsigma_Acontant0       uint32
	AVFS_meanNsigma_Acontant1       uint32
	AVFS_meanNsigma_Acontant2       uint32
	AVFS_meanNsigma_DC_tol_sigma    uint16
	AVFS_meanNsigma_Platform_mean   uint16
	AVFS_meanNsigma_Platform_sigma  uint16
	GB_VDROOP_TABLE_CKSOFF_a0       uint32
	GB_VDROOP_TABLE_CKSOFF_a1       uint32
	GB_VDROOP_TABLE_CKSOFF_a2       uint32
	GB_VDROOP_TABLE_CKSON_a0        uint32
	GB_VDROOP_TABLE_CKSON_a1        uint32
	GB_VDROOP_TABLE_CKSON_a2        uint32
	AVFSGB_FUSE_TABLE_CKSOFF_m1     uint32
	AVFSGB_FUSE_TABLE_CKSOFF_m2     uint16
	AVFSGB_FUSE_TABLE_CKSOFF_b      uint32
	AVFSGB_FUSE_TABLE_CKSON_m1      uint32
	AVFSGB_FUSE_TABLE_CKSON_m2      uint16
	AVFSGB_FUSE_TABLE_CKSON_b       uint32
	MaxVoltage_0_25mv               uint16
	EnableGB_VDROOP_TABLE_CKSOFF    byte
	EnableGB_VDROOP_TABLE_CKSON     byte
	EnableGB_FUSE_TABLE_CKSOFF      byte
	EnableGB_FUSE_TABLE_CKSON       byte
	PSM_Age_ComFactor               uint16
	EnableApplyAVFS_CKS_OFF_Voltage byte
	Reserved                        byte
}
