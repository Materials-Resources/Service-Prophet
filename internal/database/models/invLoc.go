package models

import (
	"github.com/uptrace/bun"
	"time"
)

type InvLoc struct {
	bun.BaseModel `bun:"table:inv_loc"`

	InvMastUid int32   `bun:"inv_mast_uid,pk"`
	LocationId float32 `bun:"location_id,pk"`
	CompanyId  string  `bun:"company_id,pk"`

	QtyOnHand             float32
	QtyInProcess          float32
	DateCreated           time.Time
	DateLastModified      time.Time
	LastMaintainedBy      string
	LastRecPo             float32
	LastRecPoWithDisc     float32
	GlAccountNo           string
	PurchOrTransfer       string
	NextDueInPoCost       float32
	RevenueAccountNo      string
	CosAccountNo          string
	Sellable              string
	MovingAverageCost     float32
	StandardCost          float32
	ProtectedStockQty     float32
	InvMin                float32
	InvMax                float32
	Stockable             string
	ReplenishmentLocation float32
	AverageMonthlyUsage   float32
	ProductGroupId        string `bun:"product_group_id"`
	PurchaseDiscountGroup string
	SalesDiscountGroup    string
	NoCharge              string
	PriceOne              float32 `bun:"price1"`
	PriceTwo              float32 `bun:"price2"`
	PriceThree            float32 `bun:"price3"`
	PriceFour             float32 `bun:"price4"`
	PriceFive             float32 `bun:"price5"`
	PriceSix              float32 `bun:"price6"`
	PriceSeven            float32 `bun:"price7"`
	PriceEight            float32 `bun:"price8"`
	PriceNine             float32 `bun:"price9"`
	PriceTen              float32 `bun:"price10"`
	ReplenishmentMethod   string

	InvMast InvMast `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
}
