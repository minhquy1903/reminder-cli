package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reminder-cli/client"
)

const dbURL = "db.json"

func main() {

	rs := Load()

	fmt.Println(rs)
}

func Load() client.Reminders{
	file, err := ioutil.ReadFile("db.json")

	rs := &client.Reminders{}

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, rs)

	if err != nil {
		fmt.Println(err.Error())
	}

	return *rs
}