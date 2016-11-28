// https://projecteuler.net/problem=5
package main

import (
	"fmt"
)

func main() {
	i := 1
loop:
	for {
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				i++
				continue loop
			}
		}
		fmt.Println("Smallest number that can be devided by each of the numbers from 1 to 20 is", i)
		break
	}
}
