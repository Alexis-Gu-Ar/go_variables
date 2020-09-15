package main

import (
	"fmt"
	"math"
)

func GetSqrArea(side float64) float64 {
	return math.Pow(side, 2)
}

func GetTriangleArea(width float64, height float64) float64 {
	return width * height / 2
}

func GetCircleArea(radius float64) float64 {
	return math.Pow(math.Pi*radius, 2)
}

func GetCelciusFromFahrenheit(temp float64) float64 {
	return (temp - 32) * 5 / 9
}

func main() {
	var (
		f1 float64
		f2 float64
	)

	fmt.Print("Dame la base y la altura del triangulo (separadas por espacio): ")
	fmt.Scanf("%f %f", &f1, &f2)
	fmt.Println(f1, f2)
	var traingleArea float64 = GetTriangleArea(f1, f2)

	fmt.Print("Dame un lado del cuadrado: ")
	fmt.Scan(&f1)
	var sqrArea float64 = GetSqrArea(f1)

	fmt.Print("Dame el radio del circulo: ")
	fmt.Scan(&f1)
	var circleArea float64 = GetCircleArea(f1)

	fmt.Print("Dame una temperatura en Fahrenheit: ")
	fmt.Scan(&f1)
	var celciusTemperature float64 = GetCelciusFromFahrenheit(f1)

	fmt.Println("El area del cuadrado es ", sqrArea)
	fmt.Println("El area del trinagulo es ", traingleArea)
	fmt.Println("El area del circulo es ", circleArea)
	fmt.Println("La temperatura en Celcius es ", celciusTemperature, "C")
}
