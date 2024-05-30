package main

import "net/http"

//if you want to define a HTTP handler
//standard library expects
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
