package main

import (
	"encoding/json"
	// "fmt"
	// "html"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := booleanGenerator()
		json.NewEncoder(w).Encode(map[string]bool{"data": data})
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func booleanGenerator() bool {
	booleans := [2]bool{true, false}
	index := rand.Intn(2)
	return booleans[index]
}
