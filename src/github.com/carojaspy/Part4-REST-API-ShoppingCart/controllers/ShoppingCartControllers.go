package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/carojaspy/Part4-REST-API-ShoppingCart/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Todo .
type Todo struct{
	Cart     models.Cart
	Articles []models.Article
}

var people []models.Person
var createdCart models.Cart

var cart = models.Cart{ID: "1234"}
var AvailableArticles []models.Article

//people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
//people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
// models.populatePerson(people)

func setJSONResponse(w *http.ResponseWriter) {
	//Set Header
	// Set content type to response
	(*w).Header().Set("Content-Type", "application/json")
}

// IndexPage .
func IndexPage(w http.ResponseWriter, r *http.Request) {

func GetArticlesFromProvider(articles *[]models.Article) error {
	/*Getting all articles available from */
	querystring := fmt.Sprintf("http://challenge.getsandbox.com/articles")
	resp, err := http.Get(querystring)
	if err != nil {
		// handle error
		log.Fatal("Error getting Articles from challenge.getsandbox.com")
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Code Status invalid: ", resp.StatusCode)
		return nil
		//http.Error(w, "Error Trying to get Available Articles", http.StatusInternalServerError)
	}
	// Closing conection
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//	Reading response
	if err != nil {
		log.Fatal("Error reading result from challenge.getsandbox.com")
	}

	// Trying to convert the response to struct Article
	errUn := json.Unmarshal(body, &articles)
	if errUn != nil {
		log.Fatal("Error in unmarshall: ", errUn)
	}
	return nil
} //end

	fp := path.Join("src", "github.com", "carojaspy", "Part4-REST-API-ShoppingCart", "templates", "index.html")
// IndexPage .
func IndexPage(w http.ResponseWriter, r *http.Request) {
	err := GetArticlesFromProvider(&AvailableArticles)

	fp := path.Join("src", "github.com", "carojaspy", "Part4-REST-API-ShoppingCart", "templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Println("IndexPage error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Todo{Cart: cart, Articles: AvailableArticles}

	// Rendering HTML, passing to the HTML the AvailableArticles variable
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("IndexPage error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println("IndexPage 200")
}

// CreateShoppingCart .
func CreateShoppingCart(w http.ResponseWriter, r *http.Request) {
	// Create new shopping cart
	log.Println("CreateShopingCart")
	createdCart = models.Cart{ID: "SoyUnCarritoCreado"}

	// Set content type to response
	setJSONResponse(&w)
	json.NewEncoder(w).Encode(createdCart)
}

// GetShopingCart .
func GetShopingCart(w http.ResponseWriter, r *http.Request) {
	// Returning the shopping cart
	log.Println("GetShopingCart")
	log.Println(createdCart)
	createdCart = models.Cart{ID: "SoyUnCarritoCreado"}

	if createdCart.ID == "" {
		http.Error(w, "Error, doesn't exist a Cart, first create a cart", http.StatusInternalServerError)
		log.Fatal("Error in GetShopingCart, No initialiced Cart")
		log.Fatal("Doesn't exists a Cart")
	} else {
		// Set content type to response
		setJSONResponse(&w)
		json.NewEncoder(w).Encode(createdCart)
	}
}

// DeleteShopingCart .
func DeleteShopingCart(w http.ResponseWriter, r *http.Request) {
	//  @TODO, delete shopping cart
	log.Println("GetShopingCart")
	json.NewEncoder(w).Encode(cart)
	if createdCart.ID == "" {
		log.Fatal("Doesn't exists a Cart")
	}
	// Reseting Cart
	createdCart = models.Cart{}

	// Set content type to response
	setJSONResponse(&w)

}

// AddToShoppingCart .
func AddToShoppingCart(w http.ResponseWriter, r *http.Request) {
	// Add a new Product to the Cart
	log.Println("AddToShoppingCart")
	err := GetArticlesFromProvider(&AvailableArticles)
	if err != nil {
		log.Println("Error trying to get Items")
	}
	log.Println("elements: ", AvailableArticles)

	vars := mux.Vars(r)
	idItem := vars["id"] // Id of the Item
	log.Println("Id of URL: ", idItem)
	// Trying to get Item Element

	json.NewEncoder(w).Encode(cart)
}

// RemoveToShoppingCart .
func RemoveToShoppingCart(w http.ResponseWriter, r *http.Request) {
	// @TODO Remove element from the shopping cart
	log.Println("RemoveToShoppingCart")
	r.ParseForm() // parse arguments,
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	json.NewEncoder(w).Encode(createdCart)
	// Set content type to response
	setJSONResponse(&w)
}

// GetAllArticles .
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	// Returning the shopping cart
	log.Println("GetShopingCart")

	/*Getting all articles available from */
	querystring := fmt.Sprintf("http://challenge.getsandbox.com/articles")
	resp, err := http.Get(querystring)
	if err != nil {
		// handle error
		log.Fatal("Error getting Articles from challenge.getsandbox.com")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Code Status invalid: ", resp.StatusCode)
		http.Error(w, "Error Trying to get Available Articles", http.StatusInternalServerError)
	}
	// Closing conection
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//	Reading response
	if err != nil {
		log.Fatal("Error reading result from challenge.getsandbox.com")
	}

	// Trying to convert the response to struct Article
	errUn := json.Unmarshal(body, &AvailableArticles)
	if errUn != nil {
		log.Fatal("Error in unmarshall: ", errUn)
	}

	// Set content type to response
	setJSONResponse(&w)
	json.NewEncoder(w).Encode(AvailableArticles)
}

// GetArticle .
func GetArticle(w http.ResponseWriter, r *http.Request) {
	// Returning the shopping cart
	log.Println("GetShopingCart")
	r.ParseForm() // parse arguments,
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// Set content type to response
	setJSONResponse(&w)
	json.NewEncoder(w).Encode(createdCart)
}

// GetPerson .
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			log.Println("Object obtained")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreatePerson .
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson .
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			log.Println("DeletePerson, object obtained", index)
			//people = append(people[:index], people[index+1]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
