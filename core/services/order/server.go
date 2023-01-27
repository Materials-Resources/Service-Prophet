package order

import (
	"buf.build/gen/go/materials-resources/prophet/grpc/go/order/v1/orderv1grpc"
	orderV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/order/v1"
	"context"
	"database/sql"
	"github.com/materials-resources/Service-Prophet/core/database/models"
	"github.com/materials-resources/Service-Prophet/core/database/repositories/address"
	"github.com/materials-resources/Service-Prophet/core/database/repositories/order"
	"github.com/uptrace/bun"
	"time"
)

type Server struct {
	orderv1grpc.UnimplementedOrderServiceServer
	DB *bun.DB
}

func (s *Server) GetOrder(ctx context.Context, request *orderV1.GetOrderRequest) (*orderV1.GetOrderResponse, error) {
	order := new(models.Order)

	err := s.DB.NewSelect().Model(order).Where("order_no = ?", request.GetOrderNo()).Relation("OrderLine").Scan(ctx)

	response := &orderV1.GetOrderResponse{
		OrderNo:              order.OrderNo,
		DeliveryInstructions: order.DeliveryInstructions.String,
		PoNo:                 order.PoNo.String,
		ContactId:            order.ContactId.String,
		ShipToName:           order.ShipToName.String,
		ShipToAdd1:           order.ShipToAdd1.String,
		ShipToAdd2:           order.ShipToAdd2.String,
		ShipToCity:           order.ShipToCity.String,
		ShipToState:          order.ShipToState.String,
		ShipToZip:            order.ShipToZip.String,
		ShipToCountry:        order.ShipToCountry.String,
	}

	for _, items := range order.OrderLine {
		response.OrderItems = append(response.OrderItems, &orderV1.GetOrderResponse_OrderItems{
			InvMastId:     items.InvMastUid,
			QtyOrdered:    items.QtyOrdered,
			Unit:          items.UnitOfMeasure.String,
			ExtendedPrice: items.ExtendedPrice.Float64,
		})
	}

	return response, err
}

func (s *Server) CreateOrder(ctx context.Context, request *orderV1.CreateOrderRequest) (*orderV1.CreateOrderResponse, error) {

	orderNo := new(string)
	oeHdrUid := new(int32)

	order.GetNewOrderNo(ctx, s.DB, orderNo)
	order.GetNewOeHdrUid(ctx, s.DB, oeHdrUid)

	shippingAddress := address.GetShippingAddress(ctx, s.DB, request.GetShippingAddressId())

	newOrder := &models.Order{
		OrderNo:              *orderNo,
		OeHdrUid:             *oeHdrUid,
		Approved:             sql.NullString{String: "N"},
		ContactId:            sql.NullString{},
		CompanyId:            "MRS",
		Completed:            sql.NullString{String: "N"},
		CustomerId:           request.GetCustomerId(),
		DeleteFlag:           "N",
		DeliveryInstructions: sql.NullString{},
		DateCreated:          time.Now(),
		DateModified:         time.Now(),
		OrderDate:            sql.NullTime{Time: time.Now()},
		PoNo: sql.NullString{
			String: request.GetPoNo(),
			Valid:  true,
		},
		ShipToName: sql.NullString{
			String: shippingAddress.Name,
			Valid:  true,
		},
		ShipToAdd1: sql.NullString{
			String: shippingAddress.Line1,
			Valid:  true,
		},
		ShipToAdd2: sql.NullString{
			String: shippingAddress.Line2,
			Valid:  true,
		},
		ShipToCity: sql.NullString{
			String: shippingAddress.City,
			Valid:  true,
		},
		ShipToState: sql.NullString{
			String: shippingAddress.State,
			Valid:  true,
		},
		ShipToZip: sql.NullString{
			String: shippingAddress.PostalCode,
			Valid:  true,
		},
		ShipToCountry: sql.NullString{
			String: "",
			Valid:  false,
		},
		OrderLine: nil,
	}

	s.DB.NewInsert().Model(newOrder).Exec(ctx)

	return &orderV1.CreateOrderResponse{
		OrderNo: newOrder.OrderNo,
	}, nil
}
func (s *Server) GetPickTicket(ctx context.Context, request *orderV1.GetPickTicketRequest) (*orderV1.GetPickTicketResponse, error) {

	orderPickTicket := new(models.OrderPickTicket)
	err := s.DB.NewSelect().Model(orderPickTicket).Where("pick_ticket_no = ?", request.GetPickTicketNo()).Scan(ctx)

	response := &orderV1.GetPickTicketResponse{
		PickTicketNo: orderPickTicket.PickTicketNo,
		OrderNo:      orderPickTicket.OrderNo,
	}

	return response, err
}
