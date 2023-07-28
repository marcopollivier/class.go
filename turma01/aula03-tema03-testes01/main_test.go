package main

import "testing"

func TestPerimetro(t *testing.T) {
	var r retangulo = retangulo{largura: 10, altura: 2}
	var c circulo = circulo{raio: 200}

	a := Perimetro(r)
	b := Perimetro(c)

	if a != 24 {
		t.Errorf("Erro ... %f; want 24", a)
	}

	if b != 1256.6370614359173 {
		t.Errorf("Erro %f; want 1256...", b)
	}
}
