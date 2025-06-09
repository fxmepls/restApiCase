package handlers

import (
	"net/http"
	"restApiCase/internal/utils"
	"sync"
	"time"
)

var requestCounts = make(map[string][]time.Time)
var mu sync.Mutex

func RateLimitedTestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		requests := requestCounts[ip]

		validRequests := []time.Time{}
		for _, t := range requests {
			if now.Sub(t) < time.Minute {
				validRequests = append(validRequests, t)
			}
		}
		if len(validRequests) >= 10 {
			utils.HandleError(w, http.StatusTooManyRequests, "Rate limit exceeded")
			return
		}

		validRequests = append(validRequests, now)
		requestCounts[ip] = validRequests

		utils.RespondWithJSON(w, http.StatusOK, "Requset accepted")
	}
}
