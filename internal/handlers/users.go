package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"restApiCase/internal/models"
	"restApiCase/internal/repository"
	"restApiCase/internal/utils"
	"strconv"
	"strings"
)

func parseUserID(w http.ResponseWriter, path string) (int, bool) {
	idStr := strings.TrimPrefix(path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid user ID")
		return 0, false
	}
	return id, true
}

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.UserInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.HandleError(w, http.StatusBadRequest, "Invalid input")
			return
		}
		user, err := repository.CreateUser(db, input.Name)
		if err != nil {
			utils.HandleError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, user)
	}
}

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := parseUserID(w, r.URL.Path)
		if !ok {
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
		utils.RespondWithJSON(w, http.StatusOK, user)
	}
}

func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := parseUserID(w, r.URL.Path)
		if !ok {
			return
		}
		input := models.UserInput{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.HandleError(w, http.StatusBadRequest, "Invalid input")
			return
		}
		user, err := repository.UpdateUser(db, id, input.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.HandleError(w, http.StatusNotFound, "User not found")
			} else {
				utils.HandleError(w, http.StatusInternalServerError, "Failed to update user")
			}
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, user)
	}
}
