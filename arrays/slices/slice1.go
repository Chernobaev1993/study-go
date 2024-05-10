/*
Дан массив, состоящий из целых чисел. Напишите программу, которая подсчитает количество элементов массива, больших предыдущего (левого соседа).

Входные данные
Сначала задано число NN — количество элементов в массиве (1≤N≤1001≤N≤100). Далее через пробел записаны NN чисел — элементы массива.
Массив состоит из целых чисел, каждое из которых по модулю не превышает 109109.

Выходные данные
Необходимо вывести единственное число — количество элементов массива, больших предыдущего элемента.
*/

package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	count := 0

	slice1 := make([]int, n)
	for i := 0; i < len(slice1); i++ {
		_, _ = fmt.Scan(&slice1[i])
	}
	for i := 1; i < len(slice1); i++ {
		if slice1[i-1] < slice1[i] {
			count++
		}
	}
	fmt.Println(count)
}

// решено
// отлично
