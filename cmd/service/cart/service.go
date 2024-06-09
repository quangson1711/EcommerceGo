package cart

import (
	"Ecommerce-Go/types"
	"fmt"
)

func getCartItemIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}

		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

func (h *Handle) createOrder(ps []types.Product, items []types.CartItem) (int, float64, error) {
	// create map product to easy check id
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}
	// caculate the total price
	// reduce quantity of products in our db
	// create the order
	// create the order items
	return 0, 0, nil
}

func checkIfCartIsInStock(items []types.CartItem, productInStore map[int]types.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		product, ok := productInStore[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d does not exist", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %d quantity less than expected %d", item.ProductID, product.Quantity)
		}
	}
	return nil
}
