package models

import "time"

type Order struct {
	customerId    float64
	requestedDate time.Time
	poNo          string
	items         []*OrderItem
}
type OrderItem struct {
	invMastUid     int
	qtyOrdered     float64
	unitPrice      float64
	requiredDate   time.Time
	customerPartNo string
}
