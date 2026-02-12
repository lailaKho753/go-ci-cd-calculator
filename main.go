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

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func requirePost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{
				Error: "method not allowed",
			})
			return
		}
		next(w, r)
	}
}

func decodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "invalid request body",
		})
		return false
	}
	return true
}

func handleOperation(
	w http.ResponseWriter,
	r *http.Request,
	op func(float64, float64) (float64, error),
) {
	var req Request

	if !decodeJSON(w, r, &req) {
		return
	}

	result, err := op(req.A, req.B)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, Response{
		Result: result,
	})
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	handleOperation(w, r, Add)
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	handleOperation(w, r, Subtract)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	handleOperation(w, r, Multiply)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	handleOperation(w, r, Divide)
}

func powerHandler(w http.ResponseWriter, r *http.Request) {
	handleOperation(w, r, Power)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/add", requirePost(addHandler))
	http.HandleFunc("/subtract", requirePost(subtractHandler))
	http.HandleFunc("/multiply", requirePost(multiplyHandler))
	http.HandleFunc("/divide", requirePost(divideHandler))
	http.HandleFunc("/power", requirePost(powerHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
	panic(err)
}

}
