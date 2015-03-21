package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var nodeServerRootURL = "http://127.0.0.1:3000/"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?module=app&props={}", nodeServerRootURL), nil)
	if err != nil {
		log.Printf("failed to create request. error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to do request. error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body. error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(string(body)))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
