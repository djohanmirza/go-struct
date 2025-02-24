package entity

type Orders struct {
	OrderID     string      `json:"order_id" gorm:"primary_key"`
	CustomerID  string      `json:"customer_id"`
	OrderDate   string      `json:"order_date"`
	TotalAmount float64     `json:"total_amount"`
	OrderItems  []OrderItem `json:"order_items" gorm:"foreignkey:OrderID;constarint:OnDelete:CASCADE"`
}

type OrderItem struct {
	ProductID  string  `json:"product_id" gorm:"primary_key"`
	Quantity   int     `json:"quantity" gorm:"primary_key"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
