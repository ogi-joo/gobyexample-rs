// _Switch naredbe_ predstavljaju uslovne izraze
// u više grana.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Ovo je klasičan `switch`.
	i := 2
	fmt.Print(i, " se piše kao ")
	switch i {
	case 1:
		fmt.Println("jedan")
	case 2:
		fmt.Println("dva")
	case 3:
		fmt.Println("tri")
	}

	// Možemo da koristimo zarez da bi smo
	// gađali uslove sa istim telom funcije.
	// U ovom primeru koristimo i _default_ naredbu.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Danas je vikend!")
	default:
		fmt.Println("Danas je radni dan.")
	}

	// `switch` bez uslova je drugi način kako možemo
	// da predstavimo if-else grananje.
	// Ovde prikazujemo kako `case` može da povlači i
	// vrednosti koje nisu konstante.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Prepodne je")
	default:
		fmt.Println("Poslepodne je")
	}

	// Tipovni `switch` upoređuje tipove umesto vrednosti,
	// ovo možemo da koristimo da otkrijemo tip vrednosti.
	// U ovom primeru, varijabla `t` ima svoj  tip kojeg
	// ispitujemo.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Ja sam bool")
		case int:
			fmt.Println("Ja sam int")
		default:
			fmt.Printf("Ne znam šta je %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
