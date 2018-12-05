package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/carojaspy/Part4-REST-API-ShoppingCart/controllers"
	"github.com/gorilla/mux"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

func main() {
	log.Print("Running Server on localhost:8000 ...")
	router := mux.NewRouter()

	// Index Router, display HTML page
	router.HandleFunc("/", controllers.IndexPage).Methods("GET")

	//	Get, Create and update Shopping Cart
	router.HandleFunc("/cart", controllers.GetShopingCart).Methods("GET")
	router.HandleFunc("/cart", controllers.CreateShoppingCart).Methods("POST")
	router.HandleFunc("/cart", controllers.DeleteShopingCart).Methods("DELETE")

	//	Operations over the shopping cart
	router.HandleFunc("/cart/{id}", controllers.AddToShoppingCart).Methods("POST")
	router.HandleFunc("/cart/{id}", controllers.RemoveToShoppingCart).Methods("DELETE")

	//	Articles operations
	router.HandleFunc("/articles", controllers.GetAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", controllers.GetArticle).Methods("GET")

	// Running server
	log.Fatal(http.ListenAndServe(":8000", router))
}
