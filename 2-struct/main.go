package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (h *user) helloName() {
	fmt.Printf("Hello, %s\n", h.Name)
}

func (h *user) userDetails() {
	const dateFormat = "January 2, 2006"
	DOB, err := time.Parse(dateFormat, h.DOB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s who was born in %s would be %d years old today.\n", h.Name, h.City, time.Now().Year()-DOB.Year())
	}
}

func main() {
	u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
	u.helloName()
	u.userDetails()
}
