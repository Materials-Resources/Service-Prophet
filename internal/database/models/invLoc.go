package models

import "github.com/uptrace/bun"

type InvLoc struct {
	bun.BaseModel `bun:"table:inv_loc"`

	InvMastUid int     `bun:"inv_mast_uid,pk"`
	LocationId float64 `bun:"location_id,pk"`
	CompanyId  string  `bun:"company_id,pk"`

	ProductGroupId string  `bun:"product_group_id"`
	PriceOne       float64 `bun:"price1"`
	PriceTwo       float64 `bun:"price2"`
	PriceThree     float64 `bun:"price3"`

	InvMast InvMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}
