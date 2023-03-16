package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	portNumber := 5555
	dirpath := "."

	fmt.Println("[*] Start Server")
	fmt.Printf("[*] listening on %v\n", portNumber)

	http.Handle("/", http.FileServer(http.Dir(dirpath)))

	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[*] End Server")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
