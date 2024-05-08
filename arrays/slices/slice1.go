package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0

	slice1 := make([]int, n)
	for i := 0; i < len(slice1); i++ {
		fmt.Scan(&slice1[i])
	}
	for i := 1; i < len(slice1); i++ {
		if slice1[i-1] < slice1[i] {
			count++
		}
	}
	fmt.Println(count)
}
