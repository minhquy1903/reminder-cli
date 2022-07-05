package main

import (
	"flag"
	"fmt"
	"os"
	"reminders-cli/client"
)

var (
	helpFlag = flag.Bool("help", false, "Display a helpful message")
)

func main() {
	flag.Parse()

	fmt.Println(*helpFlag)

	s := client.NewSwitch()

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	
	if err != nil {
		fmt.Printf("cmd switch error: %s", err)
		os.Exit(2)
	}
}