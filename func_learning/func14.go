/*
Количество букв bb
Даны два предложения. Найти общее количество букв 'bb' в них.

Входные данные
Вводятся две строки, каждая из которых по длине не превосходит 10001000.

Выходные данные
Выведите  общее количество букв 'bb' в них.
*/
package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	z := findLitB(s1) + findLitB(s2)
	fmt.Println(z)
	
}

func findLitB (x string) int {
	count := 0
	for i := 0; i < len(x); i++ {
		if string(x[i]) == "b" {
			count++
		}
	}
	return count
}
// Решено