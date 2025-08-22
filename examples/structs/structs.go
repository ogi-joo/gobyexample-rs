// Go _strukture_ su tipovane kolekcije polja.
// Korisne su za grupisanje podataka zajedno
// da bi napravili rekorde.

package main

import "fmt"

// Ova `person` struktura ima `name` i `age` polja.
type person struct {
	name string
	age  int
}

// `newPerson` pravi novu person strukturu sa datim imenom.
func newPerson(name string) *person {
	// Go ima ugrađeno sakupljanje smeća (garbage collecting).
	// Možemo slobodno vratiti pokazivač na lokalnu varijablu,
	// biće obrisan čim ne postoji ni jedna referenca koja
	// ukazuje na njega.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Ova sintaksa pravi novu strukturu.
	fmt.Println(person{"Bob", 20})

	// Možemo i eksplicitno pozvati polja.
	fmt.Println(person{name: "Alice", age: 30})

	// Zanemarena polja će imati zero-value vrednosti.
	fmt.Println(person{name: "Fred"})

	// Prefiks `&` vraća pokazivač na strukturu.
	fmt.Println(&person{name: "Ann", age: 40})

	// Normalno se koriste konstruktivne funkcije.
	fmt.Println(newPerson("Jon"))

	// Dobijamo vrednosti polja pomoću tačke.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Možemo koristiti tačke i sa pokazivačima na strukture,
	// pokazivači su automatski dereferencirani.
	sp := &s
	fmt.Println(sp.age)

	// Vrednosti polja struktura su promenjiva.
	sp.age = 51
	fmt.Println(sp.age)

	// Ne moramo da dodeljujemo ime strukturi ako je koristimo
	// samo za jednu varijablu. Vrednost može da ima anonimnu
	// strukturu. Ova tehnika se često koristi za
	// [testove](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
