package account

import (
	"context"
	"fmt"
	accountV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/account/v1"
	"github.com/materials-resources/Service-Prophet/internal/database/models"
	"github.com/uptrace/bun"
)

type Server struct {
	accountV1.UnimplementedAccountServiceServer
	DB *bun.DB
}

func (s *Server) GetContact(ctx context.Context, request *accountV1.GetContactRequest) (*accountV1.GetContactResponse, error) {
	contact := new(models.Contact)

	s.DB.NewSelect().Model(contact).Where("id = ?", request.GetId()).Scan(ctx)
	fmt.Println(contact)
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

	return response, nil
}
