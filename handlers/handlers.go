package handlers

import(
  "fmt"
  "net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Se obtiene todos los usuarios!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Se obtiene todos los usuarios!")
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
