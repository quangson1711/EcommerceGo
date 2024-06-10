package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductByID(id int) (*Product, error)
	GetProductByIDs(ids []int) ([]Product, error)
	CreateProduct(Product) error
	UpdateQuantityProduct(Product) error
}

type OrderStore interface {
	CreateOrder(*Order) error
	CreateOrderItem(*OrderItem) error
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreateAt  time.Time `json:"createAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Images      []string  `json:"images"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreateAt    time.Time `json:"createAt"`
}
