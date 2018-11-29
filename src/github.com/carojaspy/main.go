package main

import (
	"fmt"
	"reflect"

	"github.com/carojaspy/db"
)

func main() {
	fmt.Println("Main db")

	/*
		myDB := db.InMermoryDB.InitDB("taco")
		fmt.Println(reflect.TypeOf(myDB))
		fmt.Printf("Name of my DB: %s \n", myDB.Name)
	*/

	myDB := db.InMermoryDB{}
	myDB.CreateDB("Taco")
	fmt.Println(reflect.TypeOf(myDB))
	fmt.Printf("Name of my DB: %s \n", myDB.Name)

	// var movieTable db.Movie
	// movieTable = myDB.CreateTable("Movies")

	myMap := map[string]interface{}{"name": "taco", "id": 1}
	myDB.Insert("key1", myMap)

}
