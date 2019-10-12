package fiji

type AtomFanTable struct {
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
