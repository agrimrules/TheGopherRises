package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "log"
)

type Health struct {
  Healthy bool `json:"Healthy"`
  Status int `json:"Status"`
}

func HealthCheck(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(Health{Healthy: true,Status: 200})
}

func main(){
  router := mux.NewRouter()
  router.HandleFunc("/health",HealthCheck).Methods("GET")
  log.Fatal(http.ListenAndServe(":8000",router))
}
