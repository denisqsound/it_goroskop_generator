package main

import (
	"log"

	"github.com/denisqsound/it_goroskop_generator/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
