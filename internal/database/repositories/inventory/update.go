package inventory

import (
	"context"
	"fmt"
	"github.com/materials-resources/Service-Prophet/internal/database/models"
	"github.com/uptrace/bun"
)

func UpdateSupplierCost() {

}
func UpdateProductByPrimarySupplier(db *bun.DB, ctx context.Context) {
	inventorySupplier := new([]models.InventorySupplier)
	db.NewSelect().
		Model(inventorySupplier).
		Relation("InventorySupplierXLoc").
		Relation("InvMast").
		Relation("InvMast.InvLoc").
		Where("supplier_id = ? and inventory_supplier_x_loc.primary_supplier = ?", 103687, "Y").
		Scan(ctx)

	for _, item := range *inventorySupplier {
		fmt.Println(item.InvMast.ItemDesc)
	}
}
