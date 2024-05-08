package main

import "fmt"

func main() {
	var n, z int
	fmt.Scan(&n)
	minV := n
	for i := 0; i < n; i++ {
		fmt.Scan(&z)
		if z < minV {
			minV = z
		}
	}
	fmt.Println(minV)

}
