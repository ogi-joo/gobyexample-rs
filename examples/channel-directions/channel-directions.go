// Kada koristimo kanal kao funkcijski parametar, možemo
// specifirati da li je kanal namenjen samo slanju ili
// primanju vrednosti. Ovim povećavamo tipovnu sigurnost
// programa.

package main

import "fmt"

// Ova `ping` funkcija prima samo kanal za slanje
// vrednosti. Ako pokušamo da pozovemo sa kanalom
// za primanje, doći će do greške prilikom kompajliranja.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Funkcija `pong` prima jedan kanal za primanje
// (`pings`) i drugi za slanje (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
