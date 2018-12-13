package handlers

import "net/http"

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}
