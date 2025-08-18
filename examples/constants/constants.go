// Go podržava konstante, _constants_, za karaktere, string-ove, bool-ove,
// i numeričke vrednosti (npr. int, float).

package main

import (
	"fmt"
	"math"
)

// Pomoću `const` deklarišemo konstantu.
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` može da se pojavi gde god i `var`
	// može.
	const n = 500000000

	// Konstantni izrazi izvode aritmetiku
	// sa proizvoljnom tačnošću.
	const d = 3e20 / n
	fmt.Println(d)

	// Numerička konstanta nema tip dok joj se ne dodeli.
	// Ovo se postiže eksplicitnom konverzijom.
	fmt.Println(int64(d))

	// Numerička konstanta može da dobije tip i ako
	// je iskoristimo u kontekstu. Npr. ako je iskoristimo
	// kao varijablu u pozivu funkcije koja očekuje neki tip.
	// Ovde npr. `math.Sin` očekuje `float64`.
	fmt.Println(math.Sin(n))
}
