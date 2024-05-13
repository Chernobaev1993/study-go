/*
Палиндром?

По данной строке определите, является ли она палиндромом? То есть, которое одинаково читается слева направо и справа налево. Например, слово "шалаш".
Входные данные

На вход подается одна строка без пробелов.
Выходные данные

Необходимо вывести  "YES", если строка является палиндромом, и "NO" - в противном случае.
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

	if s == reverseString(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func reverseString(x string) string {
	temp := ""
	for i := len(x) - 1; i >= 0; i-- {
		temp += string(x[i])
	}
	return temp
}