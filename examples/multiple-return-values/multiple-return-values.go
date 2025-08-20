// Go ima ugrađenu podršku za _višestruko vraćanje vrednosti_.
// Ova karakteristika nam dopušta da pišemo idiomatski kod,
// npr. da vratimo i vrednosti i grešku iz neke funkcije.

package main

import "fmt"

// Potpis `(int, int)` u ovoj funkciji nam govori
// da ova funkcija vraća dva `int`-a.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Višestruko dodeljivanje vrednosti.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Ako nam ne trebaju sve vraćene vrednosti,
	// koristimo `_` za ignorisanje.
	_, c := vals()
	fmt.Println(c)
}
