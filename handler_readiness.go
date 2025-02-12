package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{} ) 
}