package db_inventory

import (
	"context"
	"github.com/materials-resources/Service-Prophet/core/database/models"
	"github.com/uptrace/bun"
)

func getNewInvMastUid(db *bun.DB, ctx context.Context, invMastUid *int) {
	db.NewRaw("SELECT next value for seq_inv_mast").Scan(ctx, &invMastUid)
}

func GetProductGroups(db *bun.DB, ctx context.Context) []models.ProductGroup {
	productGroups := new([]models.ProductGroup)
	db.NewSelect().Column("product_group_id", "product_group_desc").Model(productGroups).Where("delete_flag = 'N'").Scan(ctx)

	return *productGroups
}

func GetProductByUid(db *bun.DB, ctx context.Context, invMastUid int32) (*ProductResponse, error) {
	product := new(models.InvMast)
	err := db.NewSelect().
		Model(product).
		Relation("InvLoc").
		Where("inv_loc.inv_mast_uid = ?", invMastUid).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &ProductResponse{
		InvMastUid:     product.InvMastUid,
		ItemId:         product.ItemId,
		ItemDesc:       product.ItemDesc,
		ExtendedDesc:   product.ExtendedDesc,
		ProductGroupId: product.InvLoc.ProductGroupId,
		Price:          product.Price1,
		QtyAvailable:   product.InvLoc.QtyOnHand,
	}, nil
}
