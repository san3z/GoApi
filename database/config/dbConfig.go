package config

type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"FName"`
	SecondName string `json:"sName"`
	CreatedAt  string `json:"createdAt"`
}
