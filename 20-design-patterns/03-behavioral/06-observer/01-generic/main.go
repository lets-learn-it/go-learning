package main

// Eg.
// Observable - Patient
// Observer - Doctor
// Doctor want to observe Patient

func main() {
	person := NewPerson("BuriBuriZaemon")
	doctor := DoctorService{}

	person.Subscribe(doctor)

	person.CatchACold()
}
