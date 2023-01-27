package catalog

import (
	"context"
	"github.com/materials-resources/Service-Prophet/core/database/models"
	"github.com/uptrace/bun"
)

// GetProductPageAfter Get list of 1000 products belonging to a product group starting at desired inv_mast_uid
func GetProductPageAfter(ctx context.Context, db *bun.DB, productGroups []string, invMastUid int32) (ProductResults, error) {
	var products []models.InvMast
	err := db.NewSelect().Model(&products).
		Where("inv_loc.product_group_id in (?) AND inv_mast.inv_mast_uid > ?", bun.In(productGroups), invMastUid).
		Relation("InvLoc").
		Limit(1000).OrderExpr("inv_mast.inv_mast_uid ASC").
		Scan(ctx)
	if err != nil {
		return ProductResults{}, err
	}

	if len(products) < 1 {
		return ProductResults{Cursor: &ProductResultsCursor{
			Start: 0,
		}}, nil
	}

	results := new(ProductResults)
	for _, product := range products {
		results.Products = append(results.Products, &ProductResultsProduct{
			InvMastUid:     product.InvMastUid,
			ItemId:         product.ItemId,
			ItemDesc:       product.ItemDesc,
			ExtendedDesc:   product.ExtendedDesc,
			ProductGroupId: product.InvLoc.ProductGroupId,
		})
	}
	results.Cursor = &ProductResultsCursor{
		Start: products[0].InvMastUid,
		End:   products[len(products)-1].InvMastUid,
	}

	return *results, nil
}

func GetProductPageBefore() {

}
