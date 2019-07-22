package atom

type AtomTongaPowerPlayTable struct {
	Header                    AtomCommonTableHeader
	TableRevision             byte
	TableSize                 uint16 /*the size of header structure */
	GoldenPPID                uint32
	GoldenRevision            uint32
	FormatID                  uint16
	VoltageTime               uint16 /*in microseconds */
	PlatformCaps              uint32 /*See ATOM_Tonga_CAPS_* */
	MaxODEngineClock          uint32 /*For Overdrive.  */
	MaxODMemoryClock          uint32 /*For Overdrive. */
	PowerControlLimit         uint16
	UlvVoltageOffset          uint16    /*in mv units */
	StateArrayOffset          uint16    /*points to ATOM_Tonga_State_Array */
	FanTableOffset            uint16    /*points to ATOM_Tonga_Fan_Table */
	ThermalControllerOffset   uint16    /*points to ATOM_Tonga_Thermal_Controller */
	_                         uint16    /*CtomThermalPolicy removed for Tonga. Keep this filed as reserved. */
	MclkDependencyTableOffset uint16    /*points to ATOM_Tonga_MCLK_Dependency_Table */
	SclkDependencyTableOffset uint16    /*points to ATOM_Tonga_SCLK_Dependency_Table */
	VddcLookupTableOffset     uint16    /*points to ATOM_Tonga_Voltage_Lookup_Table */
	VddgfxLookupTableOffset   uint16    /*points to ATOM_Tonga_Voltage_Lookup_Table */
	MMDependencyTableOffset   uint16    /*points to ATOM_Tonga_MM_Dependency_Table */
	VCEStateTableOffset       uint16    /*points to ATOM_Tonga_VCE_State_Table uint16 */
	PPMTableOffset            uint16    /*points to ATOM_Tonga_PPM_Table */
	PowerTuneTableOffset      uint16    /*points to ATOM_PowerTune_Table */
	HardLimitTableOffset      uint16    /*points to ATOM_Tonga_Hard_Limit_Table */
	PCIETableOffset           uint16    /*points to ATOM_Tonga_PCIE_Table */
	GPIOTableOffset           uint16    /*points to ATOM_Tonga_GPIO_Table */
	_                         [6]uint16 /*TODO: modify reserved size to fit structure aligning */
}

type AtomTongaFanTable struct {
	RevID                   byte
	THyst                   byte
	TMin                    uint16
	TMed                    uint16
	THigh                   uint16
	PWMMin                  uint16
	PWMMed                  uint16
	PWMHigh                 uint16
	TMax                    uint16
	FanControlMode          byte
	FanPWMMax               uint16
	FanOutputSensitivity    uint16
	FanRPMMax               uint16
	MinFanSCLKAcousticLimit uint32
	TargetTemperature       byte
	MinimumPWMLimit         byte
	_                       uint16
}

type AtomFijiFanTable struct {
	RevID                   byte
	THyst                   byte
	TMin                    uint16
	TMed                    uint16
	THigh                   uint16
	PWMMin                  uint16
	PWMMed                  uint16
	PWMHigh                 uint16
	TMax                    uint16
	FanControlMode          byte
	FanPWMMax               uint16
	FanOutputSensitivity    uint16
	FanRPMMax               uint16
	MinFanSCLKAcousticLimit uint32
	TargetTemperature       byte
	MinimumPWMLimit         byte
	FanGainEdge             uint16
	FanGainHotspot          uint16
	FanGainLiquid           uint16
	FanGainVrVddc           uint16
	FanGainVrMvdd           uint16
	FanGainPlx              uint16
	FanGainHbm              uint16
	_                       uint16
}

type AtomPolarisMClkTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomPolarisMClkEntry
}

type AtomPolarisMClkEntry struct {
	VddcInd      byte
	Vddci        uint16
	VddgfxOffset uint16
	Mvdd         uint16
	Mclk         uint32
	_            uint16
}
