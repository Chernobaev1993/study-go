package main

import (
	"fmt"
	"math"
)

func main() {
	var n, z int
	fmt.Scan(&n)

	var m float64
	m = math.Sqrt(float64(n))
	z = int(m)

	for i := 2; i <= z; i++ {
		for n%i == 0 {
			fmt.Print(i, " ")
			n /= i

		}
	}

	if n > 1 {
		fmt.Print(n)
	}

}
