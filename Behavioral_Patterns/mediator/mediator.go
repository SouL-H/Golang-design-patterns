package main

import "fmt"

type ChatRoom struct {
	participants []*Person
}

func (c *ChatRoom) Join(p *Person) {
	if p == nil {
		return
	}

	p.Room = c
	c.participants = append(c.participants, p)
	c.Broadcast("All", fmt.Sprintf("%s is come in!", p.Name))
}

func (c *ChatRoom) Broadcast(name, msg string) {
	fmt.Printf("[%s]: %s\n", name, msg)
}

func (c *ChatRoom) Chat(sender, receiver, msg string) {
	for _, p := range c.participants {
		if p.Name == receiver {
			fmt.Printf("[%s] to [%s]: %s\n", sender, receiver, msg)
		}
	}
}

type Person struct {
	_    struct{}
	Name string
	Room *ChatRoom
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func (p *Person) Say(msg string) {
	p.Room.Broadcast(p.Name, msg)
}

func (p *Person) PrivateMessage(name, msg string) {
	p.Room.Chat(p.Name, name, msg)
}

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi everyone!")
	jane.Say("oh, hey john")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage("Simon", "glad you could join us!")
}
