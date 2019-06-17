package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-crud/app/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllImages(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productName := vars["name"]
	product := getImageProductOr404(db, productName, w, r)
	if product == nil {
		return
	}

	images := []model.Image{}
	if err := db.Model(&product).Related(&images).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, images)
}

func CreateImage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productName := vars["name"]
	product := getImageProductOr404(db, productName, w, r)
	if product == nil {
		return
	}

	image := model.Image{ProductID: product.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&image); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&image).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, image)
}

func GetImage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productName := vars["name"]
	product := getImageProductOr404(db, productName, w, r)
	if product == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	image := getImageOr404(db, id, w, r)
	if image == nil {
		return
	}
	respondJSON(w, http.StatusOK, image)
}

func UpdateImage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productName := vars["name"]
	product := getImageProductOr404(db, productName, w, r)
	if product == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	image := getImageOr404(db, id, w, r)
	if image == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&image); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&image).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, image)
}

func DeleteImage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productName := vars["name"]
	product := getImageProductOr404(db, productName, w, r)
	if product == nil {
		return
	}

	id, _ := strconv.Atoi(vars["id"])
	image := getImageOr404(db, id, w, r)
	if image == nil {
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getImageOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Image {
	image := model.Image{}
	if err := db.First(&image, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &image
}
