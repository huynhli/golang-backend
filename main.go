package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	print("Welcome to the to do list app!")

	//handler -> (url, function to handle)
	http.HandleFunc("/", welcomePage)
	http.HandleFunc("/loggedIn", loggedInPage)
	http.HandleFunc("/showTasksPage", showTasksPage)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		print(err.Error())
	}
}
