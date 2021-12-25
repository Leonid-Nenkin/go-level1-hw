package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("Введите длину праямоугольника: ")
	fmt.Scanln(&a)

	fmt.Print("Введите ширину праямоугольника: ")
	fmt.Scanln(&b)

	fmt.Print("Площадь праямоугольника: ", a*b)
}
