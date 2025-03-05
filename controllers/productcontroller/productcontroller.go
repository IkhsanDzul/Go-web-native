package productcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func IndexProduct(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		product.Name = r.FormValue("name")
		categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
		product.Category.Id = uint(categoryID)
		product.Stock, _ = strconv.ParseInt(r.FormValue("stock"), 10, 64)
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			temp, _ := template.ParseFiles("views/product/create.html")
			temp.Execute(w, nil)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp, _ := template.ParseFiles("views/product/create.html")
		temp.Execute(w, data)

		http.Redirect(w, r, "/product", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product":    product,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		if ok := productmodel.Update(id, product); !ok {
			temp, _ := template.ParseFiles("views/product/create.html")
			temp.Execute(w, nil)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		product.Name = r.FormValue("name")
		categoryID, _ := strconv.Atoi(r.FormValue("category_id"))
		product.Category.Id = uint(categoryID)
		product.Stock, _ = strconv.ParseInt(r.FormValue("stock"), 10, 64)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		temp, _ := template.ParseFiles("views/product/create.html")
		temp.Execute(w, data)

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idString)
    if err != nil {
        panic(err)
    }

    if err := productmodel.Delete(id); err != nil {
        http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
    }

    http.Redirect(w, r, "/products", http.StatusSeeOther)
}