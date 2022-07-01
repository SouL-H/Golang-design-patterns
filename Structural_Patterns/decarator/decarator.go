package main

import "fmt"

type Pizza interface {
	GetPrice() int
}

type VeggieMania struct{}

func (v VeggieMania) GetPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza Pizza
}

func (t TomatoTopping) GetPrice() int {
	price := t.pizza.GetPrice()
	return price + 5
}

type CheeseTopping struct {
	pizza Pizza
}

func (c CheeseTopping) GetPrice() int {
	price := c.pizza.GetPrice()
	return price + 10
}

func main() {
	v := &VeggieMania{}
	t := &TomatoTopping{v}
	c := &CheeseTopping{t}
	fmt.Println("Price of pizza is:", c.GetPrice())
}