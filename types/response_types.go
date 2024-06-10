package types

type LoginResponseBody struct {
	Token string `json:"token"`
}

func NewLoginResponseBody(token string) *LoginResponseBody {
	return &LoginResponseBody{
		Token: token,
	}
}

type CheckoutCartResponseBody struct {
	TotalPrice float64 `json:"total_price"`
	OrderID    int     `json:"order_id"`
}

func NewCheckoutCartResponseBody(totalPrice float64, orderID int) *CheckoutCartResponseBody {
	return &CheckoutCartResponseBody{totalPrice, orderID}
}
