package main

import (
	"log"
	"os/exec"
	"strings"
)

func send(m msg) error {
	rec := strings.Replace(m.rec, "'", "\\'", -1)
	text := strings.Replace(m.text, "'", "\\'", -1)
	text = strings.Replace(text, "\\r\\n", "<br>", -1)
	text = strings.Replace(text, "\\n", "<br>", -1)
	text = strings.Replace(text, "\\r", "<br>", -1)

	cmd := "echo \"msg " + rec + " " + text + " \" | sudo docker exec -i tg nc 127.0.0.1 1234 -w 1"
	log.Println(cmd)
	lsCmd := exec.Command("sh", "-c", cmd)
	_, err := lsCmd.Output()
	return err
}
