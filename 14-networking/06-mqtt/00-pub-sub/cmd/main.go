package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pskp-95/mqtt-pub-sub/internal/types"
)

var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "127.0.0.1"
	var port = 1883
	var topic = "topic/test"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client, topic)
	publish(client, topic)

	client.Disconnect(250)
}

func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, messageSubHandler)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}

func publish(client mqtt.Client, topic string) {
	num := 10
	for i := 0; i < num; i++ {

		entity := types.Entity{
			Type: "Motor",
			Facets: map[string]types.Facet{
				"rotation": map[string]any{
					"speed":     30,
					"direction": "clockwise",
				},
			},
		}

		buffer := bytes.NewBuffer([]byte{})
		json.NewEncoder(buffer).Encode(entity)

		token := client.Publish(topic, 0, false, buffer.Bytes())
		token.Wait()
		time.Sleep(time.Second)
	}
}
