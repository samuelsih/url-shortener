package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samuelsih/url-shortener/helper"
)

const port = ":8080"

func main() {
	mux := mux.NewRouter()

	//path is map of shortened URL and real URL
	//key = short version url
	//value = real URL
	path := map[string]string{
		"/youtube" : "https://youtube.com",
		"/github": "https://github.com/samuelsih", 
	}

	mapHelper := helper.MapHelper(path, mux)

	mux.HandleFunc("/", defaultResponse)

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, mapHelper)
}

func defaultResponse(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is default response for http://localhost%s", port)
}