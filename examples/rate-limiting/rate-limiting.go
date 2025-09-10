// [_Kontrola protoka, Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// je bitan mehanizam za kontrolisanje i upotrebu
// resursa i održavanje kvaliteta servisa. Go
// elegantno podržava rate limiting uz go-rutine,
// kanale, i [ticker-e](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Prvo ćemo proći kroz klasičan rate limiting.
	// Želimo da ograničimo dolazeće request-e na 200ms
	// po request-u. Zarad primera, napravićemo i popuniti
	// request kanal sa 5 vrednosti.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Ovaj `limiter` kanal će primati vrednost
	// svakih 200 milisekundi. On glumi regulator
	// protoka u ovoj našoj šemi.
	limiter := time.Tick(200 * time.Millisecond)

	// Ovde čekamo `limiter` kanal da prosledi vrednosti
	// pre nego se nastavi sa request-om, limitiramo for petlju
	// na 1 request svakih 200 milisekundi.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Ponekad želimo da rešavamo kraće praskove, bursts,
	// request-ova ali ujedno i zadržimo naš rate-limiting
	// nakon burst-a. Ovo možemo postići baferovanjem
	// našeg limiter kanala. Ovaj `burstyLimiter`
	// kanal će dozvoliti burst-ove do 3 request-a.
	burstyLimiter := make(chan time.Time, 3)

	// Naš limiter na početku dozvoljava samo 3 odjednom.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// A onda nam svakih 200ms oslobađa mesto za jedan request.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Hajde sada da simuliramo 5 requesta.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// Prva 3 request-a će odmah proći jer limiter ima šta da
	// vrati u njegov kanal kao signal. Poslednja dva će ići
	// na 200ms.
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
