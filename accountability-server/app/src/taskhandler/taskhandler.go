package taskhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../env"
	"../models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to create task. Sample Input: { Name: X, Description: Y, Workers: User, Trackers: List<Users>, Milestones: List<TaskMileston> }")
		return
	}
	env.DbConnection.Create(&task)
	fmt.Fprintln(w, "Successfully created task")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	env.DbConnection.Where("Id = ?", updatedTask.Id).Find(&task)

	task.Name = updatedTask.Name
	task.Description = updatedTask.Description
	task.Trackers = updatedTask.Trackers
	task.Milestones = updatedTask.Milestones
	task.Workers = updatedTask.Workers

	env.DbConnection.Save(task)

	fmt.Fprintln(w, "Successfully updated task")
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	env.DbConnection.Where("Id = ?", updatedTask.Id).Find(&task)

	env.DbConnection.Delete(&task)
	fmt.Fprintln(w, "This will remove a task")
}

func GetUserTasks(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	name := queryValues.Get("username")
	fmt.Fprintln(w, "This is the user "+name)

	//TODO guruis - query tasks here for the given user
	// var task models.Task
	// env.DbConnection.Where("Workers.UserName = ?", name).Find(&task)
}
