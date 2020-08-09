package reminder

import (
	"fmt"
	"net/http"

	authmiddleware "../../middleware"

	"../../models"
)

func SendReminder(w http.ResponseWriter, r *http.Request) {

	var user = authmiddleware.GetCurrentUser(r)

	if user.ID == 0 {
		http.Error(w, "Unable to find the current user", http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to retrieve")
		return
	}

	var reminder models.Reminder
	env.DbConnection.Create(&reminder)
}
