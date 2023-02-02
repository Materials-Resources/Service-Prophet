package receiving_repository

import (
	warehouseV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/warehouse/v1"
	"context"
	"fmt"
	"github.com/materials-resources/Service-Prophet/core/database/models"
	"github.com/uptrace/bun"
	"strconv"
)

func GetAllocatedOrdersForReceiptItem(db *bun.DB, ctx context.Context, receiptId float32, invMastUid int32, primaryBins []string) (*warehouseV1.GetAllocatedOrdersForReceiptItemResponse, error) {
	grpcResponse := new(warehouseV1.GetAllocatedOrdersForReceiptItemResponse)
	model := new(models.InventoryReceiptLine)

	query := db.NewSelect().Model(model).
		Relation("OrderTransactions").
		Relation("InvMast.InvLoc").
		Where("inventory_receipt_line.receipt_number = ? AND inventory_receipt_line.inv_mast_uid = ?", receiptId, invMastUid)

	fmt.Println(len(primaryBins))

	if len(primaryBins) > 0 {
		query.Where("inv_mast__inv_loc.primary_bin in (?)", bun.In(primaryBins))
	}

	err := query.Scan(ctx)

	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}

	for _, transaction := range model.OrderTransactions {
		grpcResponse.AllocatedOrders = append(grpcResponse.AllocatedOrders, &warehouseV1.GetAllocatedOrdersForReceiptItemResponse_AllocatedOrder{
			OrderId:      strconv.Itoa(int(transaction.DocumentNo)),
			QtyAllocated: transaction.QtyAllocated,
		})
	}

	return grpcResponse, nil
}
