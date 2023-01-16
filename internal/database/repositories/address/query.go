package address

import (
	"context"
	"github.com/materials-resources/Service-Prophet/internal/database/models"
	"github.com/uptrace/bun"
	"log"
)

func GetShippingAddress(ctx context.Context, db *bun.DB, addressId float64) ShippingAddress {
	address := new(models.Address)
	err := db.NewSelect().Model(address).
		Column("name", "mail_address1", "mail_address2", "mail_city", "mail_state", "mail_postal_code", "mail_country").
		Where("id = ?", addressId).Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return ShippingAddress{
		Name:       address.Name,
		Line1:      address.MailAddress1.String,
		Line2:      address.MailAddress2.String,
		City:       address.MailCity.String,
		State:      address.MailState.String,
		PostalCode: address.MailPostalCode.String,
		Country:    address.MailCountry.String,
	}
}
