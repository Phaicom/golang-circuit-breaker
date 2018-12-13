package main

import (
	"log"
	"net/http"
	"os"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gorilla/mux"
	"github.com/phaicom/golang-circuit-breaker/upstream/handlers"
	"github.com/urfave/negroni"
)

const commandName = "upstream_api"

func main() {
	PORT := ":28090"
	os.Setenv("COMMAND_NAME", commandName)

	hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
		Timeout:                500,
		MaxConcurrentRequests:  100,
		ErrorPercentThreshold:  50,
		RequestVolumeThreshold: 3,
		SleepWindow:            1000,
	})

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	log.Printf("listening on %s\n", PORT)
	http.ListenAndServe(PORT, n)
}
