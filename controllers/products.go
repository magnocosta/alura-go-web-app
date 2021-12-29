package controllers

import (
	"go-web-app/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")
		amount := r.FormValue("amount")

		priceConv, e := strconv.ParseFloat(amount, 64)
		if e != nil {
			panic(e.Error())
		}

		quantityConv, e := strconv.Atoi(quantity)
		if e != nil {
			panic(e.Error())
		}

		models.Create(name, description, quantityConv, priceConv)
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.Delete(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.GetProductById(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")
		amount := r.FormValue("amount")

		priceConv, e := strconv.ParseFloat(amount, 64)
		if e != nil {
			panic(e.Error())
		}

		quantityConv, e := strconv.Atoi(quantity)
		if e != nil {
			panic(e.Error())
		}

		models.Update(id, name, description, quantityConv, priceConv)
	}

	http.Redirect(w, r, "/", 301)
}
