package handlers

import(
  "net/http"
  "../models"
  "github.com/gorilla/mux"
  "strconv"
  "encoding/json"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
  models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {

  if user, err := getUserByRequest(r); err != nil {
    models.SendNotFound(w)
  } else  {
    models.SendData(w, models.SaveUser(user))
  }

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
  user := models.User{}
  decoder := json.NewDecoder(r.Body)

  if err := decoder.Decode(&user); err != nil {
    models.SendUnprocessableEntity(w)
  } else {
    models.SaveUser(user)
    models.SendData(w, user)
  }
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  user, err := getUserByRequest(r)

  if err != nil {
    models.SendNotFound(w)
    return
  }

  userResponse := models.User{}
  decoder := json.NewDecoder(r.Body)

  if err := decoder.Decode(&userResponse); err != nil {
    models.SendUnprocessableEntity(w)
    return
  }

  user = models.UpdateUser(user, userResponse.Username, userResponse.Password)
  models.SendData(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  if user, err := getUserByRequest(r); err != nil {
    models.SendNotFound(w)
  } else {
    models.DeleteUser(user.Id)
    models.SendNoContent(w)
  }
}

func getUserByRequest(r *http.Request) (models.User, error) {
  vars := mux.Vars(r)
  userId, _ := strconv.Atoi(vars["id"])

  if user, err := models.GetUser(userId); err != nil {
    return user, err
  } else {
    return user, nil
  }
}
