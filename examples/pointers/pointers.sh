# `zeroval` ne menja `i` u `main`-u, ali
# `zeroptr` menja jer predstavlja referencu
# na memorijsku adresu te varijable.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
