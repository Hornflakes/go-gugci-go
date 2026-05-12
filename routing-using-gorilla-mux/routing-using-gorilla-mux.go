package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/bands/{band}/songs/{song}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		band := vars["band"]
		song := vars["song"]

		fmt.Fprintf(w, "you've requested the song: %s by band %s\n", song, band)
	})

	http.ListenAndServe(":8080", r)
}
