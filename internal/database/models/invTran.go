package models

import "github.com/uptrace/bun"

type InvTran struct {
	bun.BaseModel `bun:"table:inv_tran"`

	TransactionNumber float32 `bun:"transaction_number,pk"`
	SubDocumentNo     int32   `bun:"sub_document_no"`
	InvMastUid        int32   `bun:"inv_mast_uid"`
	QtyAllocated      float32 `bun:"qty_allocated"`
	DocumentNo        float32 `bun:"document_no"`
	TransType         string  `bun:"trans_type"`
}
