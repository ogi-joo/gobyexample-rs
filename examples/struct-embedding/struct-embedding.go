// Go podržava _embedovanje_ struktura i interfejsa
// radi glađe _kompozicije_ tipova.
// Ovo embedovanje ne bi trebalo pomešati sa [`//go:embed`](embed-directive)
// Go direktivom dostupnom u verzijama 1.16+, koja služi za
// embedovanje fajlova i foldera u binarni kod aplikacije.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Struktura `container`, _embeduje_ `base`. Embedovanje liči
// na polje bez naziva.
type container struct {
	base
	str string
}

func main() {

	// Kada kreiramo strukture sa vrednostima, moramo da
	// inicijalizujemo embedovane strukture eksplicitno;
	// ovde je naziv polja - tip embedovane strukture.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// `base`-ina polja možemo pristupiti direktno preko `co`,
	// npr. `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternativno, možemo da koristimo i punu putanju.
	fmt.Println("also num:", co.base.num)

	// Kako `container` embeduje `base`, metode
	// `base`-a takođe postaju i metode `container`-a.
	// Ovo je primer.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedovane strukture sa metodama se mogu koristiti za
	// davanje implementacije interfejsa drugim strukturama.
	// Ovde `container` sada implementira
	// `describer` interfejs jer on embeduje `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
