package atom

type AtomRomPCIInfo struct {
	Signature string `struct:"[4]byte"`
	VendorID  uint16
	DeviceID  uint16
}
