package db

import "fmt"

// TableDB struct to store info
type TableDB struct {
	name string
	data map[string]interface{}
}

// InMermoryDB Declaring Structure to store info
type InMermoryDB struct {
	tables []TableDB
	Name   string
	data   map[string]interface{}
}

// Movie : Example of data to persists
type Movie struct {
	TableDB
	id           int
	title        string
	directorName string
	rate         float32
}

// InitDB .
func (db InMermoryDB) InitDB(name string) InMermoryDB {
	fmt.Println("CreateDB")
	fmt.Printf("Creating database '%s'... \n", name)
	db.Name = name
	db.data = make(map[string]interface{})
	fmt.Println(db.Name, db.data, db.tables)
	return InMermoryDB{Name: name}
}

// CreateDB .
func (db InMermoryDB) CreateDB(name string) {
	fmt.Println("CreateDB")
	fmt.Printf("Creating database '%s'... \n", name)
	db.Name = name
	db.data = make(map[string]interface{})
	fmt.Println(db.Name, db.data, db.tables)
}

// CreateTable .
func (db InMermoryDB) CreateTable(name string) TableDB {
	fmt.Printf("CreateTable of %s ...\n", name)
	fmt.Println(db.Name, db.data, db.tables)
	return TableDB{name: name}
}

// Insert .
func (db InMermoryDB) Insert(key string, data interface{}) {
	fmt.Println("Insert data")
	fmt.Printf("db.name: %s   - db.data: %s \n", db.Name, db.data)

	fmt.Println("Data to insert : ", key, ":", data)
	db.data[key] = data
}

// Update .
func (db InMermoryDB) Update() {
	fmt.Println("Update")
}

// Delete .
func (db InMermoryDB) Delete() {
	fmt.Println("Delete")
}

// Retrieve .
func (db InMermoryDB) Retrieve() {
	fmt.Println("Retrieve")
}
