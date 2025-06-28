package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//each product has a fixed set of fields: ID, Name, and Quantity.
type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

//products is a slice (dynamic array) of Product structs.
var products []Product

//nextID is used to assign unique IDs to each product.
var nextID = 1

// CORS middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:2222")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	result := products 

	//Below block is for filter by name functionality
	nameFilter := r.URL.Query().Get("name")
	if nameFilter != "" {
		filtered := []Product{}
		for _, p := range result {
			if strings.Contains(strings.ToLower(p.Name), strings.ToLower(nameFilter)) {
				filtered = append(filtered, p)
			}
		}
		result = filtered
	}


    // Sort products by quantity if '?sort=quantity' query parameter is present same for name.
	sortBy := r.URL.Query().Get("sort")
	switch sortBy {
	case "quantity":
		sort.Slice(result, func(i, j int) bool {
			return result[i].Quantity < result[j].Quantity
		})
	case "name":
		sort.Slice(result, func(i, j int) bool {
			return strings.ToLower(result[i].Name) < strings.ToLower(result[j].Name)
		})
	}

    // Encode the result slice as JSON and write it to the response.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
func addProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	newProduct.ID = nextID
	nextID++
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	
	http.Error(w, "Product not found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updated Product
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.ID == id {
			products[i].Name = updated.Name
			products[i].Quantity = updated.Quantity
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

//In main we will setup HTTP server,initialize sample data, and define API endpoints.
func main(){
	r := mux.NewRouter()

    // Add sample products to the in-memory slice.
    // Using append allows us to grow the slice dynamically.
    products = append(products, Product{ID: nextID, Name: "iPhone 16", Quantity: 10})
    nextID++
    products = append(products, Product{ID: nextID, Name: "Samsung Galaxy S25", Quantity: 5})
    nextID++

	// Apply CORS middleware to all endpoints
	r.HandleFunc("/products", enableCORS(getProducts)).Methods("GET")
	r.HandleFunc("/product", enableCORS(addProduct)).Methods("POST")
	r.HandleFunc("/product/{id}", enableCORS(deleteProduct)).Methods("DELETE")
	r.HandleFunc("/product/{id}", enableCORS(updateProduct)).Methods("PUT")
	
	// Add OPTIONS handler for preflight requests
	r.HandleFunc("/products", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/product", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/product/{id}", handleOptions).Methods("OPTIONS")

	log.Println("Server starting on :1111")
    // Start the HTTP server on port 1111 and use the mux router to handle requests.
    log.Fatal(http.ListenAndServe(":1111", r))
}
// Handle OPTIONS requests
func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:2222")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}