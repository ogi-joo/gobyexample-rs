// U [prethodnom](range-over-built-in-types) range primeru,
// videli smo kako `for` i `range` mogu da iteriraju preko klasičnih
// data struktura. Ove sintakse možemo takođe koristiti i za
// iteriranje preko vrednosti kanala!

package main

import "fmt"

func main() {

	// Iteriraćemo preko 2 vrednosti u `queue` kanalu.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Ovaj `range` iterira preko svakog elementa koji
	// se prima iz `queue`. Zbog toga što smo zatvorili
	// kanal gore, iteracija staje nakon što primi
	// 2 elementa iz kanala.
	for elem := range queue {
		fmt.Println(elem)
	}
}
