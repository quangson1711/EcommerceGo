package product

import (
	"Ecommerce-Go/types"
	"database/sql"
	"fmt"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

func (s *Store) GetProductByIDs(ids []int) ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)

	for rows.Next() {
		p, err = scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return p, nil
}

func (s *Store) CreateProduct(p types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity, createdAt) VALUES (?, ?, ?, ?, ?, ?)", p.Name, p.Description, p.Images, p.Price, p.Quantity, p.CreateAt)
	if err != nil {
		return err
	}
	return nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	p := new(types.Product)
	var imagesDB string
	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&imagesDB,
		&p.Price,
		&p.Quantity,
		&p.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	p.Images = strings.Split(imagesDB, ",")

	return p, nil
}

func (s *Store) UpdateQuantityProduct(product types.Product) error {

	_, err := s.db.Exec("UPDATE products SET quantity = ? WHERE id = ?", product.Quantity, product.ID)
	if err != nil {
		return err
	}
	return nil
}
