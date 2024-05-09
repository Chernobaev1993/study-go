/*
Количество минимумов

Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N(N≤105)N(N≤105), а затем NN чисел. Все числа по модулю не превышают 106106.

Выходные данные
Выведите количество минимальных элементов.
*/

package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	sli := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&sli[i])
	}
	minV := sli[0]
	count := 0

	for i := 1; i < size; i++ {
		if sli[i] < minV {
			minV = sli[i]
		}
	}

	for i := 0; i < size; i++ {
		if sli[i] == minV {
			count++
		}
	}
	fmt.Println(count)
}

// решено
