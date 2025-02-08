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

//http is how data from is transferred between backend and frontend
//http request from client, http response from server
//security -> https -> encryption and verification

//backend
//serves on opening
//mediates frontend and database as well as complex functions

//url?
//protocal + hostname + port

//port 80 = default e.g. google.com:80

//index.html
//scripts or images to use -> not localhost

//host names mapped to ip addresses
