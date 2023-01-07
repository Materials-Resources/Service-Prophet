package models

import "database/sql"

func (Customer) TableName() string {
	return "customer"
}

type Customer struct {
	CustomerID      float32        `gorm:"column:customer_id;primaryKey"`
	CompanyID       string         `gorm:"column:company_id;primaryKey"`
	CustomerName    sql.NullString `gorm:"column:customer_name"`
	CreditLimit     float32        `gorm:"column:credit_limit"`
	CreditLimitUsed float32        `gorm:"column:credit_limit_used"`

	Orders []Order `gorm:"foreignKey:customer_id,company_id;references:customer_id,company_id;"`
}
