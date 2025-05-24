package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Result struct {
	A     float64 `json:"a"`
	B     float64 `json:"b"`
	Value float64 `json:"result"`
}

func errHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

func parseParams(r *http.Request) (float64, float64, error) {
	aNum := r.URL.Query().Get("a")
	bNum := r.URL.Query().Get("b")

	a, err1 := strconv.ParseFloat(aNum, 64)
	if err1 != nil {
		return 0, 0, err1
	}

	b, err2 := strconv.ParseFloat(bNum, 64)
	if err2 != nil {
		return 0, 0, err2
	}

	return a, b, nil
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		errHandler(w, err)
		return
	}

	res := Result{A: a, B: b, Value: a + b}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		errHandler(w, err)
		return
	}

	res := Result{A: a, B: b, Value: a * b}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func main() {
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	fmt.Println("The server is running on port 8080, example http://localhost:8080/sum?a=5&b=7 or http://localhost:8080/multiply?a=5&b=7")
	http.ListenAndServe(":8080", nil)
}
