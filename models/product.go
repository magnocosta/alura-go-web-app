package models

import "go-web-app/db"

type Product struct {
	ID          int
	Name        string
	Description string
	Amount      float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.OpenConnection()
	result, err := db.Query("SELECT * FROM products ORDER BY id")

	if err != nil {
		panic(err)
	}

	p := Product{}
	products := []Product{}

	for result.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := result.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		p.ID = id
		p.Name = name
		p.Description = description
		p.Quantity = quantity
		p.Amount = price

		products = append(products, p)

	}

	defer db.Close()

	return products
}

func Create(name, description string, quantity int, amount float64) {
	db := db.OpenConnection()
	defer db.Close()

	query := "INSERT INTO products (name, description, quantity, price) VALUES ($1, $2, $3, $4)"

	s, e := db.Prepare(query)

	if e != nil {
		panic(e)
	}

	s.Exec(name, description, quantity, amount)
}

func Delete(id string) {
	db := db.OpenConnection()
	defer db.Close()

	query := "DELETE FROM products WHERE id=$1"

	s, e := db.Prepare(query)

	if e != nil {
		panic(e)
	}

	s.Exec(id)
}

func GetProductById(id string) Product {
	db := db.OpenConnection()
	defer db.Close()

	query := "SELECT * FROM products WHERE id=$1"

	s, e := db.Prepare(query)

	if e != nil {
		panic(e)
	}

	r := s.QueryRow(id)

	p := Product{}

	var pid, quantity int
	var name, description string
	var price float64

	err := r.Scan(&pid, &name, &description, &price, &quantity)

	if err != nil {
		panic(err)
	}

	p.ID = pid
	p.Name = name
	p.Description = description
	p.Quantity = quantity
	p.Amount = price

	return p
}

func Update(id, name, description string, quantity int, amount float64) {
	db := db.OpenConnection()
	defer db.Close()

	query := "UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5"

	s, e := db.Prepare(query)

	if e != nil {
		panic(e)
	}

	s.Exec(name, description, amount, quantity, id)
}
