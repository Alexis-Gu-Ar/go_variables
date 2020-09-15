package main

import (
	"fmt"
	"math"
)

func GetSqrArea(side float64) float64 {
	return math.Pow(side, 2)
}

func main() {
	var f1 float64

	fmt.Print("Dame un lado del cuadrado: ")
	fmt.Scan(&f1)
	var sqrArea float64 = GetSqrArea(f1)

	fmt.Println("El area del cuadrado es ", sqrArea)
}
