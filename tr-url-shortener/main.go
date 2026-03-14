package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello world") })
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
