package main

import "fmt"

func apendAssync(mensagem string, c chan string) {
	c <- mensagem + " passou pelo append"
}

func main() {

	c := make(chan string, 2)

	go apendAssync("mensagem 1", c)
	go apendAssync("mensagem 2", c)
	go apendAssync("mensagem 3", c)
	go apendAssync("mensagem 4", c)
	go apendAssync("mensagem 5", c)
	go apendAssync("mensagem 6", c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
