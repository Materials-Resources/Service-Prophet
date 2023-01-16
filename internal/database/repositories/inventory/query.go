package inventory

import (
	"context"
	"github.com/uptrace/bun"
)

func getNewInvMastUid(db *bun.DB, ctx context.Context, invMastUid *int) {
	db.NewRaw("SELECT next value for seq_inv_mast").Scan(ctx, &invMastUid)
}
