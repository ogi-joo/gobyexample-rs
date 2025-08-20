// Go podržava [_anonimne funkcije_](https://sr.wikipedia.org/wiki/Anonimna_funkcija),
// koje mogu da formiraju <a href="https://sr.wikipedia.org/wiki/%D0%97%D0%B0%D1%82%D0%B2%D0%BE%D1%80%D0%B5%D1%9A%D0%B5_(%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%B8%D1%80%D0%B0%D1%9A%D0%B5)"><em>closure</em></a>.
// Anonimne funkcije su korisne kada želimo da definišemo
// unutrašnju funkciju a da ne moramo da joj dodelimo naziv.

package main

import "fmt"

// Ova funkcija `intSeq` vraća drugu funkciju, koju
// smo mi anonimno deklarisali u telu funkcije `intSeq`.
// Vraćena funkcija _zatvara_ varijablu `i` i tako
// formira closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Zovemo `intSeq`, vezujući vrednost (funkciju)
	// za `nextInt`. Ova funkcijska vrednost vezuje
	// sopstvenu vrednost `i`, koja će biti ažurirana
	// kada zovemo `nextInt`.
	nextInt := intSeq()

	// Možemo videti mogućnosti closure-a zvanjem `nextInt`
	// nekoliko puta.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Da proverimo da je kontekst jedinstven za tu
	// partikularnu funkciju, možemo da napravimo novu varijablu.
	newInts := intSeq()
	fmt.Println(newInts())
}
