package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// global vars
// TODO creation and deletion of custom task lists
// TODO sorting and filtering tasks
var taskListOne []string
var taskListTwo []string
var taskListThree []string
var taskLists = [][]string{taskListOne, taskListTwo, taskListThree}
var currentTaskListInt = 1

func showTasksPage(writer http.ResponseWriter, request *http.Request) {

	//tells client to interpret response as html
	writer.Header().Set("Content-Type", "text/html")

	fmt.Fprintln(writer, `
		<p>This is the tasks page. Note: Each task list has a maximum size of 10 tasks. Each task must be under 30 characters long.<br /> 
		To add a task, simply click on the button for the respective task list, then add one or more tasks, finally clicking on the <br />
		task list button again.</p>`)

	//button press submits a form, name is used when submitted -> when called, sends value
	fmt.Fprintln(writer, `
		<form method="GET">
			<input type="hidden" name="formType" value="display">
			<button type="submit" name="action" value="the Reset button">Reset</button>
			<button type="submit" name="action" value="1">Task List 1</button>
			<button type="submit" name="action" value="2">Task List 2</button>
			<button type="submit" name="action" value="3">Task List 3</button>
			<button type="submit" name="action" value="the Show All button">Show all</button>
		</form>
	`)

	fmt.Fprintln(writer, `
		<form method="POST">
			<label>Task to add: </label>
			<input type="text" id="inputBox" name="user_input">
			<label>Priority Level: </label>
			<select name="priority" id="priority">
				<option value="1">1</option>
				<option value="2">2</option>
				<option value="3">3</option>
				<option value="4">4</option>
				<option value="5">5</option>
				<option value="6">6</option>
				<option value="7">7</option>
				<option value="8">8</option>
				<option value="9">9</option>
				<option value="10">10</option>
			</select>
			<button type="submit" name="action" value="add_task">Add task</button>
		</form>
	`)

	//GET only retrieves data -> get all tasks //
	//POST modifies data/creates new resource -> adding task //
	//priority levels for tasks //
	//PUT modifies data/updates resources -> editing tasks //
	//DELETE remove resources -> deleting trasks //

	//checks "did client click a button and submit a form (POST request)"
	if request.Method == http.MethodGet && request.URL.Query().Get("action") != "" {
		buttonPressed := request.FormValue("action")
		if buttonPressed == "the Reset button" { //reset button
			return
		} else if buttonPressed == "the Show All button" { // all task list
			for index, taskList := range taskLists {
				fmt.Fprintf(writer, `<h2>List %d:</h2>`, index+1)
				printTasksOfList(taskList, writer)
			}
		} else { //specific task list
			buttonNum := string(buttonPressed[0])
			fmt.Fprintf(writer, "<h1>Displaying Task List %s:\n</h1>", buttonNum)
			var taskList []string
			taskList = taskListChooser(buttonNum, taskList) //picking task list
			currentTaskListInt, _ = strconv.Atoi(buttonNum)

			printTasksOfList(taskList, writer) //print list items

			//rename section
			renameDropdownString := renameDropdownHelper(taskList)
			fmt.Fprintf(writer, `
				<form method="POST">
					<label>Change task </label>
						<select name="priority" id="priority">
							%s
						</select>
					<label> to </label>
						<input type="text" id="inputBox" name="changed_task">
					<button type="submit" name="action" value="rename">Rename</button>
				</form>
			`, renameDropdownString.String())

			//deletion section
			dropdownString := dropdownDeleteHelper(taskList)
			fmt.Fprintf(writer, `
				<form method="POST">
					<label>Delete task </label>
						<select name="priority" id="priority">
							%s
						</select>
					<label>? </label>
					<button type="submit" name="action" value="delete">Delete</button>
				</form>
			`, dropdownString.String())
		}
	}

	if request.Method == http.MethodPost {
		if request.FormValue("action") == "add_task" {
			addTask(request, writer)
		} else if request.FormValue("action") == "rename" {
			renameTask(request)
		} else if request.FormValue("action") == "delete" {
			deleteTask(request)
		}
	}

}

func renameDropdownHelper(taskList []string) strings.Builder {
	var optionString strings.Builder
	for i := 1; i <= len(taskList); i++ {
		optionString.WriteString("<option value=\"task-")
		optionString.WriteString(strconv.Itoa(i))
		optionString.WriteString("\">")
		optionString.WriteString(strconv.Itoa(i))
		optionString.WriteString("</option>")
	}
	return optionString
}

func dropdownDeleteHelper(taskList []string) strings.Builder {
	var dropdownDeleteString strings.Builder
	for i := 1; i <= len(taskList); i++ {
		dropdownDeleteString.WriteString("<option value=\"task-")
		dropdownDeleteString.WriteString(strconv.Itoa(i))
		dropdownDeleteString.WriteString("\">")
		dropdownDeleteString.WriteString(strconv.Itoa(i))
		dropdownDeleteString.WriteString("</option>")
	}
	return dropdownDeleteString
}

func addTask(request *http.Request, writer http.ResponseWriter) {
	if !regexp.MustCompile(`^.{1,30}$`).MatchString(request.FormValue("user_input")) { //invalid input
		fmt.Fprintln(writer, "Not a valid input. Please retry.")
	} else if len(taskLists[currentTaskListInt-1]) >= 10 { //too many tasks already
		fmt.Fprintln(writer, "Task list it full. Complete a task and try again.")
	} else { // adding based on priority selected
		var currentTaskList = &taskLists[currentTaskListInt-1] //needs & and * for pointer stuff
		var priorityLevel, _ = strconv.Atoi(request.FormValue("priority"))
		if priorityLevel <= len(*currentTaskList) {
			*currentTaskList = append(*currentTaskList, "")
			copy((*currentTaskList)[priorityLevel:], (*currentTaskList)[priorityLevel-1:]) // moves everything to the right
			(*currentTaskList)[priorityLevel-1] = request.FormValue("user_input")
			fmt.Fprintf(writer, "<h2>Task added to list %d!</h2>", currentTaskListInt)
		} else {
			*currentTaskList = append(*currentTaskList, request.FormValue("user_input"))
			fmt.Fprintf(writer, "<h2>Task added to list %d!</h2>", currentTaskListInt)
		}
	}
}

func deleteTask(request *http.Request) {
	priorityString := request.FormValue("priority")
	var priorityLevelNum, _ = strconv.Atoi(priorityString[5:])
	var currentTaskList = &taskLists[currentTaskListInt-1]
	*currentTaskList = append((*currentTaskList)[:priorityLevelNum-1], (*currentTaskList)[priorityLevelNum:]...)
}

func renameTask(request *http.Request) {
	priorityString := request.FormValue("priority")
	var priorityLevelNum, _ = strconv.Atoi(priorityString[5:])
	var currentTaskList = &taskLists[currentTaskListInt-1]
	(*currentTaskList)[priorityLevelNum-1] = request.FormValue("changed_task")
}

func taskListChooser(buttonNum string, taskList []string) []string {
	switch buttonNum {
	case "1":
		taskList = taskLists[0]
	case "2":
		taskList = taskLists[1]
	case "3":
		taskList = taskLists[2]
	}
	return taskList
}

func printTasksOfList(taskList []string, writer http.ResponseWriter) {
	if len(taskList) == 0 {
		fmt.Fprintln(writer, "This list is empty.")
	}
	fmt.Fprintln(writer, "<ol>")
	for i := 0; i <= len(taskList)-1; i++ {
		fmt.Fprintf(writer, "<li>%s</li>", taskList[i])
	}
	fmt.Fprintln(writer, "</ol>")
}
