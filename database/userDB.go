package users

import (
	"time"

	config "GoodApi/database/config"
)

func Getusers() []config.Employee {
	CreatedAtDB := time.Now()
	users := []config.Employee{
		{ID: 1, FirstName: "John", SecondName: "Doe", CreatedAt: CreatedAtDB},
		{ID: 2, FirstName: "Jane", SecondName: "Doe", CreatedAt: CreatedAtDB},
		{ID: 3, FirstName: "Bob", SecondName: "Smith", CreatedAt: CreatedAtDB},
	}
	return users
}
