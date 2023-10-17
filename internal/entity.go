package internal

import "time"

type Product struct {
	Id    int
	Name  string
	Desc  string
	Price int
}

type Order struct {
	Id       int
	Products []Product
	Total    int
	Created  time.Time
}

type User struct {
	Id         int
	Name       string
	Total      int
	Registered time.Time
}
