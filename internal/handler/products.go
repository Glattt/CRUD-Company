package handler

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	ID      int64  `json:"id,omitempty"`
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int64  `json:"price"`
}

func (p Product) Create(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(ctx, "insert into products (model, company, price) values ($1, $2, $3)", p.Model, p.Company, p.Price)

	return err
}

func (p *Product) Get(ctx context.Context, db *pgx.Conn) error {
	rows := db.QueryRow(ctx, "select * from products where id = $1", p.ID)
	if rows == nil {
		return pgx.ErrNoRows
	}

	if scanErr := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price); scanErr != nil {
		return scanErr
	}

	return nil
}

func (p Product) Delete(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(ctx, "delete from products where id=$1", p.ID)

	return err
}

func (p Product) Update(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(ctx, "update products set price =$1 where id=$2", p.Price, p.ID)

	return err
}

func Products(ctx context.Context, db *pgx.Conn) ([]Product, error) {
	fmt.Println("get products")

	rows, err := db.Query(ctx, `select * from products`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]Product, 0)

	for rows.Next() {
		p := Product{}

		if scanErr := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price); scanErr != nil {
			return nil, scanErr
		}

		products = append(products, p)
	}

	return products, nil
}
