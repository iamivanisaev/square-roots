package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Println("Реши уравнение")
	fmt.Println("Введи a:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Ошибка!")
	}

	fmt.Println("Введи b:")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Ошибка!")
	}

	fmt.Println("Введи c:")
	_, err = fmt.Scan(&c)
	if err != nil {
		fmt.Println("Ошибка!")
	}

	D := (b * b) - 4*(a*c)
	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнение имеет 2 корня\nD = " + fmt.Sprint(D))
		fmt.Println("x1: " + fmt.Sprint(x1) + "\nx2: " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println("Ваше уравнение имеет 1 корень")
		fmt.Println("x: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Ваше уравнение не имеет корней \nD <0\nD = " + fmt.Sprint(D))
	}
}
