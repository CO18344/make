package main 

import (
	"fmt"
)

type User struct {
	name string
	age int 
	dob string 
}
func main(){
	// fmt.Println("Hello");	
	fmt.Println("Uses of make keyword in Go");

	// with Slices
	var users []User; 
	users = make([]User, 5);


	users[0] = User{
		name: "Qarim",
		age: 44,
		dob: "21/5/1999",
	}

	users[1] = User{
		name: "Flexa",
		age: 32,
		dob: "10/10/1992",
	}
	fmt.Println(users)
	

}	