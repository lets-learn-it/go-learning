package main

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	room.Join(john)

	jane := NewPerson("Jane")
	room.Join(jane)

	jane.Say("Hi All")
	jane.PrivateMessage("John", "Hi John")
}
