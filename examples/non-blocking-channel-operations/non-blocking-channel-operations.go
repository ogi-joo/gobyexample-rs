// Klasično slanje i primanje preko kanala blokira.
// Međutim, možemo da koristimo `select` sa `default` klauzulom
// i omogućimo _deblokirana_ slanja, primanja, i
// deblokirani više-smerni `select`.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Ovo je deblokirano primanje. Ako je vrednost
	// dostupna na `messages` onda će `select` uzeti
	// slučaj `<-messages` sa tom vrednošću. Ako nema
	// on će se odmah prebaciti na `default` slučaj.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Deblokirano slanje radi slično. Ovde, `msg`
	// ne može da se pošalje na kanal `messages` zato
	// što kanal nema bafer i ne postoji risiver.
	// `Default` slučaj je izabran.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Možemo imati više `slučajeva` iznad `default`-a
	// klauzule da implementiramo više-smerni deblokirani
	// select. Ovde pokušavamo da primimo na oba
	// kanala `messages` i `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
