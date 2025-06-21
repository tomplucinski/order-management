package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type OrderRequest struct {
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go processOrder(req)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Order received"))
}

func OrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "Missing order ID", http.StatusBadRequest)
		return
	}

	orderID := parts[2]
	w.Write([]byte("Fetching order: " + orderID))
}

func processOrder(req OrderRequest) {
	log.Printf("Processing order: %+v\n", req)
}
