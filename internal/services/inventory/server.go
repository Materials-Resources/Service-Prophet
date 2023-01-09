package inventory

import (
	"context"
	inventoryV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/inventory/v1"
	"github.com/materials-resources/Service-Prophet/internal/database/models"
	"github.com/uptrace/bun"
)

type Server struct {
	inventoryV1.UnimplementedInventoryServiceServer
	DB *bun.DB
}

func (s *Server) GetProduct(ctx context.Context, request *inventoryV1.GetProductRequest) (*inventoryV1.GetProductResponse, error) {

	product := new(models.InvMast)
	err := s.DB.NewSelect().Model(product).Where("inv_mast_uid = ?", request.GetInvMastId()).Scan(ctx)

	response := &inventoryV1.GetProductResponse{
		ItemId:       product.ItemId,
		ItemDesc:     product.ItemDesc,
		ExtendedDesc: product.ExtendedDesc,
		Price1:       product.Price1,
	}
	return response, err
}

func (s *Server) GetProductsByGroup(ctx context.Context, request *inventoryV1.GetProductsByGroupRequest) (*inventoryV1.GetProductsByGroupResponse, error) {

	var products []models.InvLoc
	err := s.DB.NewSelect().Model(&products).
		Where("product_group_id = ?", request.GetProductGroup()).
		Relation("InvMast").
		Limit(1000).OrderExpr("inv_mast_uid ASC").Offset(int(request.GetPage())).
		Scan(ctx)

	response := &inventoryV1.GetProductsByGroupResponse{}

	for _, product := range products {
		response.Products = append(response.Products, &inventoryV1.GetProductsByGroupResponseProduct{
			ItemId:       product.InvMast.ItemId,
			ItemDesc:     product.InvMast.ItemDesc,
			ExtendedDesc: product.InvMast.ExtendedDesc,
			Price1:       product.PriceOne,
		})
	}
	return response, err
}

func (s *Server) GetReceivingReport(ctx context.Context, request *inventoryV1.GetReceivingReportRequest) (*inventoryV1.GetReceivingReportResponse, error) {

	//TODO finish allocated orders
	report := new(models.InventoryReceipt)

	s.DB.NewSelect().Model(report).
		Relation("InventoryReceiptItems").
		Relation("InventoryReceiptItems.OrderTransactions ").
		Where("receipt_number = ?", request.GetReceiptNumber()).
		Scan(ctx)

	response := &inventoryV1.GetReceivingReportResponse{
		ReceiptNumber: report.ReceiptNumber,
	}

	for _, reportLine := range report.InventoryReceiptItems {
		item := &inventoryV1.GetReceivingReportResponse_Item{
			Quantity:      reportLine.UnitQuantity,
			UnitOfMeasure: reportLine.UnitOfMeasure,
			InvMastId:     reportLine.InvMastUid,
		}

		for _, order := range item.AllocatedOrders {

			item.AllocatedOrders = append(item.AllocatedOrders, &inventoryV1.GetReceivingReportResponse_Item_AllocatedOrders{
				DocumentNo:   order.DocumentNo,
				QtyAllocated: order.QtyAllocated,
			})
		}

		response.Items = append(response.Items, item)

	}

	return response, nil
}
