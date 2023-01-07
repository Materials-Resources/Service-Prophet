package models

import "database/sql"

type Contacts struct {
	ID string `gorm:"column:id"`

	FirstName     string         `gorm:"column:first_name"`
	MiddleInitial sql.NullString `gorm:"column:mi"`
	LastName      string         `gorm:"column:last_name"`
	Title         sql.NullString `gorm:"column:title"`
	EmailAddress  sql.NullString `gorm:"column:email_address"`
	DirectPhone   sql.NullString `gorm:"column:direct_phone"`
}
