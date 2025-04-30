package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	reasons           []string
	rateLimitEnabled  = os.Getenv("ENABLE_RATE_LIMIT") == "true"
	clientLimiters    = make(map[string]*rate.Limiter)
	clientLimitersMux sync.Mutex
	limiterRate       = rate.Every(time.Minute / 10) // 10 requests/min
	limiterBurst      = 10
	rng               = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func loadReasons() {
	data, err := os.ReadFile("reasons.json")
	if err != nil {
		log.Fatalf("Failed to read reasons.json: %v", err)
	}
	if err := json.Unmarshal(data, &reasons); err != nil {
		log.Fatalf("Invalid JSON: %v", err)
	}
}

func getIP(r *http.Request) string {
	ip := r.RemoteAddr
	if strings.Contains(ip, ":") {
		ip, _, _ = net.SplitHostPort(ip)
	}
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ip = strings.Split(xff, ",")[0]
	}
	return ip
}

func getLimiter(ip string) *rate.Limiter {
	clientLimitersMux.Lock()
	defer clientLimitersMux.Unlock()

	limiter, exists := clientLimiters[ip]
	if !exists {
		limiter = rate.NewLimiter(limiterRate, limiterBurst)
		clientLimiters[ip] = limiter
	}
	return limiter
}

func noHandler(w http.ResponseWriter, r *http.Request) {
	if rateLimitEnabled {
		ip := getIP(r)
		if !getLimiter(ip).Allow() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error":"Too many requests, please try again later."}`))
			return
		}
	}

	reason := reasons[rng.Intn(len(reasons))]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"reason": reason})
}

func main() {
	loadReasons()

	http.HandleFunc("/no", noHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("No-as-a-Service running on port %s (rate limit: %v)", port, rateLimitEnabled)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
