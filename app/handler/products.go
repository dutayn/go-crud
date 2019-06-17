package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud-product/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categoryName := vars["name"]
	category := getCategoryOr404(db, categoryName, w, r)
	if category == nil {
		return
	}

	products := []model.Product{}
	if err := db.Model(&category).Related(&products).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, products)
}

func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categoryName := vars["name"]
	category := getCategoryOr404(db, categoryName, w, r)
	if category == nil {
		return
	}

	product := model.Product{CategoryID: category.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, product)
}

func GetProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categoryName := vars["name"]
	category := getCategoryOr404(db, categoryName, w, r)
	if category == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	product := getProductOr404(db, id, w, r)
	if product == nil {
		return
	}
	respondJSON(w, http.StatusOK, product)
}

func UpdateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categoryName := vars["name"]
	category := getCategoryOr404(db, categoryName, w, r)
	if category == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	product := getProductOr404(db, id, w, r)
	if product == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, product)
}

func DeleteProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categoryName := vars["name"]
	category := getCategoryOr404(db, categoryName, w, r)
	if category == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	product := getProductOr404(db, id, w, r)
	if product == nil {
		return
	}

	if err := db.Delete(&category).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getProductOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Product {
	product := model.Product{}
	if err := db.First(&product, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &product
}

func getImageProductOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Product {
	imageProduct := model.Product{}
	if err := db.First(&imageProduct, model.Product{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &imageProduct
}
