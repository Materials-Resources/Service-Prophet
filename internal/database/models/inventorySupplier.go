package models

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplier struct {
	bun.BaseModel          `bun:"table:inventory_supplier"`
	InvMastUid             int32           `bun:"inv_mast_uid,unique"`
	SupplierId             float32         `bun:"supplier_id,unique"`
	DivisionId             float32         `bun:"division_id,unique"`
	LeadTimeDays           float32         `bun:"lead_time_days"`
	UpcCode                sql.NullString  `bun:"upc_code"`
	CheckDigit             sql.NullFloat64 `bun:"check_digit"`
	CatalogName            sql.NullString  `bun:"catalog_name"`
	DeleteFlag             string          `bun:"delete_flag"`
	DateCreated            time.Time       `bun:"date_created"`
	DateLastModified       time.Time       `bun:"date_last_modified"`
	LastMaintainedBy       string          `bun:"last_maintained_by"`
	SupplierPartNo         sql.NullString  `bun:"supplier_part_no"`
	ListPrice              float32         `bun:"list_price"`
	Cost                   float32         `bun:"cost"`
	BackhaulType           string          `bun:"backhaul_type"`
	InventorySupplierUid   int32           `bun:"inventory_supplier_uid,pk"`
	ContractNumber         string          `bun:"contract_number"`
	CreatedBy              string          `bun:"created_by"`
	DateCostLastModified   sql.NullTime    `bun:"date_cost_last_modified"`
	UpcCodeSourceTypeCd    sql.NullInt32   `bun:"upc_code_source_type_cd"`
	FreightFactor          float32         `bun:"freight_factor"`
	MinimumPurchaseQty     float32         `bun:"minimum_purchase_qty"`
	IncrementalPurchaseQty float32         `bun:"incremental_purchase_qty"`

	InventorySupplierXLoc InventorySupplierXLoc `bun:"rel:has-one,join:supplier_id=inventory_supplier_uid"`
	InvMast               InvMast               `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
}
