// Go podržava <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pokazivače (pointers)</a></em>,
// dopuštajući nam onda da koristimo zanimljive stvari kao reference.

package main

import "fmt"

// Pokazaćemo kako pokazivači rade u odnosu na vrednosti
// pomoću 2 funkcije: `zeroval` i `zeroptr`. `zeroval` ima
// `int` parametar, dakle argumenti će joj se dodeljivati kao
// vrednost. `zeroval` će dobiti kopiju `ival` koja nije
// ista varijabla kao i u liniji gde se poziva ova funkcija.
func zeroval(ival int) {
	ival = 0
}

// Na drugu ruku, `zeroptr` ima `*int` parametar, što znači
// da ona uzima `int` pokazivač. `*iptr` označava dereferenciranje.
// Ovo u suštini znači: uzmi adresu `iptr` varijable i na toj adresi
// menjaj vrednost. Ova funkcija ne menja kopiju već originalnu
// vrednost varijable iz main() funkcije gde je zeroptr() pozvana.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// Sintaksa `&i` daje adresu varijable `i`,
	// tj. pokazivač na `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pokazivači se takođe mogu ispisivati.
	fmt.Println("pointer:", &i)
}
