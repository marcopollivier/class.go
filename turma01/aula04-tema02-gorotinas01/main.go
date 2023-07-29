package main

import (
	"fmt"
	"time"
)

func a() {
	fmt.Println("Inicio --- Funcao a")
	time.Sleep(time.Second * 2)
	fmt.Println("Final --- Funcao a")
}

func b() {
	fmt.Println("Inicio --- Funcao b")
	time.Sleep(time.Second * 1)
	fmt.Println("Final --- Funcao b")
}

func c() {
	fmt.Println("Inicio --- Funcao c")
	time.Sleep(time.Second * 5)
	fmt.Println("Final --- Funcao c")
}

func main() {
	go a()
	go b()
	go c()

	time.Sleep(time.Minute * 3)
}
