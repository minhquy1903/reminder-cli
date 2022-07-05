package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"golang.org/x/exp/slices"
)


type Reminder struct {
	ID       	int			`json:"id"`
	Title    	string		`json:"title"`
	Message  	string		`json:"message"`
	Repeat		string		`json:"repeat"`
	Handled		bool		`json:"handled"`
}

type Reminders []Reminder

// Create calls the create API endpoint
func (rs *Reminders) Add(title string, message string, repeat string) {
	
	rmd := Reminder{
		ID: len(*rs) + 1,
		Title:    title,
		Message:  message,
		Repeat: repeat,
		Handled: false,
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

