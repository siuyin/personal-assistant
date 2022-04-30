package main

import "github.com/siuyin/personal-assistant/device"

func main() {
	customHardwarePowerOn()

	select {} // wait forever
}

func customHardwarePowerOn() {
	device.PoweredOn()
}
