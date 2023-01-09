package order

import (
	"context"
	"database/sql"
	orderV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/order/v1"
	"github.com/materials-resources/Service-Prophet/internal/database/models"
	"github.com/uptrace/bun"
	"time"
)

type Server struct {
	orderV1.UnimplementedOrderServiceServer
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

	//TODO finish creat orders and identify how OrderNo is generated
	order := &models.Order{
		OrderNo:              "10000024",
		OeHdrUid:             400000,
		Approved:             sql.NullString{String: "N"},
		ContactId:            sql.NullString{},
		CompanyId:            "MRS",
		Completed:            sql.NullString{String: "N"},
		CustomerId:           100105,
		DeleteFlag:           "N",
		DeliveryInstructions: sql.NullString{},
		DateCreated:          time.Now(),
		DateModified:         time.Now(),
		OrderDate:            sql.NullTime{Time: time.Now()},
		PoNo:                 sql.NullString{},
		ShipToName:           sql.NullString{},
		ShipToAdd1:           sql.NullString{},
		ShipToAdd2:           sql.NullString{},
		ShipToCity:           sql.NullString{},
		ShipToState:          sql.NullString{},
		ShipToZip:            sql.NullString{},
		ShipToCountry:        sql.NullString{},
		OrderLine:            nil,
	}

	s.DB.NewInsert().Model(order).Exec(ctx)

	return nil, nil
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
