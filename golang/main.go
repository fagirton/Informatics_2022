package main

import (
	"fmt"
	"math"
)

func main() {
	const a float64 = 4.1
	const b float64 = 2.7

	fmt.Println(taskA(a, b))
	fmt.Println(taskB(a, b))
	test()
}

func formula(x float64, a float64, b float64) float64 {
	var y float64 = ((a*a*a)*(math.Cbrt(x)) - b*(math.Log10(x)/math.Log10(5))) / (math.Pow(math.Log10(x-1), 3))
	return y
}

func taskA(a float64, b float64) []float64 {
	fmt.Println("Задание А")
	var i float64 = 1.5
	var res []float64
	for ; i < 3.5; i = i + 0.4 {
		res = append(res, formula(i, a, b))
	}
	return res
}
func taskB(a float64, b float64) []float64 {
	fmt.Println("Задание B")
	var x []float64 = []float64{1.9, 2.15, 2.34, 2.74, 3.16}
	var res []float64
	var i int8
	for ; i < 5; i++ {
		res = append(res, formula(x[i], a, b))
	}
	return res
}
