package vega10

import "github.com/jfoster/atombios/atom"

type AtomPowerPlayTable struct {
	Header                       atom.AtomCommonTableHeader
	TableRevision                byte
	TableSize                    uint16
	GoldenPPID                   uint32
	GoldenRevision               uint32
	FormatID                     uint16
	PlatformCaps                 uint32
	MaxODEngineClock             uint32
	MaxODMemoryClock             uint32
	PowerControlLimit            uint16
	UlvVoltageOffset             uint16
	UlvSmnclkDid                 uint16
	UlvMp1clkDid                 uint16
	UlvGfxclkBypass              uint16
	GfxclkSlewRate               uint16
	GfxVoltageMode               byte
	SocVoltageMode               byte
	UclkVoltageMode              byte
	UvdVoltageMode               byte
	VceVoltageMode               byte
	Mp0VoltageMode               byte
	DcefVoltageMode              byte
	StateArrayOffset             uint16
	FanTableOffset               uint16
	ThermalControllerOffset      uint16
	SocclkDependencyTableOffset  uint16
	MclkDependencyTableOffset    uint16
	GfxclkDependencyTableOffset  uint16
	DcefclkDependencyTableOffset uint16
	VddcLookupTableOffset        uint16
	VddmemLookupTableOffset      uint16
	MMDependencyTableOffset      uint16
	VCEStateTableOffset          uint16
	_                            uint16
	PowerTuneTableOffset         uint16
	HardLimitTableOffset         uint16
	VddciLookupTableOffset       uint16
	PCIETableOffset              uint16
	PixclkDependencyTableOffset  uint16
	DispClkDependencyTableOffset uint16
	PhyClkDependencyTableOffset  uint16
}

type AtomFanTable struct {
	RevId                byte   /* Change this if the table format changes or version changes so that the other fields are not the same. */
	FanOutputSensitivity uint16 /* Sensitivity of fan reaction to temepature changes. */
	FanRPMMax            uint16 /* The default value in RPM. */
	ThrottlingRPM        uint16
	FanAcousticLimit     uint16 /* Minimum Fan Controller Frequency Acoustic Limit. */
	TargetTemperature    uint16 /* The default ideal temperature in Celcius. */
	MinimumPWMLimit      uint16 /* The minimum PWM that the advanced fan controller can set. */
	TargetGfxClk         uint16 /* The ideal Fan Controller GFXCLK Frequency Acoustic Limit. */
	FanGainEdge          uint16
	FanGainHotspot       uint16
	FanGainLiquid        uint16
	FanGainVrVddc        uint16
	FanGainVrMvdd        uint16
	FanGainPlx           uint16
	FanGainHbm           uint16
	EnableZeroRPM        byte
	FanStopTemperature   uint16
	FanStartTemperature  uint16
}

type AtomFanTableV2 struct {
	RevId                byte
	FanOutputSensitivity uint16
	FanAcousticLimitRpm  uint16
	ThrottlingRPM        uint16
	TargetTemperature    uint16
	MinimumPWMLimit      uint16
	TargetGfxClk         uint16
	FanGainEdge          uint16
	FanGainHotspot       uint16
	FanGainLiquid        uint16
	FanGainVrVddc        uint16
	FanGainVrMvdd        uint16
	FanGainPlx           uint16
	FanGainHbm           uint16
	EnableZeroRPM        byte
	FanStopTemperature   uint16
	FanStartTemperature  uint16
	FanParameters        byte
	FanMinRPM            byte
	FanMaxRPM            byte
}

type AtomFanTableV3 struct {
	RevId                byte
	FanOutputSensitivity uint16
	FanAcousticLimitRpm  uint16
	ThrottlingRPM        uint16
	TargetTemperature    uint16
	MinimumPWMLimit      uint16
	TargetGfxClk         uint16
	FanGainEdge          uint16
	FanGainHotspot       uint16
	FanGainLiquid        uint16
	FanGainVrVddc        uint16
	FanGainVrMvdd        uint16
	FanGainPlx           uint16
	FanGainHbm           uint16
	EnableZeroRPM        byte
	FanStopTemperature   uint16
	FanStartTemperature  uint16
	FanParameters        byte
	FanMinRPM            byte
	FanMaxRPM            byte
	MGpuThrottlingRPM    uint16
}

type AtomMClkTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomMClkEntry
}

type AtomMClkEntry struct {
	MemClk    uint32 /* Clock Frequency */
	VddInd    byte   /* SOC_VDD index */
	VddMemInd byte   /* MEM_VDD - only non zero for MCLK record */
	VddciInd  byte   /* VDDCI   = only non zero for MCLK record */
}

type AtomGFXClkTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomGFXClkEntry
}

type AtomGFXClkEntry struct {
	Clk                  uint32 /* Clock Frequency */
	VddInd               byte   /* SOC_VDD index */
	CKSVOffsetandDisable uint16 /* Bits 0~30: Voltage offset for CKS, Bit 31: Disable/enable for the GFXCLK level. */
	AVFSOffset           uint16 /* AVFS Voltage offset */
}

type AtomGFXClkTableV2 struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomGFXClkEntryV2
}

type AtomGFXClkEntryV2 struct {
	Clk                  uint32 /* Clock Frequency */
	VddInd               byte   /* SOC_VDD index */
	CKSVOffsetandDisable uint16 /* Bits 0~30: Voltage offset for CKS, Bit 31: Disable/enable for the GFXCLK level. */
	AVFSOffset           uint16
	ACGEnable            byte
	_                    [3]byte
}

type AtomVoltageEntry struct {
	Vdd uint16 /* Base voltage */
}

type AtomVoltageTable struct {
	RevId      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomVoltageEntry
}

type AtomHardLimitEntry struct {
	SOCCLKLimit uint32
	GFXCLKLimit uint32
	MCLKLimit   uint32
	VddcLimit   uint16
	VddciLimit  uint16
	VddMemLimit uint16
}

type AtomHardLimitTable struct {
	RevId      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomHardLimitEntry
}
