// _Kanali, Channels_ su "cevi" koje spajaju istovremene
// go-rutine. Možemo da šaljemo vrednosti iz jedne
// go-rutine i prihvatimo iste u drugim go-rutinama.

package main

import "fmt"

func main() {

	// Pravimo novi kanal sa `make(chan val-type)`.
	// Kanali su tipovani u odnosu na vrednost koju
	// prenose.
	messages := make(chan string)

	// _Šaljemo_ vrednost u kanal koristeći `channel <-`
	// sintaksu. Ovde, šaljemo `"ping"` u `messages`
	// kanal koji smo gore napravili iz nove go-rutine.
	go func() { messages <- "ping" }()

	// Sintaksa `<-channel` _prima_ vrednost iz kanala.
	// Ovde, primamo `"ping"` poruku koju smo poslali
	// gore i ispisujemo je.
	msg := <-messages
	fmt.Println(msg)
}
