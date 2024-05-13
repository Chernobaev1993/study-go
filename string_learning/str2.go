/*
Перевертыш
Дана строка. Найдите перевернутую ей строку.

Входные данные
На вход подается строка, длина которой не превосходит 10001000.
Выходные данные

Необходимо вывести ее перевертыш.
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

	for _, value := range s {
		defer fmt.Print(string(value))
	}

	
}