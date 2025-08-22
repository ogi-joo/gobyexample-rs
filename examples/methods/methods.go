// Go podržava _metode_ definisane na strukturi.

package main

import "fmt"

type rect struct {
	width, height int
}

// Ovaj metod `area` ima _receiver tip_ `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Metode mogu da se definišu i za pokazivače i za
// vrednosti receiver tipova. Evo jedan primer za vrednost
// resiver-a.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Ovde zovemo 2 metode definisane na našoj strukturi.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatski vrši konverziju između pokazivača i vrednosti
	// za metode struktura. Ponekad možemo da koristimo
	// pokazivač receiver-a da zaobiđemo kopiranje na pozivu metode
	// ili dozvolimo metodi da menja strukturu koju dobija.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
