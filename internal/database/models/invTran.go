package models

import "github.com/uptrace/bun"

type InvTran struct {
	bun.BaseModel `bun:"table:inv_tran"`

	TransactionNumber float64 `bun:"transaction_number,pk"`
	SubDocumentNo     float64 `bun:"sub_document_no"`
	InvMastUid        int     `bun:"inv_mast_uid"`
	QtyAllocated      float64 `bun:"qty_allocated"`
	DocumentNo        float64 `bun:"document_no"`
}
