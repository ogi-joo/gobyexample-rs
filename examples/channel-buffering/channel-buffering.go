// Po default-u, kanali su _nebaferovani, unbuffered_,
// što znači da će oni moći da primaju slanja (`chan <-`)
// samo ako druga strana ima primanja (`<- chan`) spremna
// za poslatu vrednost. _Baferovani kanali_ primaju ograničen
// broj vrednosti bez potrebe za risiverom.

package main

import "fmt"

func main() {

	// Ovde kreiramo kanal string-ova koji ima
	// bafer za 2 vrednosti.
	messages := make(chan string, 2)

	// Zato što je ovaj kanal baferovan, možemo slati
	// ove vrednosti u kanal bez potrebe za
	// istovremenim risiverom.
	messages <- "buffered"
	messages <- "channel"

	// Kasnije ih možemo prihvatiti kao i obično.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
