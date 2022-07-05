package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/exp/slices"
)


type Reminder struct {
	ID       int        `json:"id"`
	Title    string        `json:"title"`
	Message  string        `json:"message"`
	Duration time.Duration `json:"duration"`
}

type Reminders []Reminder

// Create calls the create API endpoint
func (rs *Reminders) Add(title string, message string, duration time.Duration) {
	rmd := Reminder{
		Title:    title,
		Message:  message,
		Duration: duration,
	}

	*rs = append(*rs, rmd)

	data, err := json.Marshal(rs)

	if err != nil {
		fmt.Println(err.Error())
	}

	ioutil.WriteFile("db.json", data, 0644)
}

func (rs *Reminders) Load() {
	file, err := ioutil.ReadFile("db.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, rs)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func (rs *Reminders) Delete(id int) error {
	
	cID := slices.IndexFunc(*rs, func(r Reminder) bool {
		return r.ID == id
	})

	if cID == -1 {
		return errors.New("Reminder not find")
	}
	return nil
}

