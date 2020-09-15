package main

import (
	"fmt"
)

func GetTriangleArea(width float64, height float64) float64 {
	return width * height / 2
}

func main() {
	var (
		f1 float64
		f2 float64
	)

	fmt.Print("Dame la base y la altura del triangulo (separadas por espacio): ")
	fmt.Scanf("%f %f", &f1, &f2)
	var traingleArea float64 = GetTriangleArea(f1, f2)

	fmt.Println("El area del trinagulo es ", traingleArea)
}
