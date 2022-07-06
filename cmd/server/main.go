package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reminder-cli/client"
	"time"

	"github.com/robfig/cron"
)

const (
	DB = "db.json"
	NotifierURL = "http://localhost:9000/notify"
)


func main() {

	c := cron.New()

	defer c.Stop()

	for {
		time.Sleep(1 * time.Second)
		rs := Load()

		if len(rs) != 0 {
			for _, v := range rs {
				if v.IsRepeat {
					go SetRepeatReminder(v)
				} else {
					go SetReminder(v, c)
				}

				SetDoneReminder(v.ID)
			}
		}
	}
}

// Set the reminder cronjob
func SetReminder(rmd client.Reminder, c *cron.Cron) {

	// Get the spec string
	spec := rmd.GetSpec()

	_, err := c.AddFunc(spec, func() {
		CallNotifier(rmd)
	})

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	c.Start()
	fmt.Println("Add reminder successfully")
}

func SetRepeatReminder(rmd client.Reminder) {
	// Get the spec string
	d := rmd.GetDuration()

	c := time.Tick(d)

	for n := range c {
		CallNotifier(rmd)
		fmt.Println(n)
	}

	fmt.Println("Add reminder successfully")
}

// Call to the notifier service to show the notification
func CallNotifier(body client.Reminder) {

	json_data, err := json.Marshal(body)

    if err != nil {
        fmt.Println(err.Error())
    }

	http.Post(NotifierURL, "application/json",
        bytes.NewBuffer(json_data))
}

func SetDoneReminder(id int) {
	file, err := ioutil.ReadFile(DB)

	rs := &client.Reminders{}

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, rs)

	if err != nil {
		fmt.Println(err.Error())
	}

	l := *rs

	for i, v := range l {
		if v.ID == id {
			l[i].Handled = true
		}
	}

	data, err := json.Marshal(*rs)

	if err != nil {
		fmt.Println(err.Error())
	}

	ioutil.WriteFile("db.json", data, 0644)
}

func Load() client.Reminders{
	file, err := ioutil.ReadFile(DB)

	rs := &client.Reminders{}

	l := client.Reminders{}

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, rs)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range *rs {
		if !v.Handled {
			l = append(l, v)
		}
	}

	return l
}