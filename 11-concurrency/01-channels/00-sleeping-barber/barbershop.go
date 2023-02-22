package main

import (
	"fmt"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		fmt.Println(barber, " goes to waiting room to check for clients...")

		for {
			// check if there are any client
			if len(shop.ClientChan) == 0 {
				fmt.Println(barber, " taking nap...")
				isSleeping = true
			}

			// channelOpen is like shop open
			client, channelOpen := <-shop.ClientChan

			if channelOpen {
				if isSleeping {
					fmt.Println(client, " wakes up ", barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	fmt.Println(barber, " is cutting ", client, "'s hair")
	time.Sleep(shop.HairCutDuration)
	fmt.Println(barber, " is done cutting ", client, "'s hair")
}

func (shop *BarberShop) sendBarberHome(barber string) {
	fmt.Println(barber, " going home...")
	shop.BarberDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	fmt.Println("Closing shop...")

	close(shop.ClientChan)
	shop.Open = false

	for a := 1; a < shop.NumberOfBarbers; a++ {
		<-shop.BarberDoneChan
	}
	close(shop.BarberDoneChan)
	fmt.Println("Shop closed...")
}

func (shop *BarberShop) addClient(client string) {
	fmt.Println(client, " New Client...")
	if shop.Open {
		select {
		case shop.ClientChan <- client:
			fmt.Println("Added client ", client)
		default:
			fmt.Println("============Shop is full, sorry...")
		}
	} else {
		fmt.Println("Sorry, ", client, " shop is closed...")
	}
}
