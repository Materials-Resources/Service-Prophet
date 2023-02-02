package warehouse

import (
	"buf.build/gen/go/materials-resources/prophet/grpc/go/warehouse/v1/warehousev1grpc"
	warehouseV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/warehouse/v1"
	"context"
	"fmt"
	receivingRepository "github.com/materials-resources/Service-Prophet/core/database/repositories/receiving"
	"github.com/uptrace/bun"
	"google.golang.org/grpc/status"
)

type Server struct {
	warehousev1grpc.UnimplementedWarehouseServiceServer
	DB *bun.DB
}

func (s *Server) GetAllocatedOrdersForReceiptItem(ctx context.Context, req *warehouseV1.GetAllocatedOrdersForReceiptItemRequest) (*warehouseV1.GetAllocatedOrdersForReceiptItemResponse, error) {
	fmt.Println(req.GetProductPrimaryBins())
	res, err := receivingRepository.GetAllocatedOrdersForReceiptItem(s.DB, ctx, req.GetReceiptId(), req.GetInvMastUid(), req.GetProductPrimaryBins())
	if err != nil {
		return res, status.Error(5, err.Error())
	}
	return res, nil
}
