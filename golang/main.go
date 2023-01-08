package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/service1", service1)
	http.HandleFunc("/service2", service2)
	http.HandleFunc("/service3", service3)
	http.ListenAndServe(":8000", nil)
}

func service1(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8000/service2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Service 1: %s", responseData)
}

func service2(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8000/service3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Service 2: %s", responseData)
}

func service3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service 3")
}
