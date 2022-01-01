package main

import (
	"fmt"
	"math/big"
	"os"
)

var cache = make(map[int]*big.Int, 0)

func Fib(n int) *big.Int {
	if cache[n] != nil {
		return cache[n]
	}

	switch n {
	case 0:
		return big.NewInt(0)
	case 1:
		return big.NewInt(1)
	default:
		a := Fib(n - 1)
		b := Fib(n - 2)

		cache[n-1] = a
		cache[n-2] = b

		value := big.NewInt(0)
		value.Add(value, a)
		value.Add(value, b)

		return value
	}
}

func main() {
	var a int
	fmt.Print("Введите номер числа Фибоначчи (номера начинаются с 1): ")
	fmt.Scanln(&a)

	if a < 1 {
		fmt.Println("Номера начинаются с 1")
		os.Exit(1)
	}

	fmt.Print("Число фибоначи: ", Fib(a-1))
}
