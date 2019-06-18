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

```bash
* Contoh GET
[
    {
        "ID": 1,
        "CreatedAt": "2019-06-18T03:23:51Z",
        "UpdatedAt": "2019-06-18T03:23:51Z",
        "DeletedAt": null,
        "name": "Elektronik",
        "enable": true,
        "categories": null
    }
] 

* Contoh POST
{
	"name" : "Elektronik",
	"enable" : true
}
```

#### http://localhost:3000/categories/:name
* `GET` : Get a categories
* `PUT` : Update a categories
* `DELETE` : Delete a categories


#### http://localhost:3000/categories/:name/product
* `GET` : Get all product of a categories
* `POST` : Create a new product in a categories

```bash
* Contoh GET Product - http://localhost:3000/categories/elektronik/product
[
    {
        "ID": 1,
        "CreatedAt": "2019-06-18T03:33:08Z",
        "UpdatedAt": "2019-06-18T03:33:08Z",
        "DeletedAt": null,
        "name": "Laptop",
        "description": "Merk DELL Latitude e7240",
        "enable": true,
        "category_id": 1,
        "products": null
    }
]


* Contoh POST Product - http://localhost:3000/categories/elektronik/product
{
	"name" : "Laptop",
	"description" : "Merk DELL Latitude e7240",
	"enable" : true
}
```

#### http://localhost:3000/categories/:name/products/:id
* `GET` : Get a product of a categories
* `PUT` : Update a product of a categories
* `DELETE` : Delete a product of a categories

#### http://localhost:3000/categories/:name/products/:name/images
* `GET` : Get a images of a products - categories
* `POST` : Create a New images of a products - categories
* `PUT` : Update a images of a products - categories
* `DELETE` : Delete a images of a products - categories

```bash
* Contoh GET Image - http://localhost:3000/categories/elektronik/product/laptop/images
[
    {
        "ID": 1,
        "CreatedAt": "2019-06-18T03:38:11Z",
        "UpdatedAt": "2019-06-18T03:38:11Z",
        "DeletedAt": null,
        "name": "Image-laptop",
        "file": "Latitude-e7240.jpg",
        "enable": true,
        "product_id": 1
    }
]

* Contoh POST Image - http://localhost:3000/categories/elektronik/product/laptop/images
{
	"name" : "Image-laptop",
	"file" : "Latitude-e7240.jpg",
	"enable" : true
}

```



