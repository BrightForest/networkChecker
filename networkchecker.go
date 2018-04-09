package main

import (
	"net"
	"html/template"
	"log"
	"net/http"
	"strings"
	"fmt"
	"time"
)

var (
	Info    *log.Logger
	Error   *log.Logger
)

func checkIPandPortAvailable (ip string, port string) bool{
	conn, err := net.DialTimeout("tcp", ip + ":" + port, 3*time.Second)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer conn.Close()
	return true
}

type Responce struct {
	Answer string
}

func checkService(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("checkpage.gtpl")
		answer := Responce{Answer: "No target selected."}
		t.Execute(w, answer)
	} else {
		r.ParseForm()
		var ipAddress = strings.Join(r.Form["ip"], " ")
		var ipPort = strings.Join(r.Form["port"], " ")
		var isAvailable = checkIPandPortAvailable(ipAddress, ipPort)
		if isAvailable == true {
			t, _ := template.ParseFiles("checkpage.gtpl")
			answer := map[string]interface{}{"Answer": "Connection to " + ipAddress + ":" + ipPort + " is OK."}
			t.Execute(w, answer)
		} else {
			t, _ := template.ParseFiles("checkpage.gtpl")
			answer := map[string]interface{}{"Answer": "Connection to " + ipAddress + ":" + ipPort + "  is failed."}
			t.Execute(w, answer)
		}
	}
}

func main() {
	http.HandleFunc("/", checkService)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		Error.Println("error: %v", err)
	}
}