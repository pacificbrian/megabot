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
	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/mqtt"
	"gobot.io/x/gobot/v2/platforms/sparki"
)

func main() {
	var connections []gobot.Connection

	mqttAdaptor := mqtt.NewAdaptor("tcp://0.0.0.0:1883", "megabotMQTT")
	connections = append(connections, mqttAdaptor)

	sparkiAdaptor := sparki.NewAdaptor(os.Args[1])
	connections = append(connections, sparkiAdaptor)
	led := gpio.NewLedDriver(sparkiAdaptor, "13")

	work := func() {
		selfPublish := false

		mqttAdaptor.On("toggleLED", func(msg mqtt.Message) {
			log.Println(msg)
			led.Toggle()
		})

		if selfPublish {
			gobot.Every(3*time.Second, func() {
				brightness := byte(100)
				mqttAdaptor.Publish("toggleLED",
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
