// Go-ov _select_ nam omogućava da čekamo na više
// operacija kanala. Kombinovanje go-rutina i kanala
// sa select-om je moćna stvar u Go-u.

package main

import (
	"fmt"
	"time"
)

func main() {

	// U našem primeru ćemo selektovati preko 2 kanala.
	c1 := make(chan string)
	c2 := make(chan string)

	// Oba kanala će dobiti vrednost posle nekog vremena,
	// ovako demonstriramo npr. asinhrone pozive
	// koji se izvršavaju u istovremenim go-rutinama.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Koristimo `select` za čekanje obe ove vrednosti
	// istovremeno, ispisivajući kako koja stigne.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
