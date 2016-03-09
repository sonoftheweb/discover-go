package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)
	m := 50

	if randomNumber > m {
		fmt.Printf("my random number is %d and is greater than %d\n", randomNumber, m)
	} else {
		fmt.Printf("my random number is %d and is NOT greater than %d\n", randomNumber, m)
	}

	fmt.Println("I am a student of the Holberton School")
	fmt.Println("It's a beautiful weather!")
	fmt.Println("Rudy Rigot is a founder")
	fmt.Println("Sylvain Kalache is a founder")
	fmt.Println("Julien Barbier is a founder")
}
