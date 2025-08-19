// Grananje sa `if` i `else` u Go-u
// je veoma jednostavno.

package main

import "fmt"

func main() {

	// Klasičan primer
	if 7%2 == 0 {
		fmt.Println("7 je paran")
	} else {
		fmt.Println("7 je neparan")
	}

	// `if` grana bez else.
	if 8%4 == 0 {
		fmt.Println("8 je deljivo sa 4")
	}

	// Logički operatori `&&` i `||` su veoma
	// korisni u grananju.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("jedan od brojeva 8 ili 7 je paran")
	}

	// Svaka varijabla deklarisana u ovom bloku se može
	// koristiti i u svim donjim granama else-if-a i else-a
	if num := 9; num < 0 {
		fmt.Println(num, "je negativan")
	} else if num < 10 {
		fmt.Println(num, "ima jednu cifru")
	} else {
		fmt.Println(num, "ima više cifra")
	}
}

// Primećujemo da if-else grane ne zahtevaju zagrade
// ali telo zahteva vitičaste zagrade.
