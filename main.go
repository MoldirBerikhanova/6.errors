package main

import (
	"errors"
	"fmt"
	"strings"
)

type user struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	u := user{firstName: "Моля", lastName: "Бериканова", age: 15}
	err := createUser(u)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User created successfully!")
	}
}

func createUser(u user) error {
	err := validateUser(u)
	if err != nil {
		return err
	}

	fmt.Printf("User %s %s created with age %d\n", u.firstName, u.lastName, u.age)
	return nil
}

func validateUser(u user) error {

	if strings.TrimSpace(u.firstName) == "" {
		return errors.New("first name is required")
	}

	if strings.TrimSpace(u.lastName) == "" {
		return errors.New("last name is required")
	}

	if u.age <= 0 {
		return errors.New("age must be greater than zero")
	}
	return nil
}
