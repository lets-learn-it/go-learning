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
