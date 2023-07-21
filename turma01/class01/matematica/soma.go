package matematica

import "fmt"

func SomaDireta(num1 float32, num2 float32) float32 {
	return num1 + num2
}

func SomaDiretaComTres(num1, num2 float32) float32 {
	return num1 + num2
}

func SomaESubtrai(num1, num2 float32) (soma float32, subtrai float32) {
	soma = num1 + num2
	subtrai = num1 - num2

	return soma, subtrai
}

func Soma(nums ...float32) (somatorio float32) {
	//inicia o resultado com 0
	var total float32 = 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

	return somatorio
}
