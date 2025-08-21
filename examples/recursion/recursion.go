// Go podržava
// <a href="https://sr.wikipedia.org/wiki/%D0%A0%D0%B5%D0%BA%D1%83%D1%80%D0%B7%D0%B8%D1%98%D0%B0_(%D0%BA%D0%BE%D0%BC%D0%BF%D1%98%D1%83%D1%82%D0%B5%D1%80%D1%81%D0%BA%D0%B5_%D0%BD%D0%B0%D1%83%D0%BA%D0%B5)"><em>rekurzivne funkcije</em></a>.
// Evo jednog klasičnog primera.

package main

import "fmt"

// Klasična funkcija za nalaženje faktorijela.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Anonimne funkcije mogu takođe biti rekurzivne, ali ovo zahteva
	// eksplicitnu deklaraciju varijable pomoću `var` sintakse da bi
	// funkciju zapisali pre nego je definišemo.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Zato što je `fib` već bio definisan u `main`-u, Go
		// zna koji `fib` da pozove.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
