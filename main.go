package main

import (
	"log"

	"github.com/butuzov/golang-sizeof/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Ciao!")
	}
}
