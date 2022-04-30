package evt

import (
	"testing"
)

func ExampleInit() {} // this is here just to call the init function when Mode != test. To see init messages: Mode=prod go test

func TestPubSub(t *testing.T) {
	Pub("pa.device.test", "gerbau")
	checkSub(t, "gerbau")
}

func checkSub(t *testing.T, expected string) {
	sub := Sub("pa.device.test", "myClientTest")

	msg := Fetch(sub)
	if msg != expected {
		t.Errorf("expected %s, but got %s", expected, msg)
	}
}
