package main

import (
   "fmt"
   "net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func greeting(message string) string {
	return "<b>" + message + "</b>"
}
