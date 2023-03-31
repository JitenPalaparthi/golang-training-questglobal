package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Server started on port 8080")
	http.HandleFunc("/getconn", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, os.Getenv("DBCONN")) // environment variable
	})

	http.HandleFunc("/call",func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprint(w, os.Getenv("DBCONN")) // environment variable
	})
	http.ListenAndServe(":8080", nil)
}
