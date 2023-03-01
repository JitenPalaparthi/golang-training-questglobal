package main

import (
	"fmt"
	"net/http"
)

// Go http servers are usually self hosted
// tomcat, IIS :-- Web Servers
func main() {
	// network-interface address:portnumber
	// 0.0.0.0 : All network interfaces
	// 127.0.0.1 : localhost : lo
	// 192.168.0.1 || 192.168.1.1 Wireless network interface
	// wireshock :- network troubleshoot tool

	// 20 22, 21 80 443,3306, 5432 , 27017
	// > 30000 : applications

	// http://username:password@address:port/example/me?name=quest-global
	// http: scheme/protocol
	// username:password : basic authentication for the site. By default for http it is not given due to securoty reasons and also to access publicly
	// address: host (DNS or direct IP address)
	// port: port that the application is accessable. If no port is given that means it runs on 80
	// /example/me : pattern/virtual path/ resources
	// ?name=quest-global : This is a query string. name is the key and quest-global is a value

	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("server started and running on port:50059")
	if err := http.ListenAndServe(":50059", nil); err != nil {
		fmt.Println("There is an error and unable to proceed")
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// handlers for each path
// http work on verbs/actions : GET,POST,PUT,DELETE,PATCH,OPTIONS,HEAD etc..
