package main

import (
	"baas/Godeps/_workspace/src/github.com/rs/cors"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := http.HandlerFunc(httpHandler)
	log.Fatal(http.ListenAndServe(getPort(), c.Handler(handler)))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := booleanGenerator()
	json.NewEncoder(w).Encode(map[string]bool{"data": data})
}

func booleanGenerator() bool {
	booleans := [2]bool{true, false}
	index := rand.Intn(2)
	return booleans[index]
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
