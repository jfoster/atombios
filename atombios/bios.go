package atombios

// Bios -
type Bios struct {
	AtomRomHeader    AtomRomHeader
	AtomCommandTable AtomMasterCommandTable
	AtomDataTable    AtomMasterDataTable
	// AtomPowerplayTable  AtomPowerplayTable
	// AtomPowertuneTable  AtomPowertuneTable
	// AtomFanTable        AtomFanTable
	// AtomMClkTable       AtomMClkTable
	// AtomSClkTable       AtomSClkTable
	// AtomVoltageTable    AtomVoltageTable
	// AtomVRAMInfo        AtomVRAMInfo
	// AtomVRAMTimingEntry []AtomVRAMTimingEntry
	// AtomVRAMEntry       []AtomVRAMEntry
}
