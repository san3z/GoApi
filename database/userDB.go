package users

import (
	"fmt"
	"time"
)

func Getusers() []Employee {
	createdAt := time.Now()
	formattedCreatedAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		createdAt.Year(), createdAt.Month(), createdAt.Day(),
		createdAt.Hour(), createdAt.Minute(), createdAt.Second())
	users := []Employee{
		{ID: 1, FirstName: "John", SecondName: "Doe", CreatedAt: formattedCreatedAt},
		{ID: 2, FirstName: "Jane", SecondName: "Doe", CreatedAt: formattedCreatedAt},
		{ID: 3, FirstName: "Bob", SecondName: "Smith", CreatedAt: formattedCreatedAt},
	}
	return users
}
