// U Go-u, nizovi, _arrays_, su sekvence elemenata kojima se
// dužina zna unapred. U tipičnom Go kodu, [slice, isečak](slices) su
// dosta više zastupljeniji nego nizovi; nizovi su korisni samo u nekim
// specifičnim slučajevima.

package main

import "fmt"

func main() {

	// Ovde kreiramo niz `a` koji će sadržati tačno
	// 5 `int`-a. Tip niza i njegova dužina spadaju
	// zajedno u njegov sintaksni tip. Po default-u niz
	// je `zero-valued`, nulte vrednosti za `int` su `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Možemo da postavimo vrednost preko indeksa pomoću
	// `array[index] = value` sintakse, i tako isto dobijemo
	// vrednost sa `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Ugrađena sintaksa `len` vraća dužinu nekog niza.
	fmt.Println("len:", len(a))

	// Ovakvu sintaksu koristimo da deklarišemo i inicijalizujemo
	// neki niz u jednoj liniji.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// A možemo da dopustimo i kompajleru da izbroji
	// dužinu niza za nas pomoću `...`
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Ako koristimo `:` u inicijalizaciji,
	// elementi između će biti nulirani.
	// (Za niz b smo stavili da je dužina 5)
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Tip niza je jednodimenzionalan, ali možemo da
	// pravimo višedimenzialne nizove.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Višedimenzialne nizove takođe možemo da
	// deklarišemo i inicijalizujemo u isto vreme
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
