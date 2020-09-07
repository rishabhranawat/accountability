package taskhandler

import (
  authmiddleware "../middleware"
  storage "../middleware/storage"
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"

  "../env"
  "../models"
  "github.com/gorilla/mux"
)

type AcknowledgmentResponse struct {
	Message string
}

type CreateTaskRequestBody struct {
	UserTask      models.Task
	TrackerEmails []string
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	// TODO: move this to a request scope?
	var user = authmiddleware.GetCurrentUser(r)

	if user.ID == 0 {
		http.Error(w, "Unable to find the current usr", http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to create task.")
	}

	var task CreateTaskRequestBody
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to create task.")
		return
	}
	task.UserTask.UserID = user.ID
	env.DbConnection.Create(&task.UserTask)

	var trackerUsers []models.User
	env.DbConnection.Where("Email IN (?)", task.TrackerEmails).Find(&trackerUsers)

	// figure out a way to insert many
	for _, item := range trackerUsers {
		var t models.Tracker
		t.UserReferID = item.ID
		t.TaskReferID = task.UserTask.ID
		env.DbConnection.Create(&t)
	}

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

func PostTaskUpdate(w http.ResponseWriter, r *http.Request) {
	var taskUpdate models.TaskUpdate
	taskReferId, err := strconv.Atoi(r.FormValue("TaskReferID"))
	if err != nil {
	  http.Error(w, err.Error(), http.StatusBadRequest)
	  return
  }
  taskUpdate.TaskReferID = taskReferId
	taskUpdate.Description = r.FormValue("Description")

	file, header, errWhileRetrievingFile := r.FormFile("uploadFile")

	if file != nil {
    if errWhileRetrievingFile != nil {
      http.Error(w, errWhileRetrievingFile.Error(), http.StatusBadRequest)
      return
    }
    defer file.Close()

    if file == nil {
      env.DbConnection.Create(&taskUpdate)
      return
    }

    fileKey := storage.GetUniqueS3Key(header.Filename)
    uploadedFileSuccessFully := storage.UploadFileToS3(file, fileKey)
    if !uploadedFileSuccessFully {
      http.Error(w, "There was an error uploading your file to S3", http.StatusBadRequest)
      return
    }

    taskUpdate.MediaURL = fileKey
  }

  env.DbConnection.Create(&taskUpdate)

  var updates []models.TaskUpdate
  env.DbConnection.Where("task_refer_id = ?", taskReferId).Find(&updates)

  jResponse, err := json.Marshal(updates)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jResponse)

}

func PostTaskComment(w http.ResponseWriter, r *http.Request) {
	var taskComment models.TaskComment

	var user = authmiddleware.GetCurrentUser(r)

	if user.ID == 0 {
		http.Error(w, "User Not Found", http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&taskComment)
	taskComment.UserReferID = user.ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	env.DbConnection.Create(&taskComment)

	var comments []models.TaskComment
	env.DbConnection.Where("task_refer_id = ?", taskComment.TaskReferID).Find(&comments)

	jResponse, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)

}

func FetchTaskComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var comments []models.TaskComment
	env.DbConnection.Where("task_refer_id = ?", vars["task-id"]).Find(&comments)

	jResponse, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}

func FetchTaskUpdates(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  var updates []models.TaskUpdate
  env.DbConnection.Where("task_refer_id = ?", vars["task-id"]).Find(&updates)

  jResponse, err := json.Marshal(updates)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }



  w.Header().Set("Content-Type", "application/json")
  w.Write(jResponse)
}

func FetchTaskDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var task models.Task
	env.DbConnection.Where("id = ?", vars["task-id"]).Find(&task)

	jResponse, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}
