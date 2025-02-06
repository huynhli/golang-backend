package main

import (
	"fmt"
	"net/http"
)

func main() {
	print("Welcome to the to do list app!")

	//handler
	//(url, function tp handle)
	http.HandleFunc("/greeting", helloUser)

	http.ListenAndServe(":8080", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user!"
	fmt.Fprintf(writer, greeting)
}

func twoList() {
	taskListOne := []string{"clean desk", "clean bed", "clean shower", "clean dishes"}
	taskListTwo := []string{"fold clothes", "fold shirts", "fold shorts", "fold socks"}
	var taskLists [][]string
	taskLists = append(taskLists, taskListOne)

	taskListTwo = addTask(taskListTwo, "make dinner")
	taskLists = append(taskLists, taskListTwo)

	printAllTasks(taskLists)
}

//http is how data from is transferred between backend and frontend
//http request from client, http response from server
//security -> https -> encryption and verification

//backend
//serves on opening
//mediates frontend and database as well as complex functions

//url?
//protocl + hostname + port

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
