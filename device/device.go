// Package device provides device abstractions.
// This is also known as Hardware Abstraction Layer.
package device

import (
	log "github.com/sirupsen/logrus"
	"github.com/siuyin/dflt"
	_ "github.com/siuyin/personal-assistant/command"
	"github.com/siuyin/personal-assistant/internal/evt"
	_ "github.com/siuyin/personal-assistant/learn"
)

func init() {
	if lvl := dflt.EnvString("LogLevel", "info"); lvl != "info" {
		log.SetLevel(log.DebugLevel)
	}
}

// PoweredOn is published when power to device is detected.
func PoweredOn() {
	event("CustomHardwareStarted", "aQimbo custom hardware started")
	debug("audio stream started")
	debug("video stream started")

	debug("phone detected")
	info("phone bluetooth connected")

	ConnectionAvailable()
}
func info(x ...interface{}) {
	log.WithFields(log.Fields{"module": "device"}).Info(x...)
}
func debug(x ...interface{}) {
	log.WithFields(log.Fields{"module": "device"}).Debug(x...)
}
func event(name string, msg ...interface{}) {
	evt.Pub("pa.device", name)
	info(msg...)
}

// ShutdownInitiated is published with a user or devices requests to shut down the device.
func ShutdownInitiated() {
}

// ConnectionAvailable is published when a means to connect to the internet is available.
func ConnectionAvailable() {
}

// ConnectionDropped is published when a established connection is dropped.
func ConnectionDropped() {
}

// StatusUpdated is published when a significant status change is detected in a device.
func StatusUpdated() {
}

//go:generate stringer -type=Status
// Status indicates a device state.
type Status int

// These list the potential states a device may be in.
const (
	LocalStorageAvailabilityLow Status = iota
	LocalStorageAvailabilityExhausted
	BatteryLow
	BatteryCharging
	BatteryFullyCharged
	ConnectionEstablished
	ConnectionNotEstablished
)

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

// Remove removes a unit from the device database.
func Remove(u Unit) error {
	return nil
}
