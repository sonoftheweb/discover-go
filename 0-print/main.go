package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Hello, we are Holberton School")
	fmt.Printf("the date is %s\n", t.Local())
	fmt.Printf("the year is %d\n", t.Year())
	fmt.Printf("the month is %s\n", t.Month())
	fmt.Printf("the day is %d\n", t.Day())
}
