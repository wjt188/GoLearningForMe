package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIp(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}
	fmt.Println("forwarded: ", forwarded)
	fmt.Println("r.RemoteAddr: ", r.RemoteAddr)
	return r.RemoteAddr
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": GetIp(r),
	})
	w.Write(resp)
}

func main() {
	http.HandleFunc("/test", TestHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
