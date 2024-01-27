package main

import "fmt"

type Item struct {
	Name        string
	IsAvailable bool
	Subscribers map[string]Subscriber
}

func NewItem(name string, isAvailable bool) Item {
	return Item{
		Name:        name,
		IsAvailable: isAvailable,
		Subscribers: make(map[string]Subscriber),
	}
}

func (i *Item) UpdateAvailability(isAvailable bool) {
	i.IsAvailable = isAvailable

	// let subscribers know
	for k, v := range i.Subscribers {
		fmt.Printf("Sending update to %s\n", k)
		v.OnUpdate(i.Name)
	}
}

func (i *Item) Subscribe(s Subscriber) {
	i.Subscribers[s.GetId()] = s
}

func (i *Item) Unsubscribe(s Subscriber) {
	delete(i.Subscribers, s.GetId())
}
