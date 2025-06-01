package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"restApiCase/internal/models"
	"restApiCase/internal/repository"
	"restApiCase/internal/utils"
	"strconv"
	"strings"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.HandleError(w, http.StatusBadRequest, "Failed to read body")
			return
		}
		fmt.Println("Request body:", string(bodyBytes))

		// теперь восстанови тело, чтобы json.Decoder мог его прочитать
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		input := models.UserInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.HandleError(w, http.StatusBadRequest, "Invalid input")
			return
		}
		user, err := repository.CreateUser(db, input.Name)
		if err != nil {
			utils.HandleError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.HandleError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}
		user, err := repository.GetUser(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.HandleError(w, http.StatusNotFound, "User not found")
			} else {
				utils.HandleError(w, http.StatusInternalServerError, "Failed to get user")
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
