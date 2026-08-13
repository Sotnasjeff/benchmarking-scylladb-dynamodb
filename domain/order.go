package domain

type Order struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Status     string  `json:"status"`
	Total      float64 `json:"total"`
}

func (order *Order) GetID() string {
	return order.ID
}

func (order *Order) GetCustomerID() string {
	return order.CustomerID
}

func (order *Order) GetStatus() string {
	return order.Status
}

func (order *Order) GetTotal() float64 {
	return order.Total
}

func (order *Order) SetTotal(total float64) {
	order.Total = total
}

func (order *Order) SetCustomerID(customerID string) {
	order.CustomerID = customerID
}

func (order *Order) SetStatus(status string) {
	order.Status = status
}
