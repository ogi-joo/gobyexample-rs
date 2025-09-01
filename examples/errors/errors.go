// Idimatski je da u Go-u komuniciramo greške (errors)
// eksplicitno, preko odvojene return vrednosti. Slično je
// `exceptions`-ima korišćene u jezicima kao Java i Ruby.
// Go-ov pristup nam omogućava da jednostavno vidimo koje
// funkcije vraćaju greške i da ih manipulišemo istom
// jezičkom sintaksom kao što i za druge,
// bezgrešne tasko-ove.
//
// Pogledaj zvaničan [errors paket](https://pkg.go.dev/errors)
// i [ovaj blog post](https://go.dev/blog/go1.13-errors) za više
// informacija.

package main

import (
	"errors"
	"fmt"
)

// Praksa je da su greške poslednja vraćena vrednost
// i da imaju tip `error`, ugrađen interfejs.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` konstruiše `error` vrednost
		// sa datom greškom u tekstualnom obliku.
		return -1, errors.New("can't work with 42")
	}

	// Vrednost `nil` na poziciji greške nagoveštava
	// da nije bilo greške.
	return arg + 3, nil
}

// Stražarska greška (sentinel error) je deklarisana greška
// i označava specifičnu grešku.
var ErrOutOfTea = fmt.Errorf("nema više čaja")
var ErrPower = fmt.Errorf("voda se ne može skuvati")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// Možemo da objedinimo stražarske i obične greške.
		// Najjednostavniji način je koristeći
		// `%w` sintaksu u `fmt.Errorf`. Objedinjene greške
		// kreiraju lanac (A zaokružuje B, koji zaokružuje C, itd.)
		// što je korisno za kasniju implementaciju `errors.Is`
		// i `errors.As` funkcija.
		return fmt.Errorf("kuvanje čaja: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// Praksa je da se greške proveravaju u inline if
		// loop-u.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` proverava da li data greška (ili bilo koja u lancu)
			// se podudara sa nekom drugom. Ovo je dosta korisno jer nam
			// pomaže da raspoznamo bilo koju grešku, bila ona objedinjena ili ne.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("Trebamo kupiti novi čaj!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Sada je noć.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Čaj je spreman!")
	}
}
