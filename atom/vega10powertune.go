package atom

type AtomVega10PowerTuneTable struct {
	RevId                   byte
	SocketPowerLimit        uint16
	BatteryPowerLimit       uint16
	SmallPowerLimit         uint16
	TdcLimit                uint16
	EdcLimit                uint16
	SoftwareShutdownTemp    uint16
	TemperatureLimitHotSpot uint16
	TemperatureLimitLiquid1 uint16
	TemperatureLimitLiquid2 uint16
	TemperatureLimitHBM     uint16
	TemperatureLimitVrSoc   uint16
	TemperatureLimitVrMem   uint16
	TemperatureLimitPlx     uint16
	LoadLineResistance      uint16
	Liquid1I2CAddress       byte
	Liquid2I2CAddress       byte
	VrI2CAddress            byte
	PlxI2CAddress           byte
	LiquidI2CLineSCL        byte
	LiquidI2CLineSDA        byte
	VrI2CLineSCL            byte
	VrI2CLineSDA            byte
	PlxI2CLineSCL           byte
	PlxI2CLineSDA           byte
	TemperatureLimitTedge   uint16
}

type AtomVega10PowerTuneTableV2 struct {
	RevId                   byte
	SocketPowerLimit        uint16
	BatteryPowerLimit       uint16
	SmallPowerLimit         uint16
	TdcLimit                uint16
	EdcLimit                uint16
	SoftwareShutdownTemp    uint16
	TemperatureLimitHotSpot uint16
	TemperatureLimitLiquid1 uint16
	TemperatureLimitLiquid2 uint16
	TemperatureLimitHBM     uint16
	TemperatureLimitVrSoc   uint16
	TemperatureLimitVrMem   uint16
	TemperatureLimitPlx     uint16
	LoadLineResistance      uint16
	Liquid1I2CAddress       byte
	Liquid2I2CAddress       byte
	LiquidI2CLine           byte
	VrI2CAddress            byte
	VrI2CLine               byte
	PlxI2CAddress           byte
	PlxI2CLine              byte
	TemperatureLimitTedge   uint16
}

type AtomVega10PowerTuneTableV3 struct {
	RevId                   byte
	SocketPowerLimit        uint16
	BatteryPowerLimit       uint16
	SmallPowerLimit         uint16
	TdcLimit                uint16
	EdcLimit                uint16
	SoftwareShutdownTemp    uint16
	TemperatureLimitHotSpot uint16
	TemperatureLimitLiquid1 uint16
	TemperatureLimitLiquid2 uint16
	TemperatureLimitHBM     uint16
	TemperatureLimitVrSoc   uint16
	TemperatureLimitVrMem   uint16
	TemperatureLimitPlx     uint16
	LoadLineResistance      uint16
	Liquid1I2CAddress       byte
	Liquid2I2CAddress       byte
	LiquidI2CLine           byte
	VrI2CAddress            byte
	VrI2CLine               byte
	PlxI2CAddress           byte
	PlxI2CLine              byte
	TemperatureLimitTedge   uint16
	BoostStartTemperature   uint16
	BoostStopTemperature    uint16
	BoostClock              uint32
	_                       [2]uint32
}
