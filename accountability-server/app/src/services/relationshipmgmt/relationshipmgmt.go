package relationshipmgmt

import (
  "../../env"
  authmiddleware "../../middleware"
  "../../models"
  "encoding/json"
  "net/http"
)

type CreateRelationshipRequest struct {
  RelationshipToUserName string
  RelationshipToUserEmail string
}

type ApproveUserRequest struct {
  RelationshipID uint
}

type DeleteRelationshipRequest struct {
  RelationshipID uint
}

func CreateRelationship(w http.ResponseWriter, r *http.Request) {

  var fromUser = authmiddleware.GetCurrentUser(r)

  if fromUser.ID == 0 {
    http.Error(w, "Unable to find the current user", http.StatusBadRequest)
    return
  }

  var request CreateRelationshipRequest
  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  var toUser models.User
  env.DbConnection.Where("user_name = ? OR  email = ?",
    request.RelationshipToUserName, request.RelationshipToUserEmail).Find(&toUser)

  var newRelationship models.Relationship
  newRelationship.Approved = false
  newRelationship.RelationshipFromID = fromUser.ID
  newRelationship.RelationshipToID = toUser.ID
  env.DbConnection.Create(&newRelationship)

  return
}

func ApproveRelationship(w http.ResponseWriter, r *http.Request) {
  var request ApproveUserRequest
  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  var relationship models.Relationship
  env.DbConnection.Model(&relationship).Where("id = ?", request.RelationshipID).Update("approved", true)

  return
}

func DeleteRelationship(w http.ResponseWriter, r *http.Request) {
  var request DeleteRelationshipRequest
  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  var relationship models.Relationship
  env.DbConnection.Where("id = ?", request.RelationshipID).Delete(&relationship)
}
