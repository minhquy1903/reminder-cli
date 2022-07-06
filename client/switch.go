package client

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Client interface {
	Add(title string, message string, isRepeat bool, at string)
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
		Error("invalid command", cmdName)
		return nil
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
		t, m, r, a := s.reminderFlags() // t = title, m = message, r = repeat, a = at

		if !ValidateTime(*a) {
			Error("input format time wrong, format required: 'hh:mm'", *a)
			return nil
		} 

		s.client.Add(*t, *m, *r, *a)
		return nil
	}
}

// reminderFlags configures reminder specific flags for a command
func (s Switch) reminderFlags() (*string, *string, *bool, *string) {
	t := flag.String("t", "Reminder", "Notification title")
	m := flag.String("m", "", "Notification message")
	r := flag.Bool("r", false, "Will repeat after duration <a>")
	a := flag.String("a", "", "A timeline or a duration with flag < -r > | Time format: 'hh:mm'")
	flag.Parse()
	return t, m, r, a
}

func ValidateTime(time string) bool {

	if !strings.Contains(time, ":") {
		return false
	}

	hh, mm, ok := SplitTime(time)

	if !ok {
		return false
	}

	if hh < 0 || hh > 23 {
		return false
	}

	if mm < 0|| mm > 60 {
		
	}

	return true
}