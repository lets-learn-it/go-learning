package main

import "fmt"

type Person struct {
	Name    string
	Room    *ChatRoom
	charLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s\n", p.Name, s)
	p.charLog = append(p.charLog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}
func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(sender, message string) {
	for _, p := range c.people {
		if p.Name != sender {
			p.Receive(sender, message)
		}
	}
}

func (c *ChatRoom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	room.Join(john)

	jane := NewPerson("Jane")
	room.Join(jane)

	jane.Say("Hi All")
	jane.PrivateMessage("John", "Hi John")
}
