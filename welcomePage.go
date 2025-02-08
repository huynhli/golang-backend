package main

import (
	"fmt"
	"net/http"
)

// TODO also explanation that it won't store anything once you close instance
func welcomePage(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(writer, `
		Hi! This is the home page. This project was built entirely 
		using VSCode and Golang. Please visit /greeting or /showTasks to view the other pages.
		`)
	fmt.Fprintln(writer, `
		<form method="GET" action="/showTasksPage">
			<button type="submit">Go to tasks page</button>
		</form>
		<form method="GET" action="/loggedIn">
			<button type="submit">Login Page</button>
		</form>`)

}
