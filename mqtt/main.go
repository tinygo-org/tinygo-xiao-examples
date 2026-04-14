package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	mqtt "github.com/soypat/natiu-mqtt"
	"tinygo.org/x/drivers/netdev"
	nl "tinygo.org/x/drivers/netlink"
	link "tinygo.org/x/espradio/netlink"
)

var (
	ssid     string
	password string
	broker   string = "broker.hivemq.com:1883"
	topic    string = "cpu/usage"
)

func main() {
	initDisplay()

	// wait a bit for serial
	time.Sleep(2 * time.Second)

	if err := connectToWifi(); err != nil {
		printMessage(err.Error())
		return
	}

	clientId := "tinygo-client-" + randomString(10)
	printMessage("ClientId: " + clientId)

	// Get a transport for MQTT packets.
	// Retry TCP connection since public brokers may reject/close connections under load.
	printMessage("Connecting to " + broker)
	var conn net.Conn
	for attempt := range 5 {
		var err error
		conn, err = net.Dial("tcp", broker)
		if err != nil {
			printMessage(fmt.Sprintf("net.Dial attempt %d failed: %s", attempt+1, err))
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	if conn == nil {
		printMessage("TCP connection failed")
		return
	}
	printMessage(fmt.Sprintf("TCP connected to %v", conn.RemoteAddr()))
	defer conn.Close()

	// Create new client
	client := mqtt.NewClient(mqtt.ClientConfig{
		Decoder: mqtt.DecoderNoAlloc{make([]byte, 1500)},
		OnPub: func(_ mqtt.Header, _ mqtt.VariablesPublish, r io.Reader) error {
			message, _ := io.ReadAll(r)
			printMessage(fmt.Sprintf("%s - %s", string(message), topic))
			return nil
		},
	})

	// Connect client
	var varconn mqtt.VariablesConnect
	varconn.SetDefaultMQTT([]byte(clientId))
	varconn.KeepAlive = 60 // seconds; some brokers reject KeepAlive=0
	printMessage("MQTT CONNECT...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx, conn, &varconn)
	if err != nil {
		printMessage("failed to connect: " + err.Error())
		return
	}
	printMessage("MQTT CONNECT succeeded")

	// Subscribe to topic
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Subscribe(ctx, mqtt.VariablesSubscribe{
		PacketIdentifier: 23,
		TopicFilters: []mqtt.SubscribeRequest{
			{TopicFilter: []byte(topic), QoS: mqtt.QoS0},
		},
	})
	if err != nil {
		printMessage("failed to subscribe to " + topic + ": " + err.Error())
		return
	}
	printMessage("Subscribed to " + topic)

	// Publish on topic
	pubFlags, _ := mqtt.NewPublishFlags(mqtt.QoS0, false, false)
	pubVar := mqtt.VariablesPublish{
		TopicName: []byte(topic),
	}

	for i := 0; i < 10; i++ {
		if !client.IsConnected() {
			printMessage("client disconnected: " + client.Err().Error())
		}

		payload := fmt.Sprintf("Random value: %d\n", randomInt(0, 100))

		pubVar.PacketIdentifier++
		err = client.PublishPayload(pubFlags, pubVar, []byte(payload))
		if err != nil {
			printMessage("transmit error: " + err.Error())
		}

		time.Sleep(time.Second)

		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		err = client.HandleNext()
		if err != nil {
			printMessage("handle next: " + err.Error())
		}

	}

	client.Disconnect(errors.New("disconnected gracefully"))
	printMessage("Disconnected from MQTT broker.")

	for {
		select {}
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func connectToWifi() error {
	link := link.Esplink{}
	netdev.UseNetdev(&link)

	printMessage("Connecting to WiFi...")
	for range 3 {
		err := link.NetConnect(&nl.ConnectParams{
			Ssid:       ssid,
			Passphrase: password,
		})
		if err == nil {
			printMessage("Connected to WiFi.")
			return nil
		}
		printMessage("Failed to connect to WiFi: " + err.Error())
		time.Sleep(5 * time.Second)
	}

	return errors.New("failed to connect to WiFi after 3 attempts")
}
