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

	QtyOnHand             float64
	QtyInProcess          float64
	DateCreated           time.Time
	DateLastModified      time.Time
	LastMaintainedBy      string
	LastRecPo             float64
	LastRecPoWithDisc     float64
	GlAccountNo           string
	PurchOrTransfer       string
	NextDueInPoCost       float64
	RevenueAccountNo      string
	CosAccountNo          string
	Sellable              string
	MovingAverageCost     float64
	StandardCost          float64
	ProtectedStockQty     float64
	InvMin                float64
	InvMax                float64
	Stockable             string
	ReplenishmentLocation float64
	AverageMonthlyUsage   float64
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
