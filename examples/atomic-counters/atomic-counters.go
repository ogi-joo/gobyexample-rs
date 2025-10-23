// Primarni mehanizam za kontrolu stanja u Go-u je
// komunikacija preko kanala. Ovo smo na primer videli i
// kod [worker pools](worker-pools). Ali postoje još nekoliko
// opcija za kontrolu stanja. Sada ćemo pogledati korišćenje
// paketa `sync/atomic` uz pomoć _atomskih brojača_ kojima
// pristupaju više go-rutina istovremeno.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Koristićemo atomski integer da prikažemo
	// naš (pozitivan) brojač.
	var ops atomic.Uint64

	// WaitGroup-a će nam pomoći da sačekamo sve
	// go-rutine dok ne završe svoj posao.
	var wg sync.WaitGroup

	// Započećemo 50 go-rutina gde će svaka povećati
	// brojač za jedan - tačno 1000 puta.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// Da atomski inkrementiramo brojač, koristimo `Add`.
				ops.Add(1)
			}
		})
	}

	// Čekamo sve go-rutine.
	wg.Wait()

	// Ovde nijedna go-rutina ne upisuje u 'ops', ali korišćenjem
	// `Load`, možemo sigurno "atomski" čitati vrednost čak i kada
	// go-rutine ih (atomski) upisuju.
	fmt.Println("ops:", ops.Load())
}
