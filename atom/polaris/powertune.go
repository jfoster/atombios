package atom

type AtomPowerTuneTable struct {
	RevID                     byte
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
	TemperatureLimitHotspot   uint16
	TemperatureLimitLiquid1   uint16
	TemperatureLimitLiquid2   uint16
	TemperatureLimitVrVddc    uint16
	TemperatureLimitVrMvdd    uint16
	TemperatureLimitPlx       uint16
	Liquid1I2CAddress         byte
	Liquid2I2CAddress         byte
	LiquidI2CLine             byte
	VrI2CAddress              byte
	VrI2CLine                 byte
	PlxI2CAddress             byte
	PlxI2CLine                byte
	_                         uint16
}
