package config

type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"FirstName"`
	SecondName string `json:"SecondName"`
	CreatedAt  string `json:"createdAt"`
}
