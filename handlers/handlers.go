package handlers

import(
  "fmt"
  "net/http"
  "../models"
  "github.com/gorilla/mux"
  "strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Se obtiene todos los usuarios!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  userId, _ := strconv.Atoi(vars["id"])

  if user, err := models.GetUser(userId); err != nil {
    models.SendNotFound(w)
  } else  {
    models.SendData(w, user)
  }

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Se crea usuario!")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Se actualiza usuario!")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Se eliminó usuario!")
}
