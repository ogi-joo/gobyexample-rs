# Vraća nam se vrednost `"one"` pa `"two"` kao
# što smo očekivali.
$ time go run select.go 
received one
received two

# Primetimo da je ukupno vreme izvršenja ~2 sekunde
# jer obe 1 i 2 sekunde izvršavaju `Sleeps` istovremeno.
real	0m2.245s
