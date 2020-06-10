package taskhandler

import (
	"fmt"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This will create a task")
}

func UpdateTask(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This will update a task")
}

func RemoveTask(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This will remove a task")
}

func GetTasks(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This will return all your tasks")
}
