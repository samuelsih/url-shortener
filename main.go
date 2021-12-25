package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/samuelsih/url-shortener/helper"
)

const port = ":8080"

func main() {
	mux := mux.NewRouter()

	//load .env file
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	//path is map of shortened URL and real URL
	//key = short version url
	//value = real URL
	path := map[string]string{
		"/youtube" : os.Getenv("YOUTUBE_URL"),
		"/github": os.Getenv("GITHUB_URL"), 
	}

	//mapHelper is variable for http.Handler comes from helper.go
	mapHelper := helper.MapHelper(path, mux)

	mux.HandleFunc("/", defaultResponse)

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, mapHelper)
}

func defaultResponse(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is default response for http://localhost%s", port)
}