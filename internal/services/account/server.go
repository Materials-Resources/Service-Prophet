package account

import (
	"context"
	"fmt"
	accountV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/account/v1"
	"github.com/materials-resources/Service-Prophet/internal/database/models"
	"gorm.io/gorm"
)

type Server struct {
	accountV1.UnimplementedAccountServiceServer
	DB *gorm.DB
}

func (s *Server) GetContact(ctx context.Context, request *accountV1.GetContactRequest) (*accountV1.GetContactResponse, error) {
	var contact models.Contacts

	s.DB.Where("id = ?", request.GetId()).Take(&contact)
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

func (s *Server) GetOrders(context.Context, *accountV1.GetOrdersRequest) (*accountV1.GetOrdersResponse, error) {
	var customerWithOrders models.Customer
	s.DB.Model(&models.Customer{}).Where("customer_id = ? AND company_id = ?", 100126, "MRS").Preload("Orders").Take(&customerWithOrders)
	fmt.Println(customerWithOrders)

	return nil, nil
}
