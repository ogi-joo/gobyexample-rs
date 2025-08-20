# Primetimo da iako je slice drugačiji tip od niza,
# slično su ispisani od strane `fmt.Println`.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Ovaj [zanimljiv blog](https://go.dev/blog/slices-intro)
# koga je napisao Go tim, ima više informacija o 
# implementaciji i dizajniranju programa sa isečcima.

# Sada kada smo pogledali nizove i isečke, razmotrićemo
# još jedan ugrađen tip vrednosti: mape, maps.
