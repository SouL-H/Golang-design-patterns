

package main

import "fmt"

type Color string
type Wheels string

type CarBuilder interface {
	Color(color Color) CarBuilder
	Wheels(wheels Wheels) CarBuilder
	Build() Car
}

type Car interface {
	Drive()
}

type MyCarBuilder struct {
	color  Color
	wheels Wheels
}

type MyCar struct {
	color  Color
	wheels Wheels
}

const (
	Red   Color = "read"
	White       = "white"
)

const (
	Sport   Wheels = "sport"
	Comfort        = "comfort"
)

func NewMyCarBuilder() CarBuilder {
	return &MyCarBuilder{}
}

func (c *MyCar) Drive() {
	fmt.Printf("My car is driving with color %s and wheels %s\n", c.color, c.wheels)
}

func (b *MyCarBuilder) Color(color Color) CarBuilder {
	b.color = color
	return b
}

func (b *MyCarBuilder) Wheels(wheels Wheels) CarBuilder {
	b.wheels = wheels
	return b
}

func (b *MyCarBuilder) Build() Car {
	return &MyCar{
		color:  b.color,
		wheels: b.wheels,
	}
}

func main() {
	dadCar := NewMyCarBuilder().Color(Red).Wheels(Sport).Build()
	momCar := NewMyCarBuilder().Color(White).Wheels(Comfort).Build()
	dadCar.Drive()
	momCar.Drive()
}