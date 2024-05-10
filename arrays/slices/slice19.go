/*
Количество элементов по условию

Дан список положительных целых чисел. Найдите все числа кратные 33 и при этом оканчивающиеся на 77, и замените каждое из этих чисел на их количество в массиве.

Входные данные
Сначала задано число NN — количество элементов в массиве (1≤N≤1001≤N≤100). Далее через пробел записаны NN чисел — элементы массива.
Массив состоит из целых положительных чисел, каждое из которых не превосходит 10001000.

Формат выходных данных
Выведите ответ на задачу.
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

	count := 0

	for i := 0; i < size; i++ {
		if (sli[i]%3 == 0) && (sli[i]%10 == 7) {
			count++
		}
	}

	for i := 0; i < size; i++ {
		if (sli[i]%3 == 0) && (sli[i]%10 == 7) {
			sli[i] = count
		}
	}

	for i := 0; i < size; i++ {
		fmt.Print(sli[i], " ")
	}

}

// решено
