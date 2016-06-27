package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(number int) bool {
	sqrtNumb := math.Sqrt(float64(number))

	for i := 2; i <= int(sqrtNumb); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: largest-prime-number <number>")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("%s is not a number", os.Args[1]))
	}

	sqrtNumb := math.Sqrt(float64(number))

	for i := int(sqrtNumb); i >= 0; i-- {
		if number%i == 0 && isPrime(i) {
			fmt.Println("Max prime number of", number, "is", i)
			break
		}
	}
}
