package client

import (
	"encoding/json"
	"io/ioutil"
	"time"
)


type Reminder struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Message  string        `json:"message"`
	Duration time.Duration `json:"duration"`
}

type Reminders []Reminder

// Create calls the create API endpoint
func (rs *Reminders) Add(title, message string, duration time.Duration) error {
	rmd := Reminder{
		Title:    title,
		Message:  message,
		Duration: duration,
	}
	
	data, err := json.Marshal(rs)

	if err != nil {
		return err
	}

	return ioutil.WriteFile("", data, 0644)
}

func (rs *Reminders) Load() error {
	file, err := ioutil.ReadFile("")

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, rs)

	if err != nil {
		return err
	}

	return nil
}

