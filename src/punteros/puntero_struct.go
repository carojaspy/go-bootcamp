package main

import ("fmt")

type User struct {
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func UpdateStatus(u *User) {
	// Using * operator
	(*u).IsActive = !(*u).IsActive
}

func update_status(u *User){
	// Using . notation
	u.IsActive = !u.IsActive
}

func main() {

	u := User{
		FirstName: "Carlos",
		LastName:  "Rojas",
		Email:     "punteros@example.com",
		IsActive:  false,
	}
	// Printing results
	fmt.Println("User antes de actualización: ", u)
	UpdateStatus(&u)
	fmt.Println("User después de actualización: ", u)
	update_status(&u)
	fmt.Println("User actualizado con notacion '.' :  ", u)
}

