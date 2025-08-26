// Počevši od verzije 1.18, Go podržava
// _generike_, drugačije poznati kao _tipovni parametri_.

package main

import "fmt"

// Kao primer generičke funkcije, `SlicesIndex` uzima
// slice bilo kog `comparable` tipa i element tog tipa
// i vraća index prvog prikazivanja `v` vrednosti u `s` array-u
// ili -1 ako nije prisutan. Ograničenje `comparable`
// omogućava nam da vrednost ovog tipa upoređujemo sa
// `==` i `!=` operatorima. Za detaljnije objašnjenje
// ovog tipovnog potpisa, pogledaj [ovaj blog post](https://go.dev/blog/deconstructing-type-parameters).
// Primetimo da se ova funkcija takođe nalazi u standardnoj
// biblioteci - [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Kao primer generičkog tipa, `List` je jednosmerno
// vezana lista sa vrednostima bilo kog tipa.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Možemo da generišemo metode sa generičkim tipovima
// isto kao i sa regularnim... ali moramo da držimo tipovne
// parametre robustno. Tip je `List[T]`, a ne `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements vraća sve List elemente kao slice.
// U sledećem primeru, videćemo više idiomatsko
// iteriranje preko svih elemenata generičkog tipa.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// Kada zovemo generičku funkciju, često se možemo
	// osloniti na tipovni zaključak - _type inference_.
	// Primetimo da ne moramo da specifiramo tipove
	// `S`-a i `E`-a kada zovemo `SlicesIndex` - kompajler
	// ih priziva automatski.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... Ali ih možemo i eksplicitno prizvati.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
