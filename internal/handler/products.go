package handler

import (
	"database/sql"
	"fmt"
)

type Products struct {
	ID      int64  `json:"id,omitempty"`
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int64  `json:"price"`
}

func (p Products) Create(db *sql.DB) error {
	_, err := db.Exec("insert into products (model, company, price) values ($1, $2, $3)", p.Model, p.Company, p.Price)

	return err
}

func (p *Products) Get(db *sql.DB) error {
	rows := db.QueryRow("select * from products where id = $1", p.ID)
	if rows.Err() != nil {
		return rows.Err()
	}

	if scanErr := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price); scanErr != nil {
		return scanErr
	}
	fmt.Println(p)

	return nil
}

func (p Products) Delete(db *sql.DB) error {
	_, err := db.Exec("delete from products where id=$1", p.ID)

	return err
}

func (p Products) Update(db *sql.DB) error {
	_, err := db.Exec("update products set price =$1 where id=$2", 9999, p.ID)

	return err
}

func Product(db *sql.DB) ([]Products, error) {
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]Products, 0)

	for rows.Next() {
		p := Products{}

		if scanErr := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price); scanErr != nil {
			return nil, scanErr
		}

		products = append(products, p)
	}

	return products, nil
}
