// Niska (string) u Go-u je read-only slice bitova. Go jezik i
// standardne biblioteke tretiraju niske na poseban način - kao
// kontejnere teksta enkodovanih u [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// U drugim jezicima, niske su skupovi "karaktera".
// U Go-u, kocept karaktera je drugačije poznat kao `rune`-a.
// Runa je integer koji se predstavlja pomoću Unicode koda.
// [Ovaj blog post](https://go.dev/blog/strings) je dobar uvod
// na temu string-ova i runa.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` je `string` kome je dodeljena vrednost
	// koja predstavlja reč "hello" na Tajlandskom.
	// Go niske su UTF-8 enkodovane.
	const s = "สวัสดี"

	// Uzimajući u obzir da su niske tipa `[]byte`, ova linija
	// će vratiti dužinu (length) niske tj. broj byte-ova.
	fmt.Println("Len:", len(s))

	// Indeksovanjem string-a, dobijamo sirove oblike karaktera
	// na svakom indeksu. Ova petlja generiše hex (heksadecimalne)
	// vrednosti byte-ova koji zajedno grade UTF-8 enkodovan `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Da bi pronašli broj _runa_ u nekom string-u, koristimo
	// `utf8` paket. Imajmo na umu da vreme izvršenja funkcije
	// `RuneCountInString` zavisi od veličine string-a,
	// jer ona mora da dekodira UTF-8 rune sekvencijalno.
	// Neki Tajlandski karakteri su reprezentovani UTF-8 kodovima
	// koji sadrže više byte-a, pa nas ova funkcija može
	// iznenaditi.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// `range` petlja iterira preko string-a na poseban način
	// i vraća svaku `runu` zajedno sa indeksom u string-u.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Istu iteraciju možemo postići i pomoću funkcije
	// `utf8.DecodeRuneInString` eksplicitno.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Ovde demonstriramo davanje `runa` vrednost funkciji.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Karakteri u jednostrukim navodnicima su _runa literali_.
	// Možemo da upoređujemo `runa` vrednost sa runa literalom direktno.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
