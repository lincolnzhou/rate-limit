package main

import (
	"fmt"
	"net/http"
	"time"
	"sync/atomic"
)

var (
	GRequest = make(chan int, 0)
	Count int64 = 0
)

func main() {
	println("【rate-limit】v2-counter")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		GRequest <- 1
		w.Write([]byte("ok"))
	})

	go clearCount()
	go handleRequest()

	addr := "0.0.0.0:7000"
	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}
	fmt.Println("http listening", addr)
	fmt.Println(s.ListenAndServe())
}

func clearCount() {
	ticker := time.Tick(time.Millisecond * 200)

	go func() {
		for range ticker {
			atomic.StoreInt64(&Count, 0)
		}
	}()
}

func handleRequest() {
	for req := range GRequest {
		if Count < 5 {
			atomic.AddInt64(&Count, 1)
			fmt.Println(time.Now(), "Handling request", req)
		} else {
			fmt.Println(time.Now(), "Denying request", req)
		}
	}
}