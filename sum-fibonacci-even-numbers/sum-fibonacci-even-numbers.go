package main

import "fmt"

func fibonacci(n uint32) uint32 {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var sum, n, i uint32

	for {
		n = fibonacci(i)
		i++
		if n >= 4000000 {
			break
		}

		if n%2 == 0 {
			sum += n
		}
	}

	fmt.Println(sum)
}
