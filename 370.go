package main

import (
	"fmt"
	"log"
	"time"
)

type Register struct {
	id,
	timestamp,
	action string
}

func main() {
	var registers []Register

	for {

		var continueRegistering string
		var register Register

		fmt.Println("enter the id")
		fmt.Scanln(&register.id)

		fmt.Println("enter the time (H:mm PM/AM)")
		fmt.Scanln(&register.timestamp)

		// TODO
		teste, err := time.Parse(time.Kitchen, register.timestamp)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(teste, teste.Unix())
		}

		fmt.Println("enter the action (pickup/dropoff)")
		fmt.Scanln(&register.action)

		fmt.Println("Would you like to add more records? (y/n)")
		fmt.Scanln(&continueRegistering)

		registers = append(registers, register)

		fmt.Println(registers)

		if continueRegistering != "y" {
			break
		}
	}
}
