package main

import (
	"fmt"
)

func main() {
	var a int64

	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&a)

	fmt.Print("Количество сотен: ", a/100)
	fmt.Print("Количество десятков: ", a/10)
}
