// _Tajm-aut, Timeout_ je bitan za programe koji se povezuju
// sa eksternim resursima ili za one koji zavise od vremena
// izvršenja. Implementacija timeout-a u Go-u je jednostavna i
// elegantna zahvaljujući kanalima i `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Zarad primera, recimo da izvršavamo eksterni
	// poziv koji vraća rezultat na kanal `c1`
	// nakon 2s. Primetimo da je kanal baferovan.
	// Ovo je česta praksa koja zaustavlja curenje
	// go-rutina ako se kanal nikad ne pročita.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Ovo je `select` koji implementira timeout.
	// `res := <-c1` čeka neku vrednost, a `<-time.After`
	// šalje vrednost nakon 1s. Zbog toga što `select`
	// nastavlja `case` sa prvom koja se desi, osiguravamo
	// da će se tajm-aut desiti nakon 1s ako se kanal c1
	// ne koristi.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Ako dozvolimo duži timeout od 3s, onda će kanal
	// `c2` biti popunjem u 2s i ispisaćemo tu vrednost.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
