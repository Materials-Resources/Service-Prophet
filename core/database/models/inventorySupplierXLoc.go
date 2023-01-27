package models

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

// InventorySupplierXLoc TODO Finish Schema
type InventorySupplierXLoc struct {
	bun.BaseModel              `bun:"table:inventory_supplier_x_loc"`
	InventorySupplierXLocUid   int32           `bun:"inventory_supplier_x_loc_uid,pk"`
	InventorySupplierUid       int32           `bun:"inventory_supplier_uid"`
	LocationId                 float64         `bun:"location_id"`
	PrimarySupplier            string          `bun:"primary_supplier"`
	AverageLeadTime            int16           `bun:"average_lead_time"`
	RowStatusFlag              int32           `bun:"row_status_flag"`
	DateCreated                time.Time       `bun:"date_created"`
	DateLastModified           time.Time       `bun:"date_last_modified"`
	LastMaintainedBy           string          `bun:"last_maintained_by"`
	LocListPrice               sql.NullFloat64 `bun:"loc_list_price"`
	LocCost                    sql.NullFloat64 `bun:"loc_cost"`
	FixedLeadTime              sql.NullInt16   `bun:"fixed_lead_time"`
	VmiStatus                  sql.NullInt32   `bun:"vmi_status"`
	CreatedBy                  sql.NullString  `bun:"created_by"`
	OverrideVmiStatus          sql.NullInt32   `bun:"override_vmi_status"`
	OverrideVmiStatusFlag      string          `bun:"override_vmi_status_flag"`
	KeyVmiIndicatorChangedFlag string          `bun:"key_vmi_indicator_changed_flag"`
}
