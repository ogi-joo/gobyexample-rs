# Očekujemo da imamo tačno 50,000 operacija. Da smo 
# koristili ne-atomski integer, i onda ga povećavali
# sa `ops++`, verovatno bismo dobili drugačiju brojku.
# Na ovaj način go-rutine se ne sudaraju međusobno.
$ go run atomic-counters.go
ops: 50000

# Sledeće što ćemo proći su mutex-i, još jedan alat
# za manipulaciju stanja u Go-u.
