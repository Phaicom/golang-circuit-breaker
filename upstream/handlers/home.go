package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/afex/hystrix-go/hystrix"
)

type Message struct {
	Message string `json:"message"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	commandName := os.Getenv("COMMAND_NAME")
	output := make(chan *http.Response, 1)
	errors := hystrix.Go(commandName, func() error {
		// talk to other services
		res, err := callChargeProducerAPI()
		// err := callWithRetry()

		if err == nil {
			output <- res
		}
		return err
	}, func(err error) error {
		log.Println("fallbackErrorDesc", err.Error())
		return errors.New("Service currently unavailable")
	})

	select {
	case res := <-output:
		// success
		message := new(Message)
		err := json.NewDecoder(res.Body).Decode(&message)
		if err != nil {
			log.Println(err)
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"data": message.Message})
	case err := <-errors:
		// failure
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
	}
}

func callChargeProducerAPI() (*http.Response, error) {
	res, err := http.Get("http://localhost:28080/api/v1/message")

	if err != nil {
		return nil, errors.New("503 error")
	}

	return res, nil
}

// func callWithRetry() ([]byte, error) {
// 	for index := 0; index < 3; index++ {
// 		// call producer API
// 		err := callChargeProducerAPI()
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	// adding backoff
// 	// adding jitter
// 	return nil, nil
// }
