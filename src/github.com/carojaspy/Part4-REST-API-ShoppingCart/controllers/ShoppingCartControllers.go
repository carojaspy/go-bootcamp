package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/carojaspy/Part4-REST-API-ShoppingCart/models"
	"github.com/gorilla/mux"
)

type Todo struct {
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

// IndexPage .
func IndexPage(w http.ResponseWriter, r *http.Request) {

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

	fp := path.Join("src", "github.com", "carojaspy", "templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Todo{Cart: cart, Articles: AvailableArticles}

	// Rendering HTML, passing to the HTML the AvailableArticles variable
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateShopingCart .
func CreateShoppingCart(w http.ResponseWriter, r *http.Request) {
	// Create new shopping cart
	log.Println("CreateShopingCart")
	createdCart = models.Cart{ID: "SoyUnCarritoCreado"}
	json.NewEncoder(w).Encode(createdCart)
}

// GetShopingCart .
func GetShopingCart(w http.ResponseWriter, r *http.Request) {
	// Returning the shopping cart
	log.Println("GetShopingCart")
	json.NewEncoder(w).Encode(cart)
}

// GetShopingCart .
func DeleteShopingCart(w http.ResponseWriter, r *http.Request) {
	//  @TODO, delete shopping cart
	log.Println("GetShopingCart")
	json.NewEncoder(w).Encode(cart)
}

// AddToShoppingCart .
func AddToShoppingCart(w http.ResponseWriter, r *http.Request) {
	// Returning the shopping cart
	log.Println("AddToShoppingCart")
	json.NewEncoder(w).Encode(cart)
}

// RemoveToShoppingCart .
func RemoveToShoppingCart(w http.ResponseWriter, r *http.Request) {
	// @TODO Remove element from the shopping cart
	log.Println("RemoveToShoppingCart")
	json.NewEncoder(w).Encode(cart)
}

// GetAllArticles .
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	// Returning the shopping cart
	log.Println("GetShopingCart")
	json.NewEncoder(w).Encode(cart)
}

// GetArticle .
func GetArticle(w http.ResponseWriter, r *http.Request) {
	// Returning the shopping cart
	log.Println("GetShopingCart")
	json.NewEncoder(w).Encode(cart)
}

// GetPeople .
func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("GetPeople")
	json.NewEncoder(w).Encode(people)
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
