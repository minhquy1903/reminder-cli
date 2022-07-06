package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Reminder struct {
	ID       	int			`json:"id"`
	Title    	string		`json:"title"`
	Message  	string		`json:"message"`
	IsRepeat	bool		`json:"isRepeat"`
	At			string		`json:"At"`	
	Handled		bool		`json:"handled"`
}

// Generate the spec string
func (r Reminder) GetSpec() string {
	hh, mm, _ := SplitTime(r.At)
	return fmt.Sprintf("%v %v * * *", mm, hh)
}

func (r Reminder) GetDuration() time.Duration {
	hh,mm, _ := SplitTime(r.At)
	return time.Duration(hh * int(time.Hour) + mm * int(time.Minute))
}

func SplitTime(time string) (hh int, mm int, ok bool) {
	t := strings.Split(time, ":")

	hh, err := strconv.Atoi(t[0]) 

	if err != nil {
		return -1, -1, false
	}

	mm, err = strconv.Atoi(t[1]) 

	if err != nil {
		return -1, -1, false
	}

	return hh, mm, true
}

type Reminders []Reminder

// Create calls the create API endpoint
func (rs *Reminders) Add(title string, message string, isRepeat bool, at string) {
	
	rmd := Reminder{
		ID: len(*rs) + 1,
		Title:    title,
		Message:  message,
		IsRepeat: isRepeat,
		At: at,
		Handled: false,
	}

	*rs = append(*rs, rmd)

	data, err := json.Marshal(rs)

	if err != nil {
		fmt.Println(err.Error())
	}

	ioutil.WriteFile(GetDBPath(), data, 0644)
}

// Load data from the db.json file
func (rs *Reminders) Load() {
	file, err := ioutil.ReadFile(GetDBPath())

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, rs)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetDBPath() string {
	
	dirPath, err := os.Getwd()

	if err != nil {
		fmt.Println("Can not get dir path")
	}

	return dirPath + "/db.json"
}
