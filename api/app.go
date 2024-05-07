package api

import (
	"log"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Msg struct {
	message string
}

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Setup Database, Router, and Routes
func (a *App) Initialize() {

	var o Orm
	o.SetupDb()
	a.DB = o.DB

	// connectionString := GetDsn()

	// var err error
	// a.DB, err = sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

// Initialize Routes
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.sayHello).Methods("GET")
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

// Run the App on the provided address
func (a *App) Run(addr string) {
	log.Printf("Listening on %v", addr)

	res := http.ListenAndServe(addr, a.Router)

	log.Fatal(res)
}

// Simple healthcheck style route
func (a *App) sayHello(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Hi Gary..."})
}

// Utility function to respond with error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Utility function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Get a list of products
func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	product := &Product{}
	products := product.GetProducts(a.DB, start, count)

	respondWithJSON(w, http.StatusOK, products)
}

// Get a product
func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := Product{ID: id}
	_, err = p.GetProduct(a.DB, id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Product ID not found")
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

// Create a product
func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p.CreateProduct(a.DB)

	respondWithJSON(w, http.StatusCreated, p)
}

// Update a Product
func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	p.UpdateProduct(a.DB, &p)

	respondWithJSON(w, http.StatusOK, p)
}

// Delete a Product
func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := Product{ID: id}
	p.DeleteProduct(a.DB)

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
