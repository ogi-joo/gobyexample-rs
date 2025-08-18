// Go ima različite tipove vrednosti, neki od njih su: niske (strings),
// ceo broj (integer), float, bulov tip (boolean), itd. Zarad
// lakšeg sporazumevanja, često ćemo koristiti englicizovane reči.
// Evo nekoliko primera:

package main

import "fmt"

func main() {

	// String-ovi mogu da se nadovezuju jedan na drugi pomoću `+`.
	fmt.Println("go" + "jezik")

	// Celi i realni brojevi
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Boolean. Praktično važe ista pravila kao i u drugim programskim jezicima
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
