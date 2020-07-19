package taskhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"

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

	var id = guuid.New()

	task.TaskId = id.String()
	env.DbConnection.Create(&task)

	var response AcknowledgmentResponse
	response.Message = "Successfully created task. Id: " + task.TaskId

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
	env.DbConnection.Where("Id = ?", updatedTask.TaskId).Find(&task)

	task.Name = updatedTask.Name
	task.Description = updatedTask.Description
	task.Trackers = updatedTask.Trackers
	task.Milestones = updatedTask.Milestones
	task.Workers = updatedTask.Workers

	env.DbConnection.Save(task)

	var response AcknowledgmentResponse
	response.Message = "Successfully updated task. Id: " + task.TaskId

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
	env.DbConnection.Where("Id = ?", updatedTask.TaskId).Find(&task)

	env.DbConnection.Delete(&task)

	var response AcknowledgmentResponse
	response.Message = "Successfully removed task. Id: " + task.TaskId

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
	env.DbConnection.Where("Workers = ?", user).Find(&tasks)

	fmt.Fprintln(w, "This will retrieve all tasks for a given user")

	jResponse, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}
