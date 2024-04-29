package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID   int
	Name string
	Age  int
	Job  string
	Sex  string
}

func FetchUsers() ([]Employee, error) {
	const (
		host   = "127.0.0.1"
		port   = 5555
		user   = "postgres"
		pass   = "mypass"
		dbname = "MyDB"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age, job, sex FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []Employee
	for rows.Next() {
		var user Employee
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Job, &user.Sex)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
