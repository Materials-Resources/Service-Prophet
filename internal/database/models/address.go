package models

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type Address struct {
	bun.BaseModel `bun:"table:address"`

	Id                  float64         `bun:"id,pk"`
	Name                string          `bun:"name"`
	MailAddress1        sql.NullString  `bun:"mail_address1"`
	MailAddress2        sql.NullString  `bun:"mail_address2"`
	MailCity            sql.NullString  `bun:"mail_city"`
	MailState           sql.NullString  `bun:"mail_state"`
	MailPostalCode      sql.NullString  `bun:"mail_postal_code"`
	MailCountry         sql.NullString  `bun:"mail_country"`
	PhysAddress1        sql.NullString  `bun:"phys_address1"`
	PhysAddress2        sql.NullString  `bun:"phys_address2"`
	PhysCity            sql.NullString  `bun:"phys_city"`
	PhysState           sql.NullString  `bun:"phys_state"`
	PhysPostalCode      sql.NullString  `bun:"phys_postal_code"`
	PhysCountry         sql.NullString  `bun:"phys_country"`
	CentralPhoneNo      sql.NullString  `bun:"central_phone_number"`
	CentralFaxNo        sql.NullString  `bun:"central_fax_number"`
	CorpAddressId       sql.NullFloat64 `bun:"corp_address_id"`
	ShippingAddressFlag sql.NullString  `bun:"shipping_address"`
}
