package models

import "github.com/uptrace/bun"

func (InvMast) TableName() string {
	return "inv_mast"
}

type InvMast struct {
	bun.BaseModel `bun:"table:inv_mast"`

	InvMastUid   int     `bun:"inv_mast_uid,pk"`
	ItemId       string  `bun:"item_id"`
	ItemDesc     string  `bun:"item_desc"`
	Price1       float64 `bun:"price1"`
	Price2       float64 `bun:"price2"`
	Price3       float64 `bun:"price3"`
	ExtendedDesc string  `bun:"extended_desc"`
}
