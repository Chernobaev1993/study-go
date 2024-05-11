/*
Произведение количества разрядов двух чисел

Дано два натуральных числа. Найдите количество разрядов каждого из них и выведите их произведение.

Входные данные
Вводятся два натуральных числа - n,m,n,m, каждое из которых не превосходит 109109.

Выходные данные
Выведите произведение количества разрядов данных чисел.
*/
package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	
	fmt.Println(razFind(n) * razFind(m))
}

func razFind(x int) int {
    count := 0
	for x > 0 {
		count++
		x = x / 10
	}
	return count
}
// Решено