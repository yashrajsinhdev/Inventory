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

func getProducts(w http.ResponseWriter, r *http.Request) {
	result := products 

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

	// RESTful API endpoints
	r.HandleFunc("/products", getProducts).Methods("GET")   // List or filter products
	r.HandleFunc("/product", addProduct).Methods("POST")	// Add a new product
	r.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE") 	// Delete a product by ID
	r.HandleFunc("/product/{id}", updateProduct).Methods("PUT")	// Update a product by ID
	
	log.Println("Server starting on :8080...")
    // Start the HTTP server on port 8080 and use the mux router to handle requests.
    log.Fatal(http.ListenAndServe(":8080", r))
}