// _Funkcije_ su centar Go jezika. Učićemo o
// funkcijama uz pomoć nekoliko dobrih primera.

package main

import "fmt"

// Ovo je funkcija koja uzima dva `int`-a i vraća
// njihov zbir tipa `int`.
func plus(a int, b int) int {

	// U jeziku Go, moramo eksplicitno da navedemo
	// tip koji vraćamo (return type). Dakle nije
	// automatski izračunat na osnovu nekog izraza.
	return a + b
}

// Ako imamo dva ili više argumenata istog tipa, možemo
// navesti na kraju niza kao:
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Zovemo funkciju kao i u drugim jezicima
	// `name(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
