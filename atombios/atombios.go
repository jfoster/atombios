package atombios

const (
	AtomRomHeaderSizeOffset    = 0x02
	AtomRomChecksumOffset      = 0x21
	AtomRomHeaderPointerOffset = 0x48
)

type AtomCommonTableHeader struct {
	StructureSize        uint16
	TableFormatRevision  byte
	TableContentRevision byte
}

type AtomRomHeader struct {
	Header                    AtomCommonTableHeader
	FirmWareSignature         [4]byte
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

type AtomMasterCommandTable struct {
	Header                          AtomCommonTableHeader
	ASICInit                        uint16
	GetDisplaySurfaceSize           uint16
	ASICRegistersInit               uint16
	VRAMBlockVenderDetection        uint16
	DIGxEncoderControl              uint16
	MemoryControllerInit            uint16
	EnableCRTCMemReq                uint16
	MemoryParamAdjt                 uint16
	DVOEncoderControl               uint16
	GPIOPinControl                  uint16
	SetEngineClock                  uint16
	SetMemoryClock                  uint16
	SetPixelClock                   uint16
	EnableDispPowerGating           uint16
	ResetMemoryDLL                  uint16
	ResetMemoryDevice               uint16
	MemoryPLLInit                   uint16
	AdjtDisplayPll                  uint16
	AdjtMemoryController            uint16
	EnableASICStaticPwrMgt          uint16
	SetUniphyInstance               uint16
	DACLoadDetection                uint16
	LVTMAEncoderControl             uint16
	HWMiscOperation                 uint16
	DAC1EncoderControl              uint16
	DAC2EncoderControl              uint16
	DVOOutputControl                uint16
	CV1OutputControl                uint16
	GetConditionalGoldenSetting     uint16
	TVEncoderControl                uint16
	PatchMCSetting                  uint16
	MCSEQControl                    uint16
	GfxHarvesting                   uint16
	EnableScaler                    uint16
	BlankCRTC                       uint16
	EnableCRTC                      uint16
	GetPixelClock                   uint16
	EnableVGARender                 uint16
	GetSCLKOverMCLKRatio            uint16
	SetCRTCTiming                   uint16
	SetCRTCOverScan                 uint16
	SetCRTCReplication              uint16
	SelectCRTCSource                uint16
	EnableGraphSurfaces             uint16
	UpdateCRTCDoubleBufferRegisters uint16
	LUTAutoFill                     uint16
	EnableHWIconCursor              uint16
	GetMemoryClock                  uint16
	GetEngineClock                  uint16
	SetCRTCingDTDTiming             uint16
	ExternalEncoderControl          uint16
	LVTMAOutputControl              uint16
	VRAMBlockDetectionByStrap       uint16
	MemoryCleanUp                   uint16
	ProcessI2cChannelTransaction    uint16
	WriteOneByteToHWAssistedI2C     uint16
	ReadHWAssistedI2CStat           uint16
	SpeedFanControl                 uint16
	PowerConnectorDetection         uint16
	MCSynchronization               uint16
	ComputeMemoryEnginePLL          uint16
	MemoryRefreshConversion         uint16
	VRAMGetCurrentInfoBlock         uint16
	DynamicMemorySettings           uint16
	MemoryTraining                  uint16
	EnableSpreadSpectrumOnPPLL      uint16
	TMDSAOutputControl              uint16
	SetVoltage                      uint16
	DAC1OutputControl               uint16
	DAC2OutputControl               uint16
	ComputeMemoryClockParam         uint16
	ClockSource                     uint16
	MemoryDeviceInit                uint16
	GetDispObjectInfo               uint16
	DIG1EncoderControl              uint16
	DIG2EncoderControl              uint16
	DIG1TransmitterControl          uint16
	DIG2TransmitterControl          uint16
	ProcessAuxChannelTransaction    uint16
	DPEncoderService                uint16
	GetVoltageInfo                  uint16
}

type AtomMasterDataTable struct {
	Header                   AtomCommonTableHeader
	UtilityPipeLine          uint16
	MultimediaCapabilityInfo uint16
	MultimediaConfigInfo     uint16
	StandardVESATiming       uint16
	FirmwareInfo             uint16
	PaletteData              uint16
	LCDInfo                  uint16
	DIGTransmitterInfo       uint16
	AnalogTVInfo             uint16
	SupportedDevicesInfo     uint16
	GPIOI2CInfo              uint16
	VRAMageByFirmware        uint16
	GPIOPinLUT               uint16
	VESAToInternalModeLUT    uint16
	ComponentVideoInfo       uint16
	PowerPlayInfo            uint16
	CompassionateData        uint16
	SaveRestoreInfo          uint16
	PPLLSSInfo               uint16
	OemInfo                  uint16
	XTMDSInfo                uint16
	MclkSSInfo               uint16
	ObjectHeader             uint16
	IndirectIOAccess         uint16
	MCInitParameter          uint16
	ASICVDDCInfo             uint16
	ASICInternalSSInfo       uint16
	TVVideoMode              uint16
	VRAMInfo                 uint16
	MemoryTrainingInfo       uint16
	IntegratedSystemInfo     uint16
	ASICProfilingInfo        uint16
	VoltageObjectInfo        uint16
	PowerSourceInfo          uint16
}

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

type AtomVEGA10PowerPlayTable struct {
	Header                       AtomCommonTableHeader
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

type AtomTongaPowerTuneTable struct {
	RevId                     byte
	TDP                       uint16
	ConfigurableTDP           uint16
	TDC                       uint16
	BatteryPowerLimit         uint16
	SmallPowerLimit           uint16
	LowCACLeakage             uint16
	HighCACLeakage            uint16
	MaximumPowerDeliveryLimit uint16
	TjMax                     uint16
	PowerTuneDataSetID        uint16
	EDCLimit                  uint16
	SoftwareShutdownTemp      uint16
	ClockStretchAmount        uint16
	_                         [2]uint16
}

type AtomFijiPowerTuneTable struct {
	RevId                     byte
	TDP                       uint16
	ConfigurableTDP           uint16
	TDC                       uint16
	BatteryPowerLimit         uint16
	SmallPowerLimit           uint16
	LowCACLeakage             uint16
	HighCACLeakage            uint16
	MaximumPowerDeliveryLimit uint16
	TjMax                     uint16 /* For Fiji, this is also TemperatureLimitEdge */
	PowerTuneDataSetID        uint16
	EDCLimit                  uint16
	SoftwareShutdownTemp      uint16
	ClockStretchAmount        uint16
	TemperatureLimitHotspot   uint16 /*The following are added for Fiji */
	TemperatureLimitLiquid1   uint16
	TemperatureLimitLiquid2   uint16
	TemperatureLimitVrVddc    uint16
	TemperatureLimitVrMvdd    uint16
	TemperatureLimitPlx       uint16
	Liquid1I2Caddress         byte /*Liquid */
	Liquid2I2Caddress         byte
	LiquidI2CLine             byte
	VrI2Caddress              byte /*VR */
	VrI2CLine                 byte
	PlxI2Caddress             byte /*PLX */
	PlxI2CLine                byte
	_                         uint16
}

// typedef struct _ATOM_Tonga_Fan_Table {
// 	UCHAR   ucRevId;						 /* Change this if the table format changes or version changes so that the other fields are not the same. */
// 	UCHAR   ucTHyst;						 /* Temperature hysteresis. Integer. */
// 	HORT  TMin; 						 /* The temperature, in 0.01 centigrades, below which we jt run at a minimal PWM. */
// 	HORT  TMed; 						 /* The middle temperature where we change slopes. */
// 	HORT  THigh;						 /* The high point above TMed for adjting the second slope. */
// 	HORT  PWMMin;						 /* The minimum PWM value in percent (0.01% increments). */
// 	HORT  PWMMed;						 /* The PWM value (in percent) at TMed. */
// 	HORT  PWMHigh;						 /* The PWM value at THigh. */
// 	HORT  TMax; 						 /* The max temperature */
// 	UCHAR   ucFanControlMode;				  /* Legacy or Fuzzy Fan mode */
// 	HORT  FanPWMMax;					  /* Maximum allowed fan power in percent */
// 	HORT  FanOutputSensitivity;		  /* Sensitivity of fan reaction to temepature changes */
// 	HORT  FanRPMMax;					  /* The default value in RPM */
// 	ULONG  ulMinFanSCLKAcoticLimit;	   /* Minimum Fan Controller SCLK Frequency Acotic Limit. */
// 	UCHAR   ucTargetTemperature;			 /* Advanced fan controller target temperature. */
// 	UCHAR   ucMinimumPWMLimit; 			  /* The minimum PWM that the advanced fan controller can set.	This should be set to the highest PWM that will run the fan at its lowest RPM. */
// 	HORT  Reserved;
// } ATOM_Tonga_Fan_Table;

// typedef struct _ATOM_Fiji_Fan_Table {
// 	UCHAR   ucRevId;						 /* Change this if the table format changes or version changes so that the other fields are not the same. */
// 	UCHAR   ucTHyst;						 /* Temperature hysteresis. Integer. */
// 	HORT  TMin; 						 /* The temperature, in 0.01 centigrades, below which we jt run at a minimal PWM. */
// 	HORT  TMed; 						 /* The middle temperature where we change slopes. */
// 	HORT  THigh;						 /* The high point above TMed for adjting the second slope. */
// 	HORT  PWMMin;						 /* The minimum PWM value in percent (0.01% increments). */
// 	HORT  PWMMed;						 /* The PWM value (in percent) at TMed. */
// 	HORT  PWMHigh;						 /* The PWM value at THigh. */
// 	HORT  TMax; 						 /* The max temperature */
// 	UCHAR   ucFanControlMode;				  /* Legacy or Fuzzy Fan mode */
// 	HORT  FanPWMMax;					  /* Maximum allowed fan power in percent */
// 	HORT  FanOutputSensitivity;		  /* Sensitivity of fan reaction to temepature changes */
// 	HORT  FanRPMMax;					  /* The default value in RPM */
// 	ULONG  ulMinFanSCLKAcoticLimit;		/* Minimum Fan Controller SCLK Frequency Acotic Limit. */
// 	UCHAR   ucTargetTemperature;			 /* Advanced fan controller target temperature. */
// 	UCHAR   ucMinimumPWMLimit; 			  /* The minimum PWM that the advanced fan controller can set.	This should be set to the highest PWM that will run the fan at its lowest RPM. */
// 	HORT  FanGainEdge;
// 	HORT  FanGainHotspot;
// 	HORT  FanGainLiquid;
// 	HORT  FanGainVrVddc;
// 	HORT  FanGainVrMvdd;
// 	HORT  FanGainPlx;
// 	HORT  FanGainHbm;
// 	HORT  Reserved;
// } ATOM_Fiji_Fan_Table;
