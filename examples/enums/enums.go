// _Enum tipovi_ (enum) su specijalan slučaj
// [zbirnih tipova](https://en.wikipedia.org/wiki/Algebraic_data_type).
// Enum tip ima fiksan broj mogućih vrednosti,
// svaka sa svojim imenom. Go nema direktnu
// sintaksu za enum tipove, ali enum tipovi
// su veoma laki za implementaciju koristeći se
// već postojećim Go konceptima.

package main

import "fmt"

// Naš enum tip `ServerState` ima njegov `int` tip.
type ServerState int

// Moguće vrednosti `ServerState`-a su definisane kao
// konstante. Sintaksa [iota](https://go.dev/ref/spec#Iota)
// generiše uzastupne vrednosti konstante automatski; u ovom
// primeru: 0, 1, 2, itd.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Implementiranjem [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// interfejsa, vrednosti `ServerState`-a se mogu ispisivati ili
// konvertovati u string.
//
// Ovo može postati zapleteno ako ima više mogućih vrednosti.
// U takvom slučaju [stringer tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// zajedno sa `go:generate` se može koristiti za automaciju
// procesa. Pogledaj [ovaj post](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// za bolje objašnjenje.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Ako imamo vrednost tipa `int`, ne možemo je dati `transition`-u - jer
	// će se kompajler buniti oko neslaganja tipova. Ovo daje dodatnu
	// sigurnost za enum tipiziranje prilikom kompajliranja.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition funkcija emulira status servera;
// uzima trenutni i vraća sledeći.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some predicates here to
		// determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
