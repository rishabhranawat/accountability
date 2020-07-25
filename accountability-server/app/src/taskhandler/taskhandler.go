package taskhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../env"
	"../models"
)

type AcknowledgmentResponse struct {
	Message string
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to create task. Sample Input: { Name: X, Description: Y, Workers: User, Trackers: List<Users>, Milestones: List<TaskMileston> }")
		return
	}

	env.DbConnection.Create(&task)

	var response AcknowledgmentResponse
	response.Message = "Successfully created task."

	jResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	env.DbConnection.Where("Id = ?", updatedTask.ID).Find(&task)

	task.Name = updatedTask.Name
	task.Description = updatedTask.Description
	task.User = updatedTask.User

	env.DbConnection.Save(task)

	var response AcknowledgmentResponse
	response.Message = "Successfully updated task. Id: " + fmt.Sprint(task.ID)

	jResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task models.Task
	env.DbConnection.Where("Id = ?", updatedTask.ID).Find(&task)

	env.DbConnection.Delete(&task)

	var response AcknowledgmentResponse
	response.Message = "Successfully removed task. Id: " + fmt.Sprint(task.ID)
	jResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}

func FetchUserTasks(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tasks []models.Task
	env.DbConnection.Where("user_id = 1").Find(&tasks)

	fmt.Fprintln(w, "This will retrieve all tasks for a given user")

	jResponse, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}
