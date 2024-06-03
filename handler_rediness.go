package main

import "net/http"

func handlerRediness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
