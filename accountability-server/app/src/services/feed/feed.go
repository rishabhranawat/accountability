package feed

import (
	"encoding/json"
	"fmt"
	"net/http"

	authmiddleware "../../middleware"

	"../../env"
	"../../models"
)

type FeedResponse struct {
	Tasks   []models.Task
	Message string
}

// todo: should get make a db call to fetch all the tasks and
// reverse sort them
func GetFeed(w http.ResponseWriter, r *http.Request) {

	var user = authmiddleware.GetCurrentUser(r)

	if user.ID == 0 {
		http.Error(w, "Unable to find the current user", http.StatusBadRequest)
		fmt.Fprintln(w, "Failed to retrieve user's feed")
		return
	}

	var tasks []models.Task
	env.DbConnection.Where("id IN (SELECT b.task_refer_id FROM trackers AS b WHERE b.user_refer_id = ?) OR user_id = ?", user.ID, user.ID).Find(&tasks)

	var response FeedResponse
	response.Tasks = tasks
	response.Message = "Found tasks"

	jResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)

}
