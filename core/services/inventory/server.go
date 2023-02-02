package inventory

import (
	"buf.build/gen/go/materials-resources/prophet/grpc/go/inventory/v1/inventoryv1grpc"
	inventoryV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/inventory/v1"
	"context"
	"github.com/materials-resources/Service-Prophet/core/database/models"
	"github.com/materials-resources/Service-Prophet/core/database/repositories/catalog"
	dbInventory "github.com/materials-resources/Service-Prophet/core/database/repositories/db-inventory"
	"github.com/uptrace/bun"
	"log"
)

type Server struct {
	inventoryv1grpc.UnimplementedInventoryServiceServer
	DB *bun.DB
}

func (s *Server) GetProduct(ctx context.Context, request *inventoryV1.GetProductRequest) (*inventoryV1.GetProductResponse, error) {

	product, err := dbInventory.GetProductByUid(s.DB, ctx, request.GetInvMastUid())
	if err != nil {
		return nil, err
	}
	return product.ToGrpcResponse(), err
}

func (s *Server) GetProductsByGroup(ctx context.Context, request *inventoryV1.GetProductsByGroupRequest) (*inventoryV1.GetProductsByGroupResponse, error) {

	response := &inventoryV1.GetProductsByGroupResponse{}

	results, err := catalog.GetProductPageAfter(ctx, s.DB, request.GetProductGroup(), request.GetInvMastUid())

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, product := range results.Products {
		response.Products = append(response.Products, &inventoryV1.GetProductsByGroupResponseProduct{
			InvMastUid:     product.InvMastUid,
			ItemId:         product.ItemId,
			ItemDesc:       product.ItemDesc,
			ExtendedDesc:   product.ExtendedDesc,
			ProductGroupId: product.ProductGroupId,
			Price1:         0,
		})
	}

	response.Cursor = &inventoryV1.GetProductsByGroupResponse_Cursor{PrevInvMastUid: results.Cursor.Start, NextInvMastUid: results.Cursor.End}

	return response, nil
}

func (s *Server) GetReceivingReport(ctx context.Context, request *inventoryV1.GetReceivingReportRequest) (*inventoryV1.GetReceivingReportResponse, error) {

	report := new(models.InventoryReceipt)

	s.DB.NewSelect().Model(report).
		Relation("InventoryReceiptItems").
		Relation("InventoryReceiptItems.OrderTransactions").
		Where("receipt_number = ?", request.GetReceiptNumber()).
		Scan(ctx)
	response := &inventoryV1.GetReceivingReportResponse{
		ReceiptNumber: report.ReceiptNumber,
	}

	for _, reportLine := range report.InventoryReceiptItems {
		item := &inventoryV1.GetReceivingReportResponse_Item{
			Quantity:      reportLine.UnitQuantity,
			UnitOfMeasure: reportLine.UnitOfMeasure,
			InvMastUid:    reportLine.InvMastUid,
		}

		response.Items = append(response.Items, item)

	}

	return response, nil
}

func (s *Server) UpdateProductBySupplier(ctx context.Context, request *inventoryV1.UpdateProductBySupplierRequest) (response *inventoryV1.UpdateProductBySupplierResponse, err error) {

	dbInventory.UpdateProductByPrimarySupplier(s.DB, ctx)

	return response, err
}

func (s *Server) GetProductGroups(ctx context.Context, request *inventoryV1.GetProductGroupsRequest) (*inventoryV1.GetProductGroupsResponse, error) {

	response := new(inventoryV1.GetProductGroupsResponse)

	productGroups := dbInventory.GetProductGroups(s.DB, ctx)

	for _, productGroup := range productGroups {
		response.ProductGroup = append(response.ProductGroup,
			&inventoryV1.GetProductGroupsResponse_ProductGroup{
				ProductGroupId:   productGroup.ProductGroupId,
				ProductGroupDesc: productGroup.ProductGroupDesc,
			})
	}

	return response, nil
}
