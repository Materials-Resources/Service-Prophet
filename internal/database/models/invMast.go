package models

import "database/sql"

func (InvMast) TableName() string {
	return "inv_mast"
}

type InvMast struct {
	InvMastUid   int            `gorm:"column:inv_mast_uid;primaryKey"`
	ItemId       string         `gorm:"column:item_id"`
	ItemDesc     string         `gorm:"column:item_desc"`
	Price1       float32        `gorm:"column:price1"`
	Price2       float32        `gorm:"column:price2"`
	Price3       float32        `gorm:"column:price3"`
	ExtendedDesc sql.NullString `gorm:"column:extended_desc"`
}
