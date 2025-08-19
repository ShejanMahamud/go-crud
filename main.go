package main

import (
	"fmt"
)

func addUser (u *User) {
	fmt.Println(*u)
}

func main () {
	user := &User{
		Name: "Shejan Mahamud",
		email: "dev.shejanmahamud@gmail.com",
		role: "super_admin",
	}

		user2 := &User{
		Name: "Shejan Mahamud2",
		email: "dev.shejanmahamud@gmail.com2",
		role: "super_admin2",
	}
	
	addUser(user)
	addUser(user2)
}
