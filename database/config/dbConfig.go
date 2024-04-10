package config

import "time"

type Employee struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"FName"`
	SecondName string    `json:"sName"`
	CreatedAt  time.Time `json:"createdAt"`
}
