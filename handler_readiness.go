package main


import "net/http"

// this is the signature you have to use if you want to define a http handler in the way that
// the go standard library expects
func handlerReadiness(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, 200, struct{}{})
}