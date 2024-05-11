/*
Сумма signsign

Найдите значение z=sign(a)+sign(b)z=sign(a)+sign(b), где
sign(x)={ −1,x<00,x=01,x>0sign(x)=⎩
⎨
⎧​ ​−1,x<00,x=01,x>0​

Входные данные
Вводится строка, в которой через пробел записаны два целых числа aa и bb.

Выходные данные
Неободимо вывести значение zz.

Примечание: данную задачу предполагается решить с помощью функций.
*/
package main

import "fmt"

func main() {
	var a, b, z int
	fmt.Scan(&a, &b)
	
	z = signFind(a) + signFind(b)
	fmt.Println(z)
}

func signFind(x int) int {
    switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	}
    return 0
}
// Решено