package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Result int `json:"result"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	res := Response{
		Result: Add(req.A, req.B),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/add", addHandler)

	http.ListenAndServe(":8080", nil)
}
