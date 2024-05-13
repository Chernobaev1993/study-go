/*
k-ый символ
Дана строка. Удалите kk-ый символ в ней.

Входные данн
На первой строке вводится строка ss. На второй строке вводится целое число 1≤k≤∣s∣1≤k≤∣s∣, где ∣s∣∣s∣ длина строки ss.

Выходные данные
Необходимо вывести строку ss без kk-го символа.
*/
package main

import (
	"fmt"
	//"bufio"
	//"os"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// _ = scanner.Scan()
	// s := scanner.Text()

	var s string
	fmt.Scan(&s)

	flag := false
	for _, value := range s {
		
		if string(value) == "x" {
			fmt.Println("x")
			flag = true
			break
		} else if string(value) == "w" {
			fmt.Println("w")
			flag = true
			break
		}
	}

	if flag == false {
		fmt.Println(-1)
	}
}
