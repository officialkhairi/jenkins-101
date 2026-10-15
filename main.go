package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, Greet("Jenkins"))
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, newMux()))
}
