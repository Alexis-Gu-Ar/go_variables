package main

import (
	"fmt"
	"math"
)

func GetCircleArea(radius float64) float64 {
	return math.Pow(math.Pi*radius, 2)
}

func main() {
	var f1 float64

	fmt.Print("Dame el radio del circulo: ")
	fmt.Scan(&f1)
	var circleArea float64 = GetCircleArea(f1)

	fmt.Println("El area del circulo es ", circleArea)
}
