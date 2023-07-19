package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Server struct to represent a server
type Server struct {
	Name string `json:"name"`
}

var servers = []Server{
	{Name: "tarikdmzmachine1"},
	{Name: "server2"},
	{Name: "server6"},
	{Name: "server3"},
}

func main() {
	http.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(servers)
	})

	http.HandleFunc("/ansible_access", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ansible_permission.html")
	})

	log.Println("Server is running on http://localhost:5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
