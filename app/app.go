package app

import (
	"fmt"
	"log"
	"net/http"

	"go-crud/app/handler"
	"go-crud/app/model"
	"go-crud/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {

	// Routing for handling the categories
	a.Get("/categories", a.GetAllCategories)
	a.Post("/categories", a.CreateCategory)
	a.Get("/categories/{name}", a.GetCategory)
	a.Put("/categories/{name}", a.UpdateCategory)
	a.Delete("/categories/{name}", a.DeleteCategory)

	// Routing for handling the product
	a.Get("/categories/{name}/products", a.GetAllProducts)
	a.Post("/categories/{name}/products", a.CreateProduct)
	a.Get("/categories/{name}/products/{id:[0-9]+}", a.GetProduct)
	a.Put("/categories/{name}/products/{id:[0-9]+}", a.UpdateProduct)
	a.Delete("/categories/{name}/products/{id:[0-9]+}", a.DeleteProduct)

	// Routing for handling the image
	a.Get("/categories/{name}/products/{name}/images", a.GetAllImages)
	a.Post("/categories/{name}/products/{name}/images", a.CreateImage)
	a.Get("/categories/{name}/products/{name}/images/{id:[0-9]+}", a.GetImage)
	a.Put("/categories/{name}/products/{name}/images/{id:[0-9]+}", a.UpdateImage)
	a.Delete("/categories/{name}/products/{name}/images{id:[0-9]+}", a.DeleteImage)

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Categories Handlers

func (a *App) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	handler.GetAllCategories(a.DB, w, r)
}

func (a *App) CreateCategory(w http.ResponseWriter, r *http.Request) {
	handler.CreateCategory(a.DB, w, r)
}

func (a *App) GetCategory(w http.ResponseWriter, r *http.Request) {
	handler.GetCategory(a.DB, w, r)
}

func (a *App) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	handler.UpdateCategory(a.DB, w, r)
}

func (a *App) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	handler.DeleteCategory(a.DB, w, r)
}

// Products Handlers

func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProducts(a.DB, w, r)
}

func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(a.DB, w, r)
}

func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	handler.GetProduct(a.DB, w, r)
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProduct(a.DB, w, r)
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProduct(a.DB, w, r)
}

// Images Handlers

func (a *App) GetAllImages(w http.ResponseWriter, r *http.Request) {
	handler.GetAllImages(a.DB, w, r)
}

func (a *App) CreateImage(w http.ResponseWriter, r *http.Request) {
	handler.CreateImage(a.DB, w, r)
}

func (a *App) GetImage(w http.ResponseWriter, r *http.Request) {
	handler.GetImage(a.DB, w, r)
}

func (a *App) UpdateImage(w http.ResponseWriter, r *http.Request) {
	handler.UpdateImage(a.DB, w, r)
}

func (a *App) DeleteImage(w http.ResponseWriter, r *http.Request) {
	handler.DeleteImage(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
