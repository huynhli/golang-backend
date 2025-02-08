package main

import (
	"fmt"
	"net/http"
)

func main() {
	print("Welcome to the to do list app!")

	//handler -> (url, function tp handle)
	http.HandleFunc("/", welcome)
	http.HandleFunc("/loggedIn", helloUser)
	http.HandleFunc("/showTasksPage", showTasksPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		print(err.Error())
	}
}

// TODO routing? also explanation that it won't store anything once you close instance
func welcome(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "text/html")
	var welcome_msg = `Hi! This is the home page. This project was built entirely 
	using VSCode and Golang. Please visit /greeting or /showTasks to view the other pages.`
	fmt.Fprintln(writer, welcome_msg)
	fmt.Fprintln(writer, `
		<form method="GET" action="/showTasksPage">
			<button type="submit">Go to tasks page</button>
		</form>
		<form method="GET" action="/loggedIn">
			<button type="submit">Login Page</button>
		</form>`)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	var greeting = `Hello user! This is a page specifically meant for greeting the user.
	You can think of it like the screen you would see after logging in? However, it's purpose has 
	been reduced to nearly nothing without proper frontend and api implementation, so 
	now it's kind of just a random page :3`
	fmt.Fprintln(writer, greeting)
	fmt.Fprintln(writer, `
		<form method="GET" action="/showTasksPage">
			<button type="submit">Go to tasks page</button>
		</form>`)

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
