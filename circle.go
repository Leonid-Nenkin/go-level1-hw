package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var pi float64 = math.Pi

	fmt.Print("Введите площадь окружности: ")
	fmt.Scanln(&a)

	var radius = math.Sqrt(a / pi)
	fmt.Print("Диаметр окружности: ", radius*2)
	fmt.Print("\n")
	fmt.Print("Длина окружности: ", 2*pi*radius)
}
