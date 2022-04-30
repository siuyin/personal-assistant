package evt

import (
	"fmt"
	"testing"

	"github.com/nats-io/nats.go"
)

func ExampleInit() {} // this is here just to call the init function when Mode != test. To see init messages: Mode=prod go test

func TestPub(t *testing.T) {
	Pub("pa.device", "gerbau")
}

func example_1() {
	// ci, err := js.AddConsumer("pa", &nats.ConsumerConfig{
	// 	Durable:        "myClient1",
	// 	DeliverSubject: "",
	// 	DeliverPolicy:  nats.DeliverAllPolicy,
	// 	AckPolicy:      nats.AckExplicitPolicy,
	// 	FilterSubject:  "pa",
	// })
	// if err != nil {
	// 	fmt.Println("addcons: ", err)
	// 	return
	// }
	// fmt.Println(ci)
	sub, err := js.PullSubscribe("pa", "myClient")
	if err != nil {
		fmt.Println("pullsub: ", err)
		return
	}

	ms, err := sub.Fetch(10)
	if err != nil {
		fmt.Println("fetch: ", err)
		return
	}
	for i := 0; i < len(ms); i++ {
		fmt.Println(ms[i].Data)
	}
	// Output:
	// abc
}
func notExample_2() {
	js.Subscribe("pa", func(msg *nats.Msg) {
		meta, _ := msg.Metadata()
		fmt.Printf("Stream Sequence  : %v\n", meta.Sequence.Stream)
		fmt.Printf("Consumer Sequence: %v\n", meta.Sequence.Consumer)
		fmt.Println(msg.Data)
	}, nats.DeliverAll())
	// Output:
	// abc
}
