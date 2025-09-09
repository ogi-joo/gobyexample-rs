// Da bi čekali na više go-rutina da se završe,
// možemo da koristimo *wait grupe*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Ovo je worker koga ćemo pokretati na svaku go-rutinu.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep da simuliramo skupi task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Ovaj WaitGroup se koristi da čeka sve pokrenute
	// go-rutine da se završe. Bitno: ako je WaitGroup
	// eksplicitno prosleđena funkciji, moramo to uraditi
	// *preko pokazivača*.
	var wg sync.WaitGroup

	// Pokrećemo više go-rutina pomoću `WaitGroup.Go`.
	// Ova metoda nije dostupna u starijim verzijama,
	// budite sigurni da koristite najnoviju Go verziju.
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// Blokiramo dok se sve go-rutine pokrenute od strane `wg`
	// ne završe. Go-rutina se završi kada interna funkcija
	// vrati - `returns`.
	wg.Wait()

	// Primetimo da ovaj pristup nam ne daje laganu
	// propagaciju grešaka od worker-a. Za komplikovanije
	// primere, možemo koristiti
	// [errgroup paket](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
