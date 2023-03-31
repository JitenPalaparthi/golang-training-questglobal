package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Server started on port 8081")
	baseURL := os.Getenv("DEMOENV_SERVICE")
	port:=os.Getenv("PORT")
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong") // environment variable
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		url:=fmt.Sprintf("http://%s:%s/getconn",baseURL,port)
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Fprintln(w, string(bytes))
	})

	http.ListenAndServe(":8081", nil)
}
