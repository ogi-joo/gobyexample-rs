// _Interfejsi_ su imenovane kolekcije
// potpisa funkcija.

package main

import (
	"fmt"
	"math"
)

// Ovo je jednostavan interfejs za geometrijske oblike.
type geometry interface {
	area() float64
	perim() float64
}

// U našem primeru, implementiraćemo ovaj interfejs na
// `rect` i `circle` strukturama.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Da bi implementirali interfejs u Go-u, potrebno je
// samo da implementiramo sve metode interfejsa. Ovde
// implementiramo interfejs `geometry` na `rect`.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// I implementacija za `circle`...
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Ako je varijabla tipa nekog interfejsa, onda možemo da
// zovemo metode tog interfejsa na njoj. Evo jedne
// generične funkcije `measure`, koja radi nad bilo kojom
// `geometry` varijablom.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Ponekad je korisno da znamo tip varijable.
// Jedna opcija je pomoću *provere tipa*
// kao u ovom primeru;
// drugi primer je [tip preko `switch`-a](switch).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Obe `circle` i `rect` strukture imaju
	// implementiran interfejs `geometry`,
	// pa možemo da koristimo instance
	// ovih struktura kao argumente za `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
