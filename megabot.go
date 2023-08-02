/*
 * SPDX-FileCopyrightText: 2023 Brian Welty
 *
 * SPDX-License-Identifier: MPL-2.0
 */

// Pass serial port to use as the first param:
// go run megabot.go /dev/rfcomm0

package main

import (
	"log"
	"os"
	"time"
	"google.golang.org/protobuf/proto"
	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/mqtt"
	"gobot.io/x/gobot/v2/platforms/sparki"
	"github.com/pacificbrian/megabot/control"
)

func main() {
	var connections []gobot.Connection
	useRobot := false

	mqttAdaptor := mqtt.NewAdaptor("tcp://0.0.0.0:1883", "megabot-controller")
	connections = append(connections, mqttAdaptor)

	sparkiAdaptor := sparki.NewAdaptor(os.Args[1])
	led := gpio.NewLedDriver(sparkiAdaptor, "13")
	if useRobot {
		connections = append(connections, sparkiAdaptor)
	}

	work := func() {
		selfPublish := false

		mqttAdaptor.On("ctrl/ToggleLED", func(msg mqtt.Message) {
			log.Println(msg)
			led.Toggle()
		})

		mqttAdaptor.On("ctrl/Motors", func(msg mqtt.Message) {
			message := &control.MegabotCtrl{}
			proto.Unmarshal(msg.Payload(), message);
			log.Println("MQTT [Motors] Payload: ", message)
			sparkiAdaptor.Move(message.Fvalue1, message.Fvalue2, -1.0)
		})

		if selfPublish {
			gobot.Every(3*time.Second, func() {
				brightness := byte(100)
				mqttAdaptor.Publish("ctrl/ToggleLED",
						    []byte{brightness})
			})
		}
	}

	robot := gobot.NewRobot("megabot",
				connections,
				[]gobot.Device{led},
				work)

	err := robot.Start()
	if err != nil {
		log.Println(err)
	}
}
