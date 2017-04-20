package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	GRequest = make(chan int, 0)
)

func main() {
	println("rate-limit v1")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		GRequest <- 1
		w.Write([]byte("ok"))
	})

	go v1_timer()

	addr := "0.0.0.0:7000"
	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}
	fmt.Println("http listening", addr)
	fmt.Println(s.ListenAndServe())
}

func v1_timer() {
	ticker := time.Tick(time.Millisecond * 200)

	for range ticker {
		req := <-GRequest
		fmt.Println(time.Now(), "Handling request", req)
	}
}