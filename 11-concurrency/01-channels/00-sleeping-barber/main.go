package main

import (
	"fmt"
	"math/rand"
	"time"
)

// variables
var seatingCapacity = 10
var arrivalRate = 200
var cutDuration = 2000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixMilli())

	// print welcome message
	fmt.Println("The Sleeping Barber Problem")
	fmt.Println("----------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientChan:      clientChan,
		BarberDoneChan:  doneChan,
		Open:            true,
	}

	// add barbers
	shop.addBarber("ABC")
	shop.addBarber("XYZ")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// add clients randomly with arrival rate
			randomMilliSeconds := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliSeconds)):
				go shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
