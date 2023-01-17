package models

import (
	"github.com/uptrace/bun"
)

type InventoryReceiptLine struct {
	bun.BaseModel `bun:"table:inventory_receipts_line"`

	ReceiptNumber float64 `bun:"receipt_number,pk"`
	LineNumber    int32   `bun:"line_number,pk"`
	InvMastUid    int32   `bun:"inv_mast_uid"`
	UnitQuantity  float64 `bun:"unit_quantity"`
	UnitOfMeasure string  `bun:"unit_of_measure"`

	InventoryReceipt  InventoryReceipt `bun:"rel:belongs-to,join:receipt_number=receipt_number"`
	OrderTransactions []*InvTran       `bun:"rel:has-many,join:receipt_number=sub_document_no,join:inv_mast_uid=inv_mast_uid,join:type=trans_type,polymorphic:IRA"`
	InvMast           InvMast          `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}
