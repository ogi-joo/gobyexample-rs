// Od verzije 1.23, Go je dodao podršku za
// [iteratore](https://go.dev/blog/range-functions),
// što nam dopušta da koristimo `range` preko
// bilo čega!

package main

import (
	"fmt"
	"iter"
	"slices"
)

// Pogledajmo tip `List` iz
// [prethodnog primera](generics) opet. U tom
// primeru smo imali `AllElements` metodu koja vraća
// slice svih elemenata u listi. Sa Go iteratorima,
// možemo uraditi bolje - kao što je prikazano.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Funkcija All vraća _iterator_, on je u Go-u ustvari funkcija
// sa [posebnim potpisom](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// Iterator funkcija uzima za parametar drugu funkciju,
		// nazvaćemo je `yield` po konvenciji (ali
		// ime je arbitrarno). Iterator funkc. će zvati `yield` za
		// svaki element koji želimo da iteriramo, i primetimo `yield`-ovu
		// return vrednost za prevremenu terminaciju.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteracija ne zahteva bilo kakvu strukturu podataka,
// čak ne mora biti ni konačna! Evo jedne funkcija koja
// vraća iterator preko Fibonacci-evih brojeva: ona se
// vrti u loop-u dok god `yield` vraća `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Zbog toga što `List.All` vraća iterator, možemo je
	// koristiti regularno u `range` loop-u.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Paketi kao [slices](https://pkg.go.dev/slices) imaju
	// brojne korisne funkcije za rad sa iteratorima.
	// Na primer, `Collect` uzima bilo koji iterator i
	// čuva sve njegove vrednosti u slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// Jednom kada loop vidi `break` ili rani `return`, funkcija `yield`
		// data iteratoru će vratiti `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
