package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Ticket struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

var tickets []Ticket

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tickets", getTickets).Methods("GET")
	r.HandleFunc("/tickets/{id}", getTicket).Methods("GET")
	r.HandleFunc("/tickets", createTicket).Methods("POST")
	r.HandleFunc("/tickets/{id}", updateTicket).Methods("PUT")
	r.HandleFunc("/tickets/{id}", deleteTicket).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getTickets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tickets)
}

func getTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range tickets {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Ticket not found", http.StatusNotFound)
}

func createTicket(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	_ = json.NewDecoder(r.Body).Decode(&ticket)
	ticket.ID = uuid.New().String()
	ticket.CreatedAt = time.Now()
	tickets = append(tickets, ticket)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

func updateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range tickets {
		if item.ID == params["id"] {
			var ticket Ticket
			_ = json.NewDecoder(r.Body).Decode(&ticket)
			ticket.ID = params["id"]
			tickets[i] = ticket
			json.NewEncoder(w).Encode(ticket)
			return
		}
	}

	http.Error(w, "Ticket not found", http.StatusNotFound)
}

func deleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range tickets {
		if item.ID == params["id"] {
			tickets = append(tickets[:i], tickets[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Ticket not found", http.StatusNotFound)
}
