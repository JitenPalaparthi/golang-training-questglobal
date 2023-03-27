package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	port string
)

func main() {
	fmt.Println("server started and working")

	port = os.Getenv("PORT")

	if port == "" {
		flag.StringVar(&port, "port", "8080", "--port=8080")
		flag.Parse()
	}

	count := 0
	go func() {
		for {
			count++
			time.Sleep(time.Second * 10)
			fmt.Println("The server is running from the past ", count*10, "seconds")
		}
	}()

	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	router.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello: %v\n", vars["name"])
	}).Methods("GET")

	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "pong")
	// })

	http.Handle("/", router)
	http.ListenAndServe(":"+port, nil)
}

// push this code to docker.
// compile and run inside docker
