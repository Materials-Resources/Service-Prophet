package models

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type Customer struct {
	bun.BaseModel   `bun:"table:customer"`
	CustomerID      float64        `bun:"customer_id,pk"`
	CompanyID       string         `bun:"company_id,pk"`
	CustomerName    sql.NullString `bun:"customer_name"`
	CreditLimit     float32        `bun:"credit_limit"`
	CreditLimitUsed float32        `bun:"credit_limit_used"`

	Orders []*Order `bun:"rel:has-many,join:customer_id=customer_id,join_on:company_id=company_id"`
}
