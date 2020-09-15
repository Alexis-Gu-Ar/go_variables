package main

import (
	"fmt"
)

func GetCelciusFromFahrenheit(temp float64) float64 {
	return (temp - 32) * 5 / 9
}

func main() {
	var f1 float64

	fmt.Print("Dame una temperatura en Fahrenheit: ")
	fmt.Scan(&f1)
	var celciusTemperature float64 = GetCelciusFromFahrenheit(f1)

	fmt.Println("La temperatura en Celcius es ", celciusTemperature, "C")
}
