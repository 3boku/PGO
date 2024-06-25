package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
)

func fibonacci(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := 50000
		result := fibonacci(n)
		w.Write([]byte(fmt.Sprintf("Fibonacci(%d) = %s", n, result.String())))
	})

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
