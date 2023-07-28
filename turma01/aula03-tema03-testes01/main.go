package main

import "math"

// retangulo
type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (r retangulo) perimetro() float64 {
	return 2*r.altura + 2*r.largura
}

// circulo
type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// um pouco mais generico

type ObjetoGeometrico interface {
	area() float64
	perimetro() float64
}

func Perimetro(obj ObjetoGeometrico) float64 {
	return obj.perimetro()
}

func main() {

	var r retangulo = retangulo{largura: 10, altura: 2}
	var c circulo = circulo{raio: 200}

	Perimetro(r)
	Perimetro(c)

}
