package main

import "log"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 20})
	users = append(users, User{"Daemon", "Targaeryan", "daemon@targaryaen.com", 32})
	users = append(users, User{"Ana", "deArmas", "ana@dearmas.com", 26})
	users = append(users, User{"Brad", "Pitt", "brad@pitt.com", 52})

	for _, user := range users {
		log.Println(user.LastName, user.FirstName)
	}
}
