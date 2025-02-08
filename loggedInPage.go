package main

import (
	"fmt"
	"net/http"
)

func loggedInPage(writer http.ResponseWriter, request *http.Request) {
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
