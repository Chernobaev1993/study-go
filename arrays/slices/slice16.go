/*
Состязания

В метании молота состязается nn спортcменов. Каждый из них сделал mm бросков. Победителем считается тот спортсмен, у которого сумма результатов по всем броскам максимальна.
Если перенумеровать спортсменов числами от 0 до n−1n−1, а попытки каждого из них – от 0 до m−1m−1, то на вход программа получает массив A[n,m]A[n,m],
состоящий из неотрицательных целых чисел. Программа должна определить максимальную сумму чисел в одной строке и вывести на экран эту сумму и номер строки,
для которой достигается эта сумма.

Входные данные
Программа получает на вход два числа nn и mm, являющиеся числом строк и столбцов в массиве. Далее во входном потоке идет nn строк по mm чисел, являющихся элементами массива.

Выходные данные
Программа должна вывести 22 числа: сумму и номер строки, для которой эта сумма достигается. Если таких строк несколько, то выводится номер наименьшей из них.
Не забудьте, что нумерация строк (спортсменов) начинается с 00.
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

	for index, value := range sli {
		if index == 0 {
			value = 5
		}
		fmt.Println(index, " - ", value)
	}

	for index, value := range sli {

		fmt.Println(index, " - ", value)
	}

}

// решено