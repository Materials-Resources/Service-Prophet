package order

import (
	"context"
	"github.com/uptrace/bun"
)

func GetNewOrderNo(ctx context.Context, db *bun.DB, orderNo *string) {
	db.NewRaw("SELECT next value for seq_WO").Scan(ctx, &orderNo)
}

func GetNewOeHdrUid(ctx context.Context, db *bun.DB, oeHdrUid *int32) {
	db.NewRaw("SELECT next value for seq_oe_hdr").Scan(ctx, &oeHdrUid)
}
