package handlers

import (
	"net/http"
	"strconv"

	"restApiCase/internal/models"
	"restApiCase/internal/utils"
)

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

func SumHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid params")
		return
	}

	res := models.Result{A: a, B: b, Value: a + b}

	utils.RespondWithJSON(w, http.StatusOK, res)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := parseParams(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid params")
		return
	}

	res := models.Result{A: a, B: b, Value: a * b}

	utils.RespondWithJSON(w, http.StatusOK, res)
}
