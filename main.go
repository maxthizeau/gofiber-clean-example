package main

import (
	"fmt"
	"log"

	"github.com/maxthizeau/gofiber-clean-boilerplate/configuration"
)

func main() {
	log.Println("Work in progres...")
	config := configuration.New()
	database := configuration.NewDatabase(config)

	fmt.Println(database)
}
