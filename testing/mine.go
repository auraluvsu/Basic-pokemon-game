package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println("Error processing operation", err)
	} else {
		log.Println("Quotient is:", result)
	}
}

func divide(x, y float32) (float32, error) {
	var result float32

	if x < y {
		return 0, errors.New("Divisor is bigger than dividend")
	} else if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	result = x / y
	return result, nil
}
