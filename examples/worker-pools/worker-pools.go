// U ovom primeru implementiraćemo _worker pool_
// (skup radnji) koristeći go-rutine i kanale.

package main

import (
	"fmt"
	"time"
)

// Ovo je worker, pokrenućemo nekoliko instance
// njega istovremeno. Ovi worker-i će primati
// task-ove na `jobs` kanalu i slati rezultate na
// kanal `results`. Sleep-ovaćemo sekundu za svaki
// posao kako bismo simulirali real-world situaciju.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Da bismo koristili naš skup radnji, a oni
	// radili i vraćali rezultate, napravićemo
	// 2 kanala.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Otpočinjemo 3 worker-a, inicijalno blokirani
	// jer i dalje nemaju poslove, jobs.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Ovde šaljemo 5 `poslova` i potom `zatvaramo` taj
	// kanal čime kažemo da nema više poslova za rad.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Na kraju prihvatamo sve rezultate poslova.
	// Ovako i osiguravamo da su se sve go-rutine
	// završile. Alternativa za čekanje više go-rutina
	// je uz pomoć [WaitGroup-e](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
