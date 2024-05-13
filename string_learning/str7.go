/*
Идентификатор
Определить, является ли введенное слово идентификатором, т.е. начинается ли оно с английской буквы в любом регистре или со знака подчеркивания,
так же она НЕ должна содержать других символов, КРОМЕ букв английского алфавита (в любом регистре), цифр и знака подчеркивания.

Входные данные
На вход программе подается одна строка.

Выходные данные
Если строка является идентификатором, необходимо вывести "YES", иначе - "NO".
*/
package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	flag := false
	for index, value := range s {
		if index == 0 {
			if value == 95 || (value >= 65 && value <= 90) || (value >= 97 && value <= 122){
				flag = true
			} else {
				flag = false
				break
			}
			continue
		}
		if (value >= 65 && value <= 90) || (value >= 97 && value <= 122) || (value >= 48 && value <= 57) || value == 95 {
			flag = true
		} else {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
