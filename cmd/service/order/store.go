package order

import (
	"Ecommerce-Go/types"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order *types.Order) error {
	result, err := s.db.Exec("INSERT INTO orders (userId, total, status, address) VALUES (?,?,?,?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	order.ID = int(lastID)

	return nil
}

func (s *Store) CreateOrderItem(orderItem *types.OrderItem) error {
	result, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES (?,?,?,?)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	orderItem.ID = int(lastID)

	return nil
}
