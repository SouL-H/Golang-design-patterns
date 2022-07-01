package main

import (
	"fmt"
)

// ########## BUTTON ##########

type button struct {
	cmd command
}

func (b *button) press() {
	fmt.Println("press btn")
	b.cmd.execute()
}

// ########## COMMAND ##########

type command interface {
	execute()
}

// ########## OFF COMMAND ##########

type offCommand struct {
	device device
}

func (c offCommand) execute() {
	fmt.Println("do off command")
	c.device.off()
}

// ########## ON COMMAND ##########

type onCommand struct {
	device device
}

func (c onCommand) execute() {
	fmt.Println("do on command")
	c.device.on()
}

// ########## DEVICE ##########

type device interface {
	on()
	off()
}

// ########## TV ##########

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	fmt.Println("tv on")
	t.isRunning = true
}

func (t *tv) off() {
	fmt.Println("tv off")
	t.isRunning = false
}

// ########## MAIN ##########

func main() {
	tv := &tv{}

	onCmd := onCommand{device: tv}
	offCmd := offCommand{device: tv}

	onBtn := button{cmd: onCmd}
	offBtn := button{cmd: offCmd}

	onBtn.press()
	offBtn.press()
}
