/*
Двойной факториал

Описать функцию Fact2(N)Fact2(N), вычисляющую двойной факториал:
N!!={ 1⋅3⋅5⋅...⋅N,N−нечетное2⋅4⋅6⋅...⋅N,N−четноеN!!={ ​1⋅3⋅5⋅...⋅N,N−нечетное2⋅4⋅6⋅...⋅N,N−четное​
С помощью этой функции найти двойные факториалы трех целых чисел.

Входные данные
Вводятся три натуральных числа A,B,C(1≤A,B,C≤16)A,B,C(1≤A,B,C≤16).

Выходные данные
Выведите двойные факториалы A,B,C.A,B,C.
*/
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Print(factTwinFind(a), factTwinFind(b), factTwinFind(c))
	
}

func factTwinFind(x int) int {
    fact := 1
	if x == 0 {
		return 1
	} else if x % 2 == 0 {
        for i := 2; i <= x; i += 2  {
			fact *= i
		}
		return fact
	} else {
		for i := 1; i <= x; i += 2  {
			fact *= i
		}
		return fact
	}
}
// Решено