package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	fmt.Println("Server is start...")
	return http.ListenAndServe(`:8090`, http.HandlerFunc(webhook))
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")
	_, _ = w.Write([]byte(`
		{
			"response": {
				"text": "Извините, ничего е умею"
			},
			"version": "1.0"
		}
	`))
}
