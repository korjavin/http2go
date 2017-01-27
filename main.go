package main

import (
	"log"
	"net/http"
)

type msg struct {
	rec  string
	text string
}

func getmsg(w http.ResponseWriter, r *http.Request) {
	rec := r.FormValue("rec")
	text := r.FormValue("text")
	m := msg{rec: rec, text: text}
	log.Printf("Recieved message: %v", m)
	err := sendmsg(m)
	if err != nil {
		log.Println(err)
	}
}

var (
	users map[string]user
)

func main() {
	users = make(map[string]user)
	loadusers(users)
	go bot_go()
	http.HandleFunc("/kiodsEp0", getmsg)
	http.ListenAndServe(":9091", nil)
}
