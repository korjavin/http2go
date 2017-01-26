package main

import (
	"fmt"
	"net/http"
)

func getmsg(w http.ResponseWriter, r *http.Request) {
	rec := r.FormValue("rec")
	text := r.FormValue("text")
	fmt.Println(rec, text)
}

func main() {
	http.HandleFunc("/kiodsEp0", getmsg)
	http.ListenAndServe(":9090", nil)
}
