package account

import (
	"buf.build/gen/go/materials-resources/prophet/grpc/go/account/v1/accountv1grpc"
	accountV1 "buf.build/gen/go/materials-resources/prophet/protocolbuffers/go/account/v1"
	"context"
	"github.com/materials-resources/Service-Prophet/core/database/models"
	"github.com/uptrace/bun"
)

type Server struct {
	accountv1grpc.UnimplementedAccountServiceServer
	DB *bun.DB
}

func (s *Server) GetContact(ctx context.Context, request *accountV1.GetContactRequest) (*accountV1.GetContactResponse, error) {
	contact := new(models.Contact)

	s.DB.NewSelect().Model(contact).Where("id = ?", request.GetId()).Scan(ctx)
	response := &accountV1.GetContactResponse{
		Id:            contact.ID,
		Title:         contact.Title.String,
		FirstName:     contact.FirstName,
		MiddleInitial: contact.MiddleInitial.String,
		LastName:      contact.LastName,
		EmailAddress:  contact.EmailAddress.String,
		DirectPhone:   contact.DirectPhone.String,
	}

	return response, nil
}

func (s *Server) GetOrders(ctx context.Context, request *accountV1.GetOrdersRequest) (*accountV1.GetOrdersResponse, error) {
	customer := new(models.Customer)
	s.DB.NewSelect().Model(customer).Where("customer_id = ? AND company_id = ?", 100711, "MRS").Relation("Orders").Scan(ctx)

	response := &accountV1.GetOrdersResponse{}

	for _, orders := range customer.Orders {
		response.Orders = append(response.Orders, &accountV1.GetOrdersResponse_Order{
			OrderNo:    orders.OrderNo,
			OrderDate:  orders.OrderDate.Time.String(),
			Completed:  orders.Completed.String,
			Approved:   orders.Approved.String,
			DeleteFlag: orders.DeleteFlag,
		})
	}

	return response, nil
}
