// [_Varijadna funkcija_](https://en.wikipedia.org/wiki/Variadic_function)
// ili funkcija sa promenljivim brojem argumenata.
// Na primer, funkcija `fmt.Println` je klasična varijadna funkcija.

package main

import "fmt"

// Ovo je funkcija koja će primiti promenljiv broj atributa
// tipa `int`.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// Unutar funkcije, tip varijable `nums` je jednak
	// tipu `[]int`. Možemo da zovemo `len(nums)`,
	// iteriramo pomoću `range`, itd.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Varijadne funkcije se klasično zovu
	// sa zarezima izmedju vrednosti.
	sum(1, 2)
	sum(1, 2, 3)

	// Ako već imamo više vrednosti u slice-u,
	// možemo da ih pozovemu u funkciji kao ovde.
	// Primetimo `nums...`
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
