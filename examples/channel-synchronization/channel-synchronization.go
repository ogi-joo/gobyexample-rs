// Možemo da koristimo kanale da sinhronizujemo
// izvršenja u različitim go-rutinama. Ovo je
// primer gde blokiramo prihvatanje i čekamo
// go-rutinu da se izvrši do kraja.
// Za čekanje na nekoliko go-rutina da se izvrše,
// [WaitGroup](waitgroups) je možda bolji izbor.

package main

import (
	"fmt"
	"time"
)

// Ovo je funkcija koju ćemo koristiti u go-rutini.
// Kanal `done` će se biti koristan da obavesti
// druge go-rutine da se ova funkcija izvršila.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Šaljemo vrednost i obaveštavamo da smo `done`.
	done <- true
}

func main() {

	// Počinjemo worker go-rutinu, dajući joj kanal
	// da nas obavesti nazad.
	done := make(chan bool, 1)
	go worker(done)

	// Blokiramo dalje izvršenje dok ne dobijemo
	// obaveštenje iz kanala.
	<-done
}
