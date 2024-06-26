/*
Тимур и забор

Тимур с друзьями идёт по запрещенной территории возле забора высотой hh. Чтобы охранник их не заметил, высота каждого из друзей не должна превышать hh.
При этом каждый из них может пригнуться, тогда его точно не будет видно охраннику. Высота ii-го из друзей равняется aiai​.
Будем считать, что ширина человека, идущего в полный рост, равна 11, а согнутого — 22.
Друзья хотят идти по дороге в один ряд, чтобы иметь возможность разговаривать друг с другом.
Какой минимальной ширины должна быть дорога, чтобы все они поместились на ней в один ряд и охранник никого не увидел?

Входные данные
В первой строке входных данных записаны два целых числа nn и hh (1 ≤ n ≤ 1000, 1 ≤ h ≤ 1000)(1 ≤ n ≤ 1000, 1 ≤ h ≤ 1000) — количество друзей и высота забора соответственно.
Во второй строке записаны nn целых чисел ai (1 ≤ ai ≤ 2h)ai​ (1 ≤ ai​ ≤ 2h), ii-е из которых равняется высоте ii-го друга.

Выходные данные
Выведите единственное целое число — минимальную подходящую ширину дороги.
*/
package main

import (
	"fmt"
)

func main() {
	var n, h int
	fmt.Scan(&n, &h)

	sli := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sli[i])
	}

	count := 0
	for _, value := range sli {
		if value > h {
			count += 2
		} else {
			count++
		}
	}

	fmt.Println(count)
}

// решено
