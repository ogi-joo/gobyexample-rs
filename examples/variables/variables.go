// U Go-u, _varijable_ su eksplicitno deklarisane i korišćene
// od strane kompajlera da bi npr. pozivi funkcija bili _type-safe_.

package main

import "fmt"

func main() {

	// Pomoću `var` deklarišemo 1 ili više varijabli.
	var a = "initial"
	fmt.Println(a)

	// Možemo da deklarišemo više varijabli od jednom.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Za varijable definisane pomoću `var`, Go će automatski
	// izvesti, _infer_, tip varijable u odnosu na tip inicijalne vrednosti.
	var d = true
	fmt.Println(d)

	// Varijable deklarisane bez početne vrednosti
	// su nulte varijable, _zero-valued_. Na primer,
	// nulta vrednost za `int` je `0`.
	var e int
	fmt.Println(e)

	// Sintaksa `:=` je skraćenica za istovremeno deklarisanje i
	// inicijalizovanje varijable, npr. za
	// `var f string = "apple"`.
	// Ova sintaksa je validna samo unutar funkcija.
	f := "apple"
	fmt.Println(f)
}
