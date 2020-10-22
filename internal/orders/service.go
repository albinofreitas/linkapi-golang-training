package orders

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ResponseError handle http error
type ResponseError struct {
	Message string `json:"message"`
}

// Store handle incoming orders
func Store(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseError{Message: "failed to parse body"})
		return
	}

	order.ID = primitive.NewObjectID()

	go func() {
		if err := order.store(); err != nil {
			log.Print(err)
		}
	}()

	json.NewEncoder(w).Encode(order)
}

// Show return an order
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var order Order

	err := order.getByID(vars["id"])

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseError{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(order)
}
