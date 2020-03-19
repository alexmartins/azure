package main

import (
	"log"
	"math"
	"net/http"
)

// Funcao para um calculo muito louco 8 ===== )
func Crash() string {
	x := 0.0001

	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	return "<h1><center>Code.education Rocks! (5)</center></h1>"
}

// Montando Servico HTTP
func Server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Crash()))
}

// Funcao principal do sistema
func main() {
	http.HandleFunc("/", Server)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
