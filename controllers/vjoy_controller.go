package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tajtiattila/vjoy"
)

// VJoyController sends input using a fake controller managed by vJoy
type VJoyController struct {
	device *vjoy.Device
}

var buttonsToJoy = map[string]uint{
	"a":      1,
	"b":      2,
	"start":  3,
	"select": 4,
	"left":   5,
	"right":  6,
	"up":     7,
	"down":   8,
}

// NewVJoyController tries to grab vJoy 1 and creates a Controller with it
func NewVJoyController() (v *VJoyController, err error) {
	v = &VJoyController{}

	device, err := vjoy.Acquire(1)
	if err != nil {
		return nil, err
	}
	v.device = device

	return v, nil
}

// Press presses the button on the virtual controller
func (v *VJoyController) Press(button string) {
	value, ok := buttonsToJoy[strings.ToLower(button)]
	if !ok {
		fmt.Printf("Button %s doesn't seem to be valid\n", button)
	}

	v.device.Button(value).Set(true)
	v.device.Update()

	time.Sleep(200 * time.Millisecond)

	v.device.Button(value).Set(false)
	v.device.Update()
}

// InteractiveMode drops the user in a shell where they can choose which button to press
func (v *VJoyController) InteractiveMode() {
	var buttonNames []string
	var scanner = bufio.NewScanner(os.Stdin)

	for k := range buttonsToJoy {
		buttonNames = append(buttonNames, k)
	}

	for {
		fmt.Println("Available buttons:", buttonNames)
		fmt.Print("> ")
		scanner.Scan()
		s := scanner.Text()
		if s == "exit" {
			return
		}

		if !IsValidButton(s) {
			fmt.Println("Button", s, "is not valid")
			continue
		}

		for t := 5; t > 0; t-- {
			fmt.Printf("%d...\n", t)
			time.Sleep(time.Second)
		}

		v.Press(s)
	}
}
