package tonga

type AtomPowerTuneTable struct {
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
