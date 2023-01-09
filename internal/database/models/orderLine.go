package models

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type OrderLine struct {
	bun.BaseModel `bun:"table:oe_line"`

	OrderNo       string          `bun:"order_no,pk"`
	LineNo        float32         `bun:"line_no,pk"`
	DeleteFlag    string          `bun:"delete_flag"`
	InvMastUid    int32           `bun:"inv_mast_uid"`
	RequiredDate  sql.NullTime    `bun:"required_date"`
	ExpediteDate  sql.NullTime    `bun:"expedite_date"`
	QtyOrdered    float64         `bun:"qty_ordered"`
	ExtendedPrice sql.NullFloat64 `bun:"extended_price"`
	UnitOfMeasure sql.NullString  `bun:"unit_of_measure"`
}
