package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/julienschmidt/httprouter"
)

func main() {
	initStorage()

	r := httprouter.New()
	r.GET("/", index)
	r.GET("/products/", allProducts)
	r.POST("/products/", addProducts)
	r.GET("/products/:id", productByID)
	r.POST("/products/:name", addProduct)

	log.Fatal(http.ListenAndServe(":8080", r))
}

type product struct {
	name string
	id   string
}

func initStorage() {
	migration, _ := migrate.New("file://migrations", "mysql://root:foomeow@tcp/mysql")
	migration.Up()
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "This is the Morena API")
}

func allProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "Shows All Products")
	// TODO: get several products from db
}

func addProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "POST to add products")
	// TODO: check whether this is correct because it looks like there's only 1 product added at a time
	// TODO: add several products from db
}

func productByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Displaying product with ID: %s", ps.ByName("id"))
	w.WriteHeader(http.StatusNoContent)
}

func addProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Trying to add product with name: %s", ps.ByName("name"))
}
