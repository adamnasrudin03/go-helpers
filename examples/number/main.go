package main

import (
	"fmt"

	help "github.com/adamnasrudin03/go-helpers"
)

func exampleRoundFloat() {
	number := 12.3456789
	fmt.Println(help.RoundUpFloat(number, 3))   // 12.346
	fmt.Println(help.RoundDownFloat(number, 3)) // 12.345

	number = -12.3456789
	fmt.Println(help.RoundUpFloat(number, 1))   // -12.3
	fmt.Println(help.RoundDownFloat(number, 1)) // -12.4

	number = 12.35
	fmt.Println(help.RoundFloat(number, 1, true))  // 12.4
	fmt.Println(help.RoundFloat(number, 1, false)) // 12.3
}

func exampleCheckValue() {
	fmt.Println(help.CheckArrayIntNil(nil))
	a := 123
	fmt.Println(help.CheckIntValue(nil))
	fmt.Println(help.CheckIntValue(&a))

	fmt.Println(help.CheckArrayFloat64Nil(nil))
	b := float64(1.23)
	fmt.Println(help.CheckFloat64Value(nil))
	fmt.Println(help.CheckFloat64Value(&b))
}

func main() {
	exampleRoundFloat()

	fmt.Println()
	arrInt := []int{}
	for i := 1; i <= 10; i++ {
		len := i
		if i >= 5 {
			len = 2
		}
		arrInt = append(arrInt, help.GenerateRandomNumber(len))
	}
	fmt.Println("arr", arrInt)
	min, max := help.GetMinMaxIntArray(arrInt)
	fmt.Println("min", min, "max", max)

	fmt.Println()
	exampleCheckValue()
}
