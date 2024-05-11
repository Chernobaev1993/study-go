/*
Число сочетаний

По данным натуральным nn и kk вычислите значение Сnk=n!k!(n−k)!Сnk​=k!(n−k)!n!​(число сочетаний из nn элементов по kk). 

Входные данные
На первой строке вводится натуральное число nn, не превосходящее 1010. 
На второй строке вводится целое неотрицательное число kk, не превосходящее 1010.

Выходные данные
Необходимо вывести значение СnkСnk​.

Примечание: данную задачу предполагается решить с помощью функций. Не забывайте, что 0!=1.0!=1.
*/
package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	factN := factFind(n)
	factK := factFind(k)
	factSub := factFind(n - k)
	x := factN / (factK * factSub)
	fmt.Println(x)
}

func factFind(a int) int {
    fact := 1
	if a == 0 {
		return 1
	} else {
        for i := 1; i <= a; i++  {
			fact *= i
		}
		return fact
	}
}

