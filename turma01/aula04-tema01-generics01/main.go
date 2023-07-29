package main

import "fmt"

func soma[X int | float32 | string](val1 X, val2 X) X {
	return val1 + val2
}

func main() {
	a := soma[int](1, 2)
	b := soma[float32](1.1, 2.1)
	c := soma[string]("Oi", " mundo")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
