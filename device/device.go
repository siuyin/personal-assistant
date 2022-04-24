// Package device provides device abstractions.
// This is also known as Hardware Abstraction Layer.
package device

// PoweredOn is published when power to device is detected.
func PoweredOn() {
}

// ShutdownInitiated is published with a user or devices requests to shut down the device.
func ShutdownInitiated() {
}

// ConnectionAvailable is published when a means to connect to the internet is available.
func ConnectionAvailable() {
}

//go:generate stringer -type=Capability

// Capability indicates what the device is capable of.
type Capability int

// These list the potential capabilities available for a device.
const (
	VideoCapture Capability = iota
	AudioCapture
	ApplicationMonitoring
	Display
	AudioOutput
)

// Capabilities are the list of capabilities a device has.
type Capabilities []Capability

// Describe returns the capabilities available for a device.
func Describe() Capabilities {
	return []Capability{}
}

// Unit represents a physical device unit / box / item.
type Unit struct {
	ID   string
	Desc string
	Capabilities
}

// List returns a list of device Units.
func List() []Unit {
	return []Unit{}
}

// Register enters a unit into the device database.
func Register(u Unit) error {
	return nil
}
