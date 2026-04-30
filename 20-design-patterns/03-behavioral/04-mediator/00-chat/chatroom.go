package main

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
