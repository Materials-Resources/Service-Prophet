package models

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type Contact struct {
	bun.BaseModel `bun:"table:contacts"`
	ID            string         `bun:"id,pk"`
	FirstName     string         `bun:"first_name"`
	MiddleInitial sql.NullString `bun:"mi"`
	LastName      string         `bun:"last_name"`
	Title         sql.NullString `bun:"title"`
	EmailAddress  sql.NullString `bun:"email_address"`
	DirectPhone   sql.NullString `bun:"direct_phone"`
}
