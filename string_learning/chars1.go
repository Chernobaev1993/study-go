/*
	Изменить регистр символа

Измените регистр символа, если он был латинской буквой: сделайте его заглавным, если он был строчной буквой и наоборот.
Входные данные

Задан единственный символ латинского алфавита c.c.
Выходные данные

Необходимо вывести получившийся символ.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	_ = scanner.Scan()
	stroka := scanner.Text()

	fmt.Println(stroka)

	count := 0
	for i := 0; i < n-2; i++ {

		if string(stroka[i]) == "x" && string(stroka[i+1]) == "x" && string(stroka[i+2]) == "x" {
			count++
		}
	}
	fmt.Println(count)
}
