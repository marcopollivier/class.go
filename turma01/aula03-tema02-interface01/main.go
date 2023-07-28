package main

import (
	"fmt"
	"math"
)

// Retangulo
type Retangulo struct {
	largura, altura float64
}

func (r Retangulo) perimetro() float64 {
	return 2*r.altura + 2*r.largura
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

// circulo
type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c Circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// interface
type ObjetosGeometricos interface {
	area() float64
	perimetro() float64
}

func perimetro(obj ObjetosGeometricos) float64 {
	return obj.perimetro()
}

func main() {

	var r Retangulo = Retangulo{largura: 10, altura: 2}
	var c Circulo = Circulo{raio: 3}

	fmt.Println(perimetro(r))
	fmt.Println(perimetro(c))

}
