package polaris

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
