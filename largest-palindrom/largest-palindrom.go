package main

import (
	"fmt"
)

func isPalindrom(number int) bool {
	var revNumber int
	x := number

	for {
		if x != 0 {
			y := x % 10
			x /= 10

			revNumber *= 10
			revNumber += y

			continue
		}
		break
	}

	return revNumber == number
}

func main() {
	var palindrom int
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isPalindrom(i * j) {
				if palindrom == 0 {
					palindrom = i * j
					continue
				}

				if palindrom < i*j {
					palindrom = i * j
				}
			}
		}
	}

	fmt.Println("Largest palindrom is", palindrom)
}
