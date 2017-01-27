package main

import (
	"log"
	"net/http"
	"time"
)

type msg struct {
	regdt time.Time
	rec   string
	text  string
}

func getmsg(w http.ResponseWriter, r *http.Request) {
	rec := r.FormValue("rec")
	text := r.FormValue("text")
	m := msg{regdt: time.Now(), rec: rec, text: text}
	log.Printf("Recieved message: %v", m)
	err := send(m)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/kiodsEp0", getmsg)
	http.ListenAndServe(":9090", nil)
}
