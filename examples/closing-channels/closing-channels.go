// _Zatvaranje_ kanala znači da se vrednosti
// više ne mogu slati. Ovo je korisno da označimo
// završetak kanala našim risiverima.

package main

import "fmt"

// U ovom primeru koristićemo kanal `jobs`
// da objavimo završetak `main()` go-rutine.
// Kada više ne budemo imali task-ove za
// naše worker-e `zatvorićemo` kanal `jobs`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Ovo je go-rutina. Ona non-stop prima taskove
	// iz `jobs` sa `j, more := <-jobs`. U ovoj
	// formi primanja sa dve vrednosti, vrednost `more`
	// će biti `false` ako je `jobs` zatvoren i sve
	// vrednosti u kanalu su primljene.
	// Koristimo ovu prednost da pošaljemo vrednost u
	// `done` kada su svi taskovi iz `jobs` gotovi.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Ovde šaljemo 3 "posla" na worker preko
	// kanala `jobs` i potom ga zatvaramo.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Čekamo `done` kanal pomoću tehnike
	// [sinhronizacije](channel-synchronization) koju smo
	// ranije videli.
	<-done

	// Čitanje iz zatvorenog kanala je odmah uspešno,
	// vraćajući `zero-valued` vrednost neke vrednosti (ovde `int`).
	// Opcionalni drugi parametar je `true` ako je
	// vrednost primljena nakon uspešnog slanja na kanal,
	// ili `false` ako je kanal zatvoren i prazan.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
