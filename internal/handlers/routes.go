package handlers

import (
	"database/sql"
	"net/http"
	"restApiCase/internal/utils"
)

func SetupRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/sum", SumHandler)
	mux.HandleFunc("/multiply", MultiplyHandler)

	mux.HandleFunc("/users", CreateUserHandler(db))
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetUserHandler(db)(w, r)
		case http.MethodPut:
			UpdateUserHandler(db)(w, r)
		default:
			utils.HandleError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
	return mux
}
