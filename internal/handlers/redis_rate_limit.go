package handlers

import (
	"net/http"
	"restApiCase/internal/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

func RateLimitedRedisHandler(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("X-Token")

		if token == "" {
			utils.HandleError(w, http.StatusUnauthorized, "Missing token")
			return
		}

		key := "rate_limit: " + token
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			utils.HandleError(w, http.StatusInternalServerError, "Resis error")
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, time.Minute)
		}

		if count > 10 {
			utils.HandleError(w, http.StatusTooManyRequests, "Rate limit exceeded")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, "Request accepted")

	}
}
