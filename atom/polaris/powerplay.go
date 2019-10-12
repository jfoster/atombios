package polaris

type AtomMClkTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomMClkEntry
}

type AtomMClkEntry struct {
	VddcInd      byte
	Vddci        uint16
	VddgfxOffset uint16
	Mvdd         uint16
	Mclk         uint32
	_            uint16
}
