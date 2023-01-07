package models

import (
	"database/sql"
)

func (Order) TableName() string {
	return "oe_hdr"
}

type Order struct {
	OrderNo              string         `gorm:"column:order_no;primaryKey"`
	Approved             sql.NullString `gorm:"column:approved"`
	ContactID            sql.NullString `gorm:"column:contact_id"`
	CompanyID            sql.NullString `gorm:"column:company_id"`
	Completed            sql.NullString `gorm:"column:completed"`
	CustomerID           float32        `gorm:"column:customer_id"`
	DeleteFlag           string         `gorm:"column:delete_flag"`
	DeliveryInstructions sql.NullString `gorm:"column:delivery_instructions"`
	OrderDate            sql.NullTime   `gorm:"column:order_date"`
}
