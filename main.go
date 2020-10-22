package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type ErrorResponse struct {
	Detail string `json:"detail"`
}

type Response struct {
	IsDisposable bool `json:"isDisposable"`
}

var domains []string

func handler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryParams := r.URL.Query()

	email := queryParams.Get("email")

	w.Header().Set("Content-Type", "application/json")

	if email == "" {
		errorResponse, _ := json.Marshal(ErrorResponse{
			Detail: "email query parameter should be set.",
		})

		w.WriteHeader(400)
		w.Write(errorResponse)

		return
	}

	parts := strings.Split(email, "@")
	domain := parts[len(parts)-1]

	isDisposable := Find(domains, domain)

	jsonBytes, _ := json.Marshal(Response{
		IsDisposable: isDisposable,
	})

	w.Write(jsonBytes)
}

func main() {
	domainsJsonFile, _ := os.Open("domains.json")

	byteValue, _ := ioutil.ReadAll(domainsJsonFile)

	_ = json.Unmarshal(byteValue, &domains)

	router := httprouter.New()

	router.GET("/", handler)

	log.Println("Starting to listen on 0.0.0.0:80...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}

	return false
}
