package models

import (
	"database/sql"
	"time"
)

// Supplier TODO finish schema
type Supplier struct {
	SupplierId                  float64         `bun:"supplier_id,pk"`
	BuyerId                     string          `bun:"buyer_id"`
	DaysEarly                   sql.NullFloat64 `bun:"days_early"`
	DaysLate                    sql.NullFloat64 `bun:"days_late"`
	DeleteFlag                  string          `bun:"delete_flag"`
	DateCreated                 time.Time       `bun:"date_created"`
	DateLastModified            time.Time       `bun:"date_last_modified"`
	LastMaintainedBy            string          `bun:"last_maintained_by"`
	ControlValue                string          `bun:"control_value"`
	TargetValue                 sql.NullFloat64 `bun:"target_value"`
	ReviewCycle                 sql.NullFloat64 `bun:"review_cycle"`
	DateOfLastReview            sql.NullTime    `bun:"date_of_last_review"`
	DefaultCarrier              sql.NullFloat64 `bun:"default_carrier"`
	DefaultShippingInstructions sql.NullString  `bun:"default_shipping_instructions"`
	AverageLeadTime             sql.NullFloat64 `bun:"average_lead_time"`
	LeadTimeSource              sql.NullString  `bun:"lead_time_source"`
	SupplierName                sql.NullString  `bun:"supplier_name"`
	TransitDays                 sql.NullInt64   `bun:"transit_days"`
	RmaRequired                 string          `bun:"rma_required"`
}
