package main

import "fmt"

type Switch struct {
	stat State
}

func (s *Switch) on() {
	s.stat.on(s)
}

func (s *Switch) off() {
	s.stat.off(s)
}

type State interface {
	on(sw *Switch)
	off(sw *Switch)
}

type OnState struct {
}

func (s *OnState) on(sw *Switch) {
	fmt.Println("already turn on")
}

func (s *OnState) off(sw *Switch) {
	fmt.Println("turn off the switch")
	sw.stat = &OffState{}
}

type OffState struct {
}

func (s *OffState) on(sw *Switch) {
	fmt.Println("turn on the switch")
	sw.stat = &OnState{}
}

func (s *OffState) off(sw *Switch) {
	fmt.Println("already turn off")
}

func main() {
	sw := &Switch{stat: &OffState{}}
	sw.on()
	sw.off()
	sw.off()
}