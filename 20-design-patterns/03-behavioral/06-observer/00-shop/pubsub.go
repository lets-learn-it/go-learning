package main

type Publisher interface {
	Subscribe(s Subscriber)
	Unsubscribe(s Subscriber)
}

type Subscriber interface {
	OnUpdate(publisher string)
	GetId() string
}
