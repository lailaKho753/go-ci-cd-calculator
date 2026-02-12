package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Response struct {
	Result float64 `json:"result"`
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

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	res := Response{
		Result: Subtract(req.A, req.B),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	res := Response{
		Result: Multiply(req.A, req.B),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	res := Response{
		Result: Divide(req.A, req.B),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func powerHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	res := Response{
		Result: Power(req.A, req.B),
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
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divideHandler)
	http.HandleFunc("/power", powerHandler)
	http.ListenAndServe(":8080", nil)
}
