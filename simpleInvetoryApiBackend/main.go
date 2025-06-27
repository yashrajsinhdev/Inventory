package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
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

	    // Filter products by name if 'name' query parameter is present.
    // Uses strings.Contains for partial, case-insensitive match.
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
	if sortBy == "quantity" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Quantity < result[j].Quantity
		})
	} else if sortBy == "name" {
		sort.Slice(result, func(i, j int) bool {
			return strings.ToLower(result[i].Name) < strings.ToLower(result[j].Name)
		})
	}


    // Encode the result slice as JSON and write it to the response.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
//In main we will setup HTTP server,initialize sample data, and define API endpoints.
func main(){
	r := mux.NewRouter()

    // Add sample products to the in-memory slice.
    // Using append allows us to grow the slice dynamically.
    products = append(products, Product{ID: nextID, Name: "Apple", Quantity: 10})
    nextID++
    products = append(products, Product{ID: nextID, Name: "Banana", Quantity: 5})
    nextID++

	// RESTful API endpoints
	r.HandleFunc("/products", getProducts).Methods("GET")       // List or filter products
	
	log.Println("Server starting on :8080...")
    // Start the HTTP server on port 8080 and use the mux router to handle requests.
    log.Fatal(http.ListenAndServe(":8080", r))
}