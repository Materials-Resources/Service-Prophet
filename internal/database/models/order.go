package models

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type Order struct {
	bun.BaseModel        `bun:"table:oe_hdr"`
	OrderNo              string         `bun:"order_no,pk"`
	Approved             sql.NullString `bun:"approved"`
	ContactId            sql.NullString `bun:"contact_id"`
	CompanyId            string         `bun:"company_id"`
	Completed            sql.NullString `bun:"completed"`
	CustomerId           float64        `bun:"customer_id"`
	DeleteFlag           string         `bun:"delete_flag"`
	DeliveryInstructions sql.NullString `bun:"delivery_instructions"`
	DateCreated          time.Time      `bun:"date_Created"`
	DateModified         time.Time      `bun:"date_last_modified"`
	LastMaintainedBy     string         `bun:"last_maintained_by"`
	OrderDate            sql.NullTime   `bun:"order_date"`
	PoNo                 sql.NullString `bun:"po_no"`
	ShipToName           sql.NullString `bun:"ship2_name"`
	ShipToAdd1           sql.NullString `bun:"ship2_add1"`
	ShipToAdd2           sql.NullString `bun:"ship2_add2"`
	ShipToCity           sql.NullString `bun:"ship2_city"`
	ShipToState          sql.NullString `bun:"ship2_state"`
	ShipToZip            sql.NullString `bun:"ship2_zip"`
	ShipToCountry        sql.NullString `bun:"ship2_country"`
	OeHdrUid             int            `bun:"oe_hdr_uid,unique"`
	SourceCodeNo         int            `bun:"source_code_no"`
	InvoiceBatchUid      int            `bun:"invoice_batch_uid,default:1"`
	FreightOut           float64        `bun:"freight_out,default:0"`
	ExcludeRebates       string         `bun:"exclude_rebates"`
	CaptureUsageDefault  string         `bun:"capture_usage_default"`
	FrontCounterRma      string         `bun:"front_counter_rma,default:N"`
	ProfitPercent        float64        `bun:"profit_percent,default:0"`
	OrderCostBasis       int            `bun:"order_cost_basis,default:1"`

	Customer Customer `bun:"rel:has-one,join:customer_id=customer_id,join_on:company_id=company_id"`

	OrderLine []*OrderLine `bun:"rel:has-many,join:order_no=order_no"`
}
