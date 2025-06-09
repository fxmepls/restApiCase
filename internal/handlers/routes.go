package handlers

import (
	"database/sql"
	"net/http"
	"restApiCase/internal/utils"

	"github.com/redis/go-redis/v9"
)

func SetupRoutes(db *sql.DB, redisClient *redis.Client) *http.ServeMux {
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

	mux.HandleFunc("/test", RateLimitedTestHandler())
	mux.HandleFunc("/cookie", RateLimitedTestCookieHandler)
	mux.HandleFunc("/redis", RateLimitedRedisHandler(redisClient))
	return mux
}
