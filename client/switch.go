package client

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Client interface {
	Add(title string, message string, duration time.Duration)
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
	cmdName := os.Args[1]
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
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(createCmd)

		fmt.Println("hello: ",*t, *m)

		s.client.Add(*t, *m, *d)
		return nil
	}
}

// reminderFlags configures reminder specific flags for a command
func (s Switch) reminderFlags(f *flag.FlagSet) (*string, *string, *time.Duration) {
	t, m, d := "", "", time.Duration(0)
	f.StringVar(&t, "title", "", "Reminder title")
	f.StringVar(&t, "t", "", "Reminder title")
	f.StringVar(&m, "message", "", "Reminder message")
	f.StringVar(&m, "m", "", "Reminder message")
	f.DurationVar(&d, "duration", 0, "Reminder time")
	f.DurationVar(&d, "d", 0, "Reminder time")
	return &t, &m, &d
}