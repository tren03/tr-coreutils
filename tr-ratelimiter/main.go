package main

import "fmt"

// import (
// 	"fmt"
// 	"time"

// 	"github.com/tren03/tr-coreutils/tr-ratelimiter/algorithms"
// 	"github.com/tren03/tr-coreutils/tr-ratelimiter/ratelimiter"
// 	"github.com/tren03/tr-coreutils/tr-ratelimiter/utils"
// )

func closure(a int) func() int {
	return func() int {
		a = a + 1
		return a
	}
}

func main() {
	temp := closure(0)
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())
	fmt.Println(temp())

	// repo := algorithms.NewInMemoryTokenBucketRepo()
	// tokenBucket := algorithms.NewTokenBucket(20, 5, repo, utils.TimeNow)
	// ratelimiter := ratelimiter.NewRateLimiter(tokenBucket)
	// for i := range 10 {
	// 	resp := ratelimiter.Allow("testUser")
	// 	fmt.Println(i, resp)
	// 	if !resp.Allow {
	// 		fmt.Println("sleeping for 3 seconds")
	// 		time.Sleep(3 * time.Second)
	// 	}
	// }
}
