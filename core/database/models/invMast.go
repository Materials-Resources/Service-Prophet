package models

import "github.com/uptrace/bun"

func (InvMast) TableName() string {
	return "inv_mast"
}

type InvMast struct {
	bun.BaseModel `bun:"table:inv_mast"`

	InvMastUid          int32   `bun:"inv_mast_uid,pk"`
	ItemId              string  `bun:"item_id"`
	ItemDesc            string  `bun:"item_desc"`
	Price1              float64 `bun:"price1"`
	Price2              float64 `bun:"price2"`
	Price3              float64 `bun:"price3"`
	DefaultProductGroup string  `bun:"default_product_group"`
	ExtendedDesc        string  `bun:"extended_desc"`

	InvLoc *InvLoc `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}
