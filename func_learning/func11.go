/*
Проверка на простоту
Проверьте, является ли число простым.

Входные данные
Вводится одно натуральное число nn, принимающие значения от 22 до 2⋅1052⋅105.
Выходные данные

Необходимо вывести "prime", если число простое, или "composite", если число составное.
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	z := isPrime(n)
	if z == true {
		fmt.Println("prime")
	} else {
		fmt.Println("composite")
	}
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x % i == 0 {
			return false
		}
	}
	return true
}
// Решено