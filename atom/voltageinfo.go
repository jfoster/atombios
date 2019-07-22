package atom

type AtomVoltageInfo struct {
	Header                AtomCommonTableHeader
	VDDCBaseLevel         uint16 //In number of 50mv unit
	_                     uint16 //For possible extension table offset
	NumOfVoltageEntries   byte
	BytesPerVoltageEntry  byte
	VoltageStep           byte //Indicating in how many mv increament is one step, 0.5mv unit
	DefaultVoltageEntry   byte
	VoltageControlI2cLine byte
	VoltageControlAddress byte
	VoltageControlOffset  byte
	VoltageEntries        []byte
}
