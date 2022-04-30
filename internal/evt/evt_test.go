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
	sub, err := js.PullSubscribe("pa.device.test", "myClientTest")
	if err != nil {
		t.Error(err)
	}

	ms, err := sub.Fetch(1)
	if err != nil {
		t.Error("fetch: ", err)
	}
	ms[0].Ack()
	if dat := string(ms[0].Data); dat != expected {
		t.Errorf("expected %s, but got %s", expected, dat)
	}
}
