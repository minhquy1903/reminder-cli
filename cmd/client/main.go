package main

import (
	"flag"
	"fmt"
	"os"
	"reminder-cli/client"
)

var (
	helpFlag = flag.Bool("help", false, "Display a helpful message")
)

func main() {
	s := client.NewSwitch()

	err := s.Switch()

	if err != nil {
		fmt.Printf("cmd switch error: %s", err)
		os.Exit(2)
	}
}