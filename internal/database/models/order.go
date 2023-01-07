package models

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel        `bun:"table:oe_hdr"`
	OrderNo              string         `bun:"order_no,pk"`
	Approved             sql.NullString `bun:"approved"`
	ContactID            sql.NullString `bun:"contact_id"`
	CompanyID            string         `bun:"company_id"`
	Completed            sql.NullString `bun:"completed"`
	CustomerID           float64        `bun:"customer_id"`
	DeleteFlag           string         `bun:"delete_flag"`
	DeliveryInstructions sql.NullString `bun:"delivery_instructions"`
	OrderDate            sql.NullTime   `bun:"order_date"`
}
