package models

import "github.com/uptrace/bun"

type OrderPickTicket struct {
	bun.BaseModel `bun:"table:oe_pick_ticket"`

	PickTicketNo float64 `bun:"pick_ticket_no,pk"`
	OrderNo      string  `bun:"order_no"`
	Order        Order   `bun:"rel:has-one,join:order_no=order_no"`
}
