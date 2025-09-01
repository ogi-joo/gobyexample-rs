// Možemo da koristimo naše posebne tipove za `error`-e
// implementirajući `Error()` metodu. Evo jednog primera
// koji koristi gornju metodu za otkrivanje grešaka
// argumenta.

package main

import (
	"errors"
	"fmt"
)

// Praksa je da se posebni tipovi grešaka
// završavaju na "Error".
type argError struct {
	arg     int
	message string
}

// Kada dodamo `Error` metod, `argError` implementira
// interfejs `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Vraćamo naš tip greške.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.As` je naprednija verzija `errors.Is`.
	// Proverava da li data greška (ili bilo koja greška u lancu)
	// se podudara sa specifičnim error tipom i konvertuje je u
	// vrednost tog tipa, vraćajući `true`. Ako nema podudaranja,
	// vraća `false`.
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
