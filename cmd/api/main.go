package main

import (
	"eva/config"
	"eva/logging"
	"fmt"
)

func main() {
	fmt.Println("EVA server starting")
	fmt.Println("Configuring EVA server...")
	cfg := config.GetConfig()
	fmt.Println("EVA server configured.")
	fmt.Println("Setting up logger...")
	logger := logging.GetLogger()
	fmt.Println("Logger set up.")

}
