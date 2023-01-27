package models

import (
	"github.com/uptrace/bun"
)

type InventoryReceipt struct {
	bun.BaseModel `bun:"table:inventory_receipts_hdr"`

	ReceiptNumber float64 `bun:"receipt_number,pk"`

	InventoryReceiptItems []InventoryReceiptLine `bun:"rel:has-many,join:receipt_number=receipt_number"`
}
