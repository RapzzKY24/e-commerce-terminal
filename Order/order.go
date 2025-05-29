package order

import (
	"time"
	cart "tubes_alpro/Cart"
)

type Order struct {
	ID           string
	CustomerName string
	Cart         cart.Cart
	TotalPrice   int
	OrderDate    time.Time
	Status       string
}

var (
	OrderHistory = []Order{} // Menyimpan riwayat semua order
)

func CreateOrder(id string, customerName string, c cart.Cart) Order {
	totalPrice := 0
	for _, item := range c.Items {
		totalPrice += item.Price * item.Quantity
	}

	newOrder := Order{
		ID:           id,
		CustomerName: customerName,
		Cart:         c,
		TotalPrice:   totalPrice,
		OrderDate:    time.Now(),
		Status:       "pending",
	}

	// Simpan ke riwayat order
	OrderHistory = append(OrderHistory, newOrder)

	return newOrder
}

func (o *Order) UpdateStatus(status string) {
	o.Status = status

	// Update status di riwayat order juga
	for i := range OrderHistory {
		if OrderHistory[i].ID == o.ID {
			OrderHistory[i].Status = status
			break
		}
	}
}

func (o *Order) CalculateTotal() {
	totalPrice := 0
	for _, item := range o.Cart.Items {
		totalPrice += item.Price * item.Quantity
	}
	o.TotalPrice = totalPrice
}

// GetAllOrders mengembalikan semua order yang pernah dibuat
func GetAllOrders() []Order {
	return OrderHistory
}

// GetOrderByID mencari order berdasarkan ID
func GetOrderByID(id string) (Order, bool) {
	for _, order := range OrderHistory {
		if order.ID == id {
			return order, true
		}
	}
	return Order{}, false
}
