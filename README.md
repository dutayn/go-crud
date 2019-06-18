# Go-CRUD

RESTful API menggunakan framework **gorilla/mux**, **gorm** (An ORM for Go) & **MYSQL**

## Install dan Menjalankannya
```bash
# Download project
go get github.com/dutayn/go-crud
```

Sebelum menjalankan server API, lakukan konfigurasi database terlebih dahulu -  [config.go](https://github.com/dutayn/go-crud/blob/master/config/config.go)


```bash
# Run
cd go-crud
go run main.go

# API Endpoint : http://localhost:3000
```

## API

#### http://localhost:3000/categories
* `GET` : Get all categories
* `POST` : Create a new categories

#### http://localhost:3000/categories/:name
* `GET` : Get a categories
* `PUT` : Update a categories
* `DELETE` : Delete a categories

#### http://localhost:3000/categories/:name/product
* `GET` : Get all product of a categories
* `POST` : Create a new product in a categories

#### http://localhost:3000/categories/:name/products/:id
* `GET` : Get a product of a categories
* `PUT` : Update a product of a categories
* `DELETE` : Delete a product of a categories

#### http://localhost:3000/categories/:name/products/:name/images
* `GET` : Get a images of a products - categories
* `POST` : Create a New images of a products - categories
* `PUT` : Update a images of a products - categories
* `DELETE` : Delete a images of a products - categories

