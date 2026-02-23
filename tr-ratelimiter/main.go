package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tren03/tr-coreutils/tr-ratelimiter/algorithms"
	"github.com/tren03/tr-coreutils/tr-ratelimiter/ratelimiter"
	"github.com/tren03/tr-coreutils/tr-ratelimiter/utils"
)

const ALLOW_HOST = "localhost"

// closure which returns a handle func
// func RateLimiterMW(func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
// 	rate

// }

/*
To Test the rate limiter, we can add a X-User-ID header
which would be used as key to rate limit
*/
func fail(w http.ResponseWriter) {
	w.WriteHeader(400)
	fmt.Fprint(w, "<h1> ID NOT FOUND </h1>")
}

func ratelimit(rl ratelimiter.RatelimterResponse, w http.ResponseWriter) {
	w.WriteHeader(429)
	jsonData, err := json.Marshal(rl)
	if err != nil {
		fmt.Println("err marshalling json, returning static resp", err)
		jsonData = []byte("static")
	}
	message := fmt.Sprintf("<h1> RATE LIMITED \n %v </h1>", string(jsonData))
	fmt.Println(message)
	fmt.Fprint(w, message)
}

type RateLimiterMW struct {
	rateLimiter *ratelimiter.Ratelimiter
}

func (rl *RateLimiterMW) ServeHTTP(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-user-id") // Must include userid here

		// 2. Handle the Preflight (OPTIONS) request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// fetch the user_id from headers
		fmt.Println("request recieved", *r)
		userIDHeader := r.Header.Get("x-user-id")
		if len(userIDHeader) == 0 {
			fmt.Println("user id not found in header")
			fail(w)
			return
		}
		userID := userIDHeader[0]
		rlResponse := rl.rateLimiter.Allow(string(userID))
		if !rlResponse.Allow {
			fmt.Println("You have been ratelimited!!!", userID)
			ratelimit(rlResponse, w)
			return
		}
		next(w, r)
	}

}

func greet(w http.ResponseWriter, r *http.Request) {
	userIDHeader := r.Header.Get("x-user-id")
	userID := userIDHeader[0]
	fmt.Fprintf(w, "Hello World!", userID)
}

func main() {
	repo := algorithms.NewInMemoryTokenBucketRepo()
	tokenBucket := algorithms.NewTokenBucket(5, 5, repo, utils.TimeNow)
	ratelimiter := ratelimiter.NewRateLimiter(tokenBucket)
	rmw := RateLimiterMW{rateLimiter: ratelimiter}
	http.Handle("/hello", rmw.ServeHTTP(greet))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
