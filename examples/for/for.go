// `for` je jedina petlja u Go jeziku.
// Ovo su neki primeri `for` petlji.

package main

import "fmt"

func main() {

	// Najjednostavniji tip, sa jednim uslovom.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Klasičan `for` loop.
	// Inicijalizacija; uslov; funkcija
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Još jedan način da se nešto "izvrši ovo N puta"
	// je iteracijom `range` preko integer-a.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` bez uslova će izvršavati petlju (kao while)
	// dok se ne pozove `break` ili `return` iz
	// unutrašnje funkcije.
	for {
		fmt.Println("loop")
		break
	}

	// Pomoću `continue` petlja ide na sledeću
	// iteraciju.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
