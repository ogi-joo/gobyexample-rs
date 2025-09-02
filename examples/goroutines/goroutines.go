// Go-rutina (goroutine) je `lightweight` nit (thread) za izvršenja.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Recimo da imamo poziv funkcije `f(s)`. Ovako
	// inače zovemo funkciju, koja se pokreće
	// sinhrono (sekvencijalno u kodu).
	f("direct")

	// Da bi je izveli u nekoj niti go-rutine,
	// koristimo `go f(s)`. Ova nova go-rutinska će
	// se izvršavati istovremeno sa gornjom funkcijom.
	go f("goroutine")

	// Takođe možemo započeti go-rutinu i sa anonimnim
	// funkcijskim pozivom.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Sada se naše dve funkcije izvršavaju asinhrono
	// u odvojenim go-rutinama. Sačekaćemo da se završe,
	// jer ne zavise od funkcije koja ih zatvara
	// (za robustniji pristup, koristimo [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
