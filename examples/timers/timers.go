// Često nam zatreba da izvršimo Golang kod u budućnosti,
// ili po nekom intervalu. Go-ove ugrađene funkcije
// _timer_ i _ticker_ nam omogućavaju da oba task-a budu
// laka za implementaciju. Prvo ćemo pogledati u tajmere
// a potom u [ticker-e](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Tajmeri predstavljaju neki događaj u budućnosti. Mi
	// govorimo tajmeru koliko dugo želimo da traje, a on
	// nam daje kanal koji će biti obavešten kada na to dođe
	// vreme. Ovaj tajmer će čekati 2 sekunde.
	timer1 := time.NewTimer(2 * time.Second)

	// Sintaksa `<-timer1.C` blokira tajmer kanal `C`
	// dok on ne pošalje vrednost obaveštavajući nas da
	// je tajmer "opalio".
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Ako poželimo samo da čekamo, možemo koristiti klasičan
	// `time.Sleep`. Jedan od razloga zašto tajmer može biti koristan
	// je taj da možemo da ga cancel-ujemo pre nego što opali.
	// Ovo je primer toga.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Ako damo tajmeru `timer2` dovoljno vremena da opali,
	// možemo dokazati da je stvarno zaustavljen. Ovu liniju
	// možemo staviti odmah iznad `stop2 := timer2.Stop()` kako
	// bismo ovo demonstrirali.
	time.Sleep(2 * time.Second)
}
