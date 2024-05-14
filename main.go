package main

import (
	"errors"
	"log"
)

func divide(x, y float32) (float32, error) {
	if y != 0 {
		return x / y, nil
	}

	return -1.0, errors.New("y cannot be 0")
}

func main() {
	value, err := divide(-1.0, -2.0)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("value: ", value)
}
