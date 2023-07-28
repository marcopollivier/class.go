package main

import "testing"

// go test -bench=.
//go test -bench=AddNomes -cpu=8 -benchmem -benchtime=5s -count 5
func BenchmarkPerimetro(b *testing.B) {
	var r1 retangulo = retangulo{largura: 10, altura: 2}

	for n := 0; n > b.N; n++ {
		Perimetro(r1)
	}
}
