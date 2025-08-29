# Da bi pokrenuli program, ubacujemo kod u `hello-world.go`
# koristimo `go run`.
$ go run hello-world.go
hello world

# Ponekad ćemo zaželeti da imamo naš program u 
# binarnoj datoteci (npr .exe). Ovo možemo uraditi pomoću `go build`.
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# Zatim možemo pokrenuti aplikaciju direktno:
$ ./hello-world
hello world

# Sada kada znamo kako ispisati i napraviti jednostavan Go program, možemo
# se baciti na teže stvari.