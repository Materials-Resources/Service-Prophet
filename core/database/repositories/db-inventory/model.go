package db_inventory

import inventoryV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/inventory/v1"

type ProductResponse struct {
	InvMastUid     int32
	ItemId         string
	ItemDesc       string
	ExtendedDesc   string
	ProductGroupId string
	Price          float64
	QtyAvailable   float64
}

func (res ProductResponse) ToGrpcResponse() *inventoryV1.GetProductResponse {

	return &inventoryV1.GetProductResponse{
		InvMastUid:     res.InvMastUid,
		ItemId:         res.ItemId,
		ItemDesc:       res.ItemDesc,
		ExtendedDesc:   res.ExtendedDesc,
		ProductGroupId: res.ProductGroupId,
		Price:          res.Price,
		QtyAvailable:   res.QtyAvailable,
	}
}
