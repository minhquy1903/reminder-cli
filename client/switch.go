package client

import (
	"flag"
	"fmt"
	"os"
)

type Client interface {
	Add(title string, message string, repeat string)
	Load()
}

type Switch struct {
	client Client
	commands map[string] func() func(string) error
}

func NewSwitch() Switch{
	s := Switch{
		client: &Reminders{},
	}

	s.commands = map[string] func() func(string) error {
		"add": s.Add,
	}
	
	s.client.Load()
	return s
}

func (s Switch) Switch() error {
	cmdName := os.Args[len(os.Args) - 1]
	cmd, ok := s.commands[cmdName]
	
	if !ok {
		return fmt.Errorf("invalid command: '%s'", cmdName)
	}

	return cmd()(cmdName)
}

func (s Switch) Help() {
	var help string

	for name := range s.commands {
		help += name + "\t\t--help\n"
	}

	fmt.Printf("Usage of: %s: \n<command>\t[<args>]\n%s", os.Args[0], help)
}

func (s Switch) Add() func(string) error {
	return func(cmd string) error {
		t, m, r := s.reminderFlags()

		s.client.Add(*t, *m, *r)
		return nil
	}
}

// reminderFlags configures reminder specific flags for a command
func (s Switch) reminderFlags() (*string, *string, *string) {
	t := flag.String("t", "", "Notification title")
	m := flag.String("m", "", "Notification message")
	r := flag.String("r", "", "Will repeat in ...time")
	flag.Parse()
	return t, m, r
}