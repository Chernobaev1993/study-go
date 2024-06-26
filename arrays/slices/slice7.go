/*
Разность индексов максимума и минимума

Дан массив, состоящий из целых чисел. Напишите программу, которая выведет разность индексов максимального и минимального элементов массива. Индексация начинается с нуля.

Входные данные
Сначала задано число NN — количество элементов в массиве (1≤N≤10001≤N≤1000). Далее через пробел записаны NN различных  чисел.
Массив состоит из целых чисел, каждое из которых по модулю не превышает 1000010000.

Выходные данные
Необходимо вывести разность индексов максимального и минимального элементов массива
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
	minI := 0
	maxV := sli[0]
	maxI := 0

	for i := 0; i < size; i++ {
		if sli[i] < minV {
			minV = sli[i]
			minI = i
		}
		if sli[i] > maxV {
			maxV = sli[i]
			maxI = i
		}
	}

	fmt.Println(maxI - minI)

}

// решено
