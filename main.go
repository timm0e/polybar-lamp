package main

import (
	"bufio"
	"flag"
	"fmt"
	"polybar-lamp/config"
	"polybar-lamp/hass"
	"polybar-lamp/lampcontroller"
	"polybar-lamp/polybar"
	"net"
	"os"
	"strings"
)

const socketPath = "/tmp/polybar-lamp.sock"

func main() {
	flag.Parse()
	action := flag.Arg(0)

	switch strings.ToLower(action) {
	case "makeconfig":
		config.WriteInitialConfig()
	case "server":
		server()
	case "sendcommand":
		sendCommand(flag.Arg(1))
	default:
		if len(action) == 0 {
			fmt.Println("Missing action")
		} else {
			fmt.Printf("Unknown action %s\n", action)
		}
		fmt.Printf("Usage: %s <action>\n", os.Args[0])
		fmt.Println("Actions: makeconfig, server, sendcommand <command>")
		fmt.Println("Commands: onoff, switchmode, increment, decrement")
	}
}

func server() {
	os.Remove(socketPath)
	sock, err := net.Listen("unix", socketPath)

	if err != nil {
		panic(err)
	}

	lampConfig := config.GetPolybarLampConfig()

	apiclient := hass.NewAPIClient(lampConfig)

	lampController, _ := lampcontroller.NewLampController(lampConfig.LampEntity, apiclient, lampConfig.ColorTempLimitOverride)

	isOn, err := lampController.GetOnState()
	if isOn {
		polybar.SwitchState(polybar.BRIGHTNESS, lampController)
	} else {
		polybar.SwitchState(polybar.OFF, lampController)
	}

	for {
		conn, err := sock.Accept()
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			text := scanner.Text()

			switch strings.ToLower(text) {
			case "onoff":
				polybar.OnOff(lampController)
			case "switchmode":
				polybar.SwitchMode(lampController)
			case "increment":
				polybar.Increment(lampController)
			case "decrement":
				polybar.Decrement(lampController)
			}
		}
	}
}

func sendCommand(command string) {
	sock, err := net.Dial("unix", socketPath)
	if err != nil {
		panic(err)
	}

	_, err = sock.Write([]byte(command))
	if err != nil {
		panic(err)
	}

	err = sock.Close()
	if err != nil {
		panic(err)
	}
}
