package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	srvPort := 4444
	srvHost := "localhost"
	srvURI := srvHost + ":" + strconv.Itoa(srvPort)
	connstring := "http://" + srvURI

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
	log.Printf("Server starting on port:\t%d", srvPort)
	log.Printf("Server connection string:\t%s", connstring)
	http.ListenAndServe(srvURI, r)
}
