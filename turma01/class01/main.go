package main

import (
	"fmt"
)

func imprimirOTexto(resultado string) {
	fmt.Println(resultado)
}

func main() {
	defer imprimirOTexto("Ollivier")

	imprimirOTexto("Marco")
	imprimirOTexto("Paulo")

}
