package main

import (
	"encoding/json"
	"fmt"
)

type Cloneable interface {
	Clone() (Cloneable, error)
}

type Car interface {
	Cloneable
	Drive()
}

type MyCar struct {
	Brand  string
	Color  string
	Wheels string
}

func NewMyCar(brand, color, wheels string) *MyCar {
	return &MyCar{
		Brand:  brand,
		Color:  color,
		Wheels: wheels,
	}
}

func (c *MyCar) Clone() (*MyCar, error) {
	result := &MyCar{}
	bin, _ := json.Marshal(c)
	return result, json.Unmarshal(bin, result)
}

func (c *MyCar) Drive() {
	fmt.Printf("My car is driving. It's %s brand, %s color and %s wheels.\n", c.Brand, c.Color, c.Wheels)
}

func main() {
	dadCar := NewMyCar("skoda", "red", "sport")
	dadCar.Drive()

	momCar, _ := dadCar.Clone()
	momCar.Color = "white"
	momCar.Wheels = "comfort"
	momCar.Drive()
}