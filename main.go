package main;

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Item struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

var items []Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item 
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	newItem.ID = fmt.Sprintf("%d", len(items) + 1)
	newItem.Name = fmt.Sprintf("Item %d", len(items) + 1)

	items = append(items, newItem)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Handle preflight request
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.WriteHeader(http.StatusOK)
		return
	}
	params := mux.Vars(r)
	itemID := params["id"]

	index := -1

	// Find the index of the item with the specified ID
	for i, item := range items {
		if item.ID == itemID {
			index = i
			break;
		}
	}

	// If the item is found, remove it from the items slice
	if index != -1 {
		items = append(items[:index], items[index+1:]...)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	router := mux.NewRouter()

	items = append(items, Item{ID: "1", Name: "Item 1", Description: "Description 1"})
	items = append(items, Item{ID: "2", Name: "Item 2", Description: "Description 2"})

    router.HandleFunc("/api/items", GetItems).Methods("GET")
	router.HandleFunc("/api/items", CreateItem).Methods("POST")
    router.HandleFunc("/api/items/{id}", DeleteItem).Methods("DELETE", "OPTIONS")

	handler := cors.Default().Handler(router)

	http.ListenAndServe(":8000", handler)
}