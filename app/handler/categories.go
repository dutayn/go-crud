package handler

import (
	"encoding/json"
	"net/http"

	"crud-product/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	categories := []model.Category{}
	db.Find(&categories)
	respondJSON(w, http.StatusOK, categories)
}

func CreateCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	category := model.Category{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&category).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, category)
}

func GetCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//title := vars["title"]
	name := vars["name"]
	category := getCategoryOr404(db, name, w, r)
	if category == nil {
		return
	}
	respondJSON(w, http.StatusOK, category)
}

func UpdateCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	category := getCategoryOr404(db, name, w, r)
	if category == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&category).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, category)
}

func DeleteCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	category := getCategoryOr404(db, name, w, r)
	if category == nil {
		return
	}
	if err := db.Delete(&category).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getCategoryOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Category {
	category := model.Category{}
	if err := db.First(&category, model.Category{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &category
}
