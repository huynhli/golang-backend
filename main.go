package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

// global vars
// TODO creating lists -> limit to max # of lists and tasks
var taskListOne = []string{"clean desk", "clean bed", "clean shower", "clean dishes"}
var taskListTwo = []string{"fold clothes", "fold shirts", "fold shorts", "fold socks"}
var taskListThree = []string{"dry clothes", "dry shirts", "dry shorts", "dry socks"}
var taskLists = [][]string{taskListOne, taskListTwo, taskListThree}
var currentTaskList int

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
			<input type="hidden" name="formType" value="display">
			<button type="submit" name="action" value="the Reset button">Reset</button>
			<button type="submit" name="action" value="Button 1">Task List 1</button>
			<button type="submit" name="action" value="Button 2">Task List 2</button>
			<button type="submit" name="action" value="the Show All button">Show all</button>
		</form>
	`)

	fmt.Fprintln(writer, `
		<form method="POST">
			<input type="hidden" name="formType" value="taskAdd">
			<p>Task to add: </p>
			<input type="text" id="inputBox" name="user_input">
			<button type="submit" name="action" value="Add task">Add task</button>
		</form>
	`)
	//GET only retrieves data -> get all tasks //
	//POST modifies data/creates new resource -> adding task //
	//PUT modifies data/updates resources -> editing tasks
	//DELETE remove resources -> deleting trasks

	//checks "did client click a button and submit a form (POST request)"
	if request.Method == http.MethodPost {

		//TODO delete
		buttonPressed := request.FormValue("action")
		fmt.Fprintf(writer, "<h3>You clicked %s.\n</h3>", buttonPressed)

		// checks if adding task
		if request.FormValue("formType") == "taskAdd" {
			if !regexp.MustCompile(`^.{1,30}$`).MatchString(request.FormValue("user_input")) {
				fmt.Fprintln(writer, "Not a valid input. Please only add tasks under 30 characters in length. Please retry.")
			} else {
				taskLists[currentTaskList] = append(taskLists[currentTaskList], request.FormValue("user_input"))
				taskListOne = taskLists[currentTaskList]
				fmt.Fprintf(writer, "Task added to list %d!", currentTaskList)
			}
		} else {

			if buttonPressed == "the Reset button" { //reset button
				return
			} else if buttonPressed == "the Show All button" { // all task list
				for _, taskList := range taskLists {
					printTasks(taskList, writer)
				}
			} else { //specific task list

				//displaying which list num
				//could alternatively use stringBuilder
				buttonNum := string(buttonPressed[7]) //for simplicity
				fmt.Fprintf(writer, "<h1>Displaying Task List %s:\n</h1>", buttonNum)

				//picking which task list to display based on button click
				var taskList []string
				taskList = taskListSwitch(buttonNum, taskList)
				currentTaskList, _ = strconv.Atoi(buttonNum)

				//print list items
				printTasks(taskList, writer)
			}
		}
	}
}

func taskListSwitch(buttonNum string, taskList []string) []string {
	switch buttonNum {
	case "1":
		taskList = taskListOne
	case "2":
		taskList = taskListTwo
	case "3":
		taskList = taskListThree
	}
	return taskList
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
