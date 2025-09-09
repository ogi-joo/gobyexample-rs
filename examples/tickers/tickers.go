// [Tajmeri](timers) su za kada želimo nešto da
// uradimo u budućnosti - _ticker-i_ su kada
// želimo nešto da uradimo sa ponavljanjem u regularnim
// intervalima. Ovo je primer ticker-a koji periodično
// "kuca" (tick) dok ga ne zaustavimo.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Ticker-i imaju sličan mehanizam kao tajmeri:
	// kanal kome se šalju vrednosti. Ovde ćemo da
	// koristimo ugrađeni `select` da čekamo
	// vrednosti koje pristižu svakih 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Ticker-i se mogu stopirati kao tajmeri. Jednom kada
	// je stopiran, ticker neće primati više vrednosti na
	// njegov kanal. Zaustavićemo naš nakon 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
