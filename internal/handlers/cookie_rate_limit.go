package handlers

import (
	"net/http"
	"restApiCase/internal/utils"
	"sync"
	"time"
)

var tokenRequestCounts = make(map[string][]time.Time)
var muCookie sync.Mutex

func RateLimitedTestCookieHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session_id")
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, "Missing session_id cookie")
		return
	}

	token := cookie.Value

	muCookie.Lock()
	defer muCookie.Unlock()

	now := time.Now()
	requests := tokenRequestCounts[token]

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
	tokenRequestCounts[token] = validRequests

	utils.RespondWithJSON(w, http.StatusOK, "Requset accepted")
}
