package main

import (
	"net/http"

	"baas/Godeps/_workspace/src/github.com/rs/cors"
	"github.com/codegangsta/negroni"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(mux)
	n.Run(":3000")
}
