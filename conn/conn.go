// Package conn provide connection management functions.
// For example it connects a device to the internet when a Device announces that a connection is available.
package conn

// ConnectionEstablished is published when connection manager when the "best" available connection is established.
func ConnectionEstablished() {
}

// ConnectionLost is published when a device loses connectivity to the internet.
func ConnectionLost() {
}
