/*
Шеренга

Петя впервые пришел на урок физкультуры в новой школе. Перед началом урока ученики выстраиваются по росту, в порядке невозрастания.
Напишите программу, которая определит на какое место в шеренге Пете нужно встать, чтобы не нарушить традицию,
если заранее известен рост каждого ученика и эти данные уже расположены по невозрастанию (то есть каждое следующее число не больше предыдущего).
Если в классе есть несколько учеников с таким же ростом, как у Пети, то программа должна расположить его после них.

Входные данные
Сначала задано число NN — количество учеников (не считая Петю)(1≤N≤1001≤N≤100).
Далее через пробел записаны NN чисел — элементы массива. Массив состоит из натуральных чисел, не превосходящих 200200 (рост учеников в сантиметрах).
Затем, на новой строке, вводится рост самого Пети.

Выходные данные
Необходимо вывести единственное число - номер Пети в шеренге учеников.
*/
package main

import "fmt"

func main() {
	var size, h int
	fmt.Scan(&size)

	sli := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&sli[i])
	}

	fmt.Scan(&h)

	q := 1

	for _, value := range sli {
		if h <= value {
			q++
		}
	}
	fmt.Println(q)

}

// решено
