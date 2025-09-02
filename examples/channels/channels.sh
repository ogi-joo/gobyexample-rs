# Kada pokrenemo program, `"ping"` poruka je
# uspešno poslata iz jedne go-rutine u drugu
# go-rutinu pomoću našeg kanala.
$ go run channels.go 
ping

# Po default-u, funkcije slanja i primanja zaustavljaju
# se dok obe nisu spremne. Ovo nam je omogućilo da
# korektno sačekamo kraj našeg programa za `"ping"`
# poruku, a da ne moramo da koristimo neku drugu
# sinhronizaciju.
