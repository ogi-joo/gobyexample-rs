// _Slices_ (slajs, isečak) su važni tipovi vrednosti u Go-u, dajući nam
// bolji interfejs za sekvence u odnosu na nizove.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Za razliku od nizova, sintaksni tip slice-a je samo
	// tip elemenata koje on sadrži (ne i broj elemenata).
	// Ne inicijalizovan slice vrednosti nil i dužne 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Da bi se napravio slice sa dužinom većom od 0,
	// koristimo ugrađen `make`. Ovde pravimo slice sa
	// elementima tipa `string`, dužine `3`.
	// Po default-u novi slice ima kapacitet jednak
	// dužini; ako znamo da će slice da raste u budućnosti,
	// možemo da mu dodelimo kapacitet eksplicitno
	// kao dodatan parametar u `make` funkciji.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Možemo da "set & get" isto kao sa nizovima.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` vraća dužinu - kao što očekujemo.
	fmt.Println("len:", len(s))

	// Osim ovih jednostavnih operacija, slice pruža
	// dodatne funkcionalnosti koje su napredak u odnosu
	// na nizove. Jedna od njih je ugrađena `append` funkcija
	// koja vraća slice koji sada sadrži jedan ili više dodatnih
	// vrednosti. Primetimo da moramo sačuvati return
	// `append`-a jer dobijamo novu slice vrednost.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Možemo da kopiramo, `copy`, slice. Ovde smo
	// napravili prazan `c` slice iste dužine ako `s`
	// i kopirali smo `s` u `c`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Koristimo "slice" operator sa sintaksom
	// `slice[low:high]` za dobijanje isečka.
	// Na primer, ovde dobijamo isečak
	// elemenata `s[2]`, `s[3]`, i `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Ovako dobijamo do `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// Ovako sečemo od `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Kao i kod nizova, možemo da inicijalizujemo
	// i deklarišemo slice u jednoj liniji
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Go paket `slices` sadrži brojne korisne
	// funkcije za upravljanje isečcima.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slice može da bude višedimenzionalan.
	// Dužina unutrašnjeg slice-a je promenljiva,
	// za razliku od višedimenzionalnih nizova.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1

		twoD[i] = make([]int, innerLen)

		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
