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

	http.HandleFunc("/employee", Employee)

	fmt.Println("Server started and listening on port 50060")
	if err := http.ListenAndServe(":50060", nil); err != nil {
		fmt.Println(err)
	}
}

func Employee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		e1 := employee.Employee{}
		err := json.NewDecoder(r.Body).Decode(&e1)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		} else {
			fmt.Fprint(w, e1)
			f, err := os.Create("files/" + fmt.Sprint(e1.Id)) // to create a file
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
	} else if r.Method == "DELETE" {
		vals := r.URL.Query()
		id := vals.Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no id found"))
			return
		}
		_, err := os.OpenFile("files/"+id, 0, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			w.Write([]byte("no file found"))
			return
		}
		w.Write([]byte(id + " to be deleted"))
		return
	} else {
		w.WriteHeader(405)
		w.Write([]byte("not implemented"))
		return
	}
}
