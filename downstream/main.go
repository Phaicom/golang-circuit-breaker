package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/phaicom/golang-circuit-breaker/downstream/handlers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	PORT := ":28080"

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/message", handlers.MessageHandler).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)
	fmt.Printf("Http server running on port %v\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, n))
}
