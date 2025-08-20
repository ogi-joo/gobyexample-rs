// _Mape_ su ugrađeni Go tipovi, drugačije poznati kao [asocijativni nizovi](https://sr.wikipedia.org/wiki/%D0%90%D1%81%D0%BE%D1%86%D0%B8%D1%98%D0%B0%D1%82%D0%B8%D0%B2%D0%BD%D0%B8_%D0%BD%D0%B8%D0%B7)
// (ponekad zvani i _hash_ ili rečnici, _dicts_, u drugim programskim jezicima).
// key - ključ, val (value) - vrednost

package main

import (
	"fmt"
	"maps"
)

func main() {

	// Da bi napravili praznu mapu, koristimo ugrađen `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Postavljamo key/value parove koristeći `name[key] = val`
	// sintaksu.
	m["k1"] = 7
	m["k2"] = 13

	// Ispisivanje mape uz pomoć `fmt.Println` će ispisati sve
	// key/value parove.
	fmt.Println("map:", m)

	// Možemo da izvučemo vrednost pomoću ključa
	// koristeći `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn't exist, the
	// [zero value](https://go.dev/ref/spec#The_zero_value) of the
	// value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Ugrađen `len` vraća broj key/value
	// parova kada je pozvan na neku mapu.
	fmt.Println("len:", len(m))

	// Ugrađen `delete` briše key/value par iz
	// neke mape.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Da bi izbrisali *sve* key/value parove iz mape, koristimo
	// ugrađenu `clear` funkciju.
	clear(m)
	fmt.Println("map:", m)

	// Kada izvlačimo vrednost iz neke mape uz pomoć ključa,
	// postoji opcionalna vrednost koju nam Go daje na return.
	// Ova vrednost nam vraća boolean true ako je ključ ikad
	// postojao. Tako možemo da raspoznamo nepostojeće ključeve
	// od onih nultih koje imaju vrednost kao npr. `0` ili `""`.
	// U ovim linijama nismo tražili samu vrednost,
	// pa smo je ignorisali _blanko znakom_ `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Kao i nizove i isečke, mape možemo takođe
	// deklarisati i inicijalizovati u jednoj liniji.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Go paket `maps` sadrži brojne korisne
	// funkcije za upravljanje mapama.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
