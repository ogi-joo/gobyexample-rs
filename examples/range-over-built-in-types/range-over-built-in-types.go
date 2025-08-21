// _range_ iterira preko elemenata raznih ugrađenih
// data struktura. Pogledaćemo kako da koristimo
// `range` sa nekim od data struktura koje smo
// već pominjali.

package main

import "fmt"

func main() {

	// Ovde koristimo `range` da saberemo brojeve u slice-u.
	// Nizovi se isto ponašaju.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` vraća dve vrednosti, indeks i
	// vrednost na svakoj iteraciji. U gornjem bloku nije
	// nam trebao indeks, pa smo ga ignorisali sa `_`.
	// Ponekad su nam indeksi korisni.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` na mapi iterira preko key/value parova.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` može da iterira i samo po ključevima na mapi.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` na niskama (strings) iterira preko Unicode koda
	// Ovde `range` vraća dve vrednosti, prva je index "rune",
	// druga vrednost je `runa` (runa - int32 oblik Unicode-a).
	// [Niske i Rune](strings-and-runes) za više detalja.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
