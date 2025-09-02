# Kada pokrenemo naš go program, prvo vidimo output
# naše sinhrone funkcije, a potom output naše dve
# go-rutine. Output naših go-rutina može biti
# isprepleten, ovo je zato što se one izvršavaju
# istovremeno u Go programu.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Sledeće što ćemo pogledati je dodatak na go-rutine u
# asinhronim Go programima: kanali, channels.
