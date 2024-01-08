package handlers

import (
	"encoding/json"
	"myapp/datastore"
	"myapp/models"
	"net/http"
)

// CreateEntry - Handler for creating a new entry
func CreateEntry(w http.ResponseWriter, r *http.Request) {
	var entry models.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := datastore.CreateEntry(entry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}

// GetAllEntries - Handler to retrieve all entries
func GetAllEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := datastore.GetAllEntries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entries)
}
