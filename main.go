package main

import (
	"fmt"
	"net/http"
)

// global vars
var taskListOne = []string{"clean desk", "clean bed", "clean shower", "clean dishes"}
var taskListTwo = []string{"fold clothes", "fold shirts", "fold shorts", "fold socks"}
var taskListThree = []string{"dry clothes", "dry shirts", "dry shorts", "dry socks"}
var taskLists = [][]string{taskListOne, taskListTwo, taskListThree}

func main() {
	print("Welcome to the to do list app!")

	//handler
	//(url, function tp handle)
	http.HandleFunc("/", welcome)
	http.HandleFunc("/greeting", helloUser)
	http.HandleFunc("/showTasksPage", showTasksPage)

	http.ListenAndServe(":8080", nil)
}

// TODO send user to greeting/showtasks? also explanation that it won't store anything once you close instance
func welcome(writer http.ResponseWriter, request *http.Request) {
	var welcome_msg = `Hi! This is the home page. This project was built entirely 
	using VSCode and Golang. Please visit /greeting or /showTasks to view the other pages.`
	fmt.Fprint(writer, welcome_msg)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = `Hello user! This is a page specifically meant for greeting the user.
	You can think of it like the screen you would see after logging in? However, it's purpose has 
	been reduced to nearly nothing without proper frontend and api implementation, so 
	now it's kind of just a random page :3`
	fmt.Fprint(writer, greeting)
}

func showTasksPage(writer http.ResponseWriter, request *http.Request) {
	//tells client to interpret response as html
	writer.Header().Set("Content-Type", "text/html")

	var task_msg = `<p>This is the tasks page. Click Button 1 to display Task List 1. Click Button 2 to display Task List 2.</p>`
	fmt.Fprintln(writer, task_msg)

	//button press submits a form, name is used when submitted -> when called, sends value
	//TODO make these variable calls
	fmt.Fprintln(writer, `
		<form method="POST">
			<button type="submit" name="action" value="the Reset button">Reset</button>
			<button type="submit" name="action" value="Button 1">Button 1</button>
			<button type="submit" name="action" value="Button 2">Button 2</button>
			<button type="submit" name="action" value="the Show All button">Show all</button>
		</form>
	`)

	//checks "did client click a button and submit a form (POST request)"
	if request.Method == http.MethodPost {
		//get button value
		buttonPressed := request.FormValue("action")
		fmt.Fprintf(writer, "<h3>You clicked %s.\n</h3>", buttonPressed)

		if buttonPressed == "the Reset button" { //reset button
			return
		} else if buttonPressed == "the Show All button" { // all task list
			for index, taskList := range taskLists {
				fmt.Print(taskList, " run ", index)
				printTasks(taskList, writer)
			}
		} else { //specific task list
			//displaying which list num
			//could alternatively use stringBuilder
			buttonNum := string(buttonPressed[7]) //for simplicity
			fmt.Fprintf(writer, "<h1>Displaying Task List %s:\n</h1>", buttonNum)

			//picking which task list to display based on button click
			var taskList []string
			switch buttonNum {
			case "1":
				taskList = taskListOne
			case "2":
				taskList = taskListTwo
			case "3":
				taskList = taskListThree
			}

			//print list items
			printTasks(taskList, writer)
		}

	}
}

func printTasks(taskList []string, writer http.ResponseWriter) {
	for i := 0; i <= len(taskList)-1; i++ {
		fmt.Fprintf(writer, "<li>%s</li>", taskList[i])
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

func printAllTasks(listOfLists [][]string) {
	for _, taskList := range listOfLists {
		for _, tasks := range taskList {
			println("--- ", tasks)
		}
	}
	for i := 0; i <= len(listOfLists)-1; i++ {
		for j := 0; j <= len(listOfLists[i])-1; j++ {

			println(">>>", listOfLists[i][j])
		}
	}
}

func addTask(taskList []string, task string) []string {
	taskList = append(taskList, task)
	return taskList
}

func test() {
	fmt.Println("test")

	fmt.Println("Welcome to my to-do list app!")

	taskListOne := []int{0}
	for i := 11; i <= 14; i++ {
		taskListOne = append(taskListOne, i)
	}

	for index, task := range taskListOne {
		println(index, task)
	}
	for _, task := range taskListOne {
		println(task)
	}

	fmt.Println("List of to dos")
	fmt.Println(taskListOne)
	fmt.Println(taskListOne[0])
	fmt.Println(taskListOne[1:])
	fmt.Println(taskListOne[1:2])
}
