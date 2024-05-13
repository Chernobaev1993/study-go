/*
Две одинаковые буквы
Дана строка. Известно, что она содержит ровно две одинаковые буквы. Найдите эти буквы. Гарантируется, что повторяются буквы только одного вида.

Входные данные
На вход подается одна строка. Строка может состоять как из строчных, так и заглавных букв.

Выходные данные
Необходимо вывести букву, которая встречается в строке дважды.
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

	fmt.Println(findDoubleLit(s))

}

func findDoubleLit(x string) string {
	for i := 0; i < len(x); i++ {
		for j := 1; j < len(x); j++ {
			if i == j {
				continue
			} else {
				if string(x[i]) == string(x[j]) {
					return string(x[i])
				}
			}

		}
	}
	return "o"
}