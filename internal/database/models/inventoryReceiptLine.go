package models

import "github.com/uptrace/bun"

type InventoryReceiptLine struct {
	bun.BaseModel `bun:"table:inventory_receipts_line"`

	ReceiptNumber float64 `bun:"receipt_number,pk"`
	LineNumber    int32   `bun:"line_number,pk"`
	InvMastUid    int32   `bun:"inv_mast_uid"`
	UnitQuantity  float64 `bun:"unit_quantity"`
	UnitOfMeasure string  `bun:"unit_of_measure"`

	OrderTransactions []InvTran `bun:"rel:has-many,join:receipt_number=sub_document_no,join_on:inv_mast_uid=inv_mast_uid"`
	InvMast           InvMast   `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}
