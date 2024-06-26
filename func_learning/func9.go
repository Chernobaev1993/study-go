/*
Сравнение количества цифр

Даны два натуральных числа. Выяснить, в каком из них больше цифр.

Входные данные
Вводятся два натуральных числа, каждый из которых не превосходит 109109.

Выходные данные
Выведите цифру 1, если количество цифр в первом числе больше, чем количество цифр во втором числе.
Выведите цифру 2, если количество цифр во втором числе больше, чем количество цифр в первом числе.
Выведите цифру 0, если количество цифр в первом числе равно количеству цифр во втором числе.
*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	a1 := howManyC(a)
	b1 := howManyC(b)

	if a1 > b1 {
		fmt.Println(1)
	} else if b1 > a1 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
	
}

func howManyC(x int) int {
    count := 0
	for x > 0 {
		count++
		x = x / 10
	}
	return count
	
}
// Решено