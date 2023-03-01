package main

import (
	"employee-demo/employee"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	//e := new(employee.Employee)
	//e1 := employee.New(100, "Jiten", "JitenP@Outlook.Com")

	http.HandleFunc("/employee", AddEmployee)

	fmt.Println("Server started and listening on port 50060")
	if err := http.ListenAndServe(":50060", nil); err != nil {
		fmt.Println(err)
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		e1 := employee.Employee{}
		err := json.NewDecoder(r.Body).Decode(&e1)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		} else {
			fmt.Fprint(w, e1)
			f, err := os.Create("files/" + fmt.Sprint(e1.Id))
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}
			_, err = f.WriteString(fmt.Sprintln(e1))
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}
			return
		}
	} else if r.Method == "GET" {
		w.Write([]byte("No employee data found"))
		return
	} else {
		w.WriteHeader(405)
		w.Write([]byte("not implemented"))
		return
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

}
