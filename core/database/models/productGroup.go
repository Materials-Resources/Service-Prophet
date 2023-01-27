package models

import (
	"github.com/uptrace/bun"
	"time"
)

type ProductGroup struct {
	bun.BaseModel             `bun:"table:product_group"`
	CompanyId                 string    `bun:"company_id,pk,notnull"`
	ProductGroupId            string    `bun:"product_group_id,pk,notnull"`
	ProductGroupDesc          string    `bun:"product_group_desc,notnull"`
	AssetAccountNo            string    `bun:"asset_account_no,notnull"`
	RevenueAccountNo          string    `bun:"revenue_account_no"`
	CosAccountNo              string    `bun:"cos_account_no"`
	DeleteFlag                string    `bun:"delete_flag,notnull"`
	DateCreated               time.Time `bun:"date_created,notnull"`
	DateLastModified          time.Time `bun:"date_last_modified,notnull"`
	LastMaintainedBy          string    `bun:"last_maintained_by,notnull"`
	ProductGroupUid           int32     `bun:"product_group_uid,unique"`
	EnvironmentalFeeFlag      string    `bun:"environmental_fee_flag,notnull,default:N"`
	AdminFleeFlag             string    `bun:"admin_fee_flag,notnull,default:N"`
	IncludeInSizeAnalysisFlag string    `bun:"include_in_size_analysis_flag,default:Y"`
	PriceRoundingFlag         string    `bun:"price_rounding_flag,notnull,default:N"`
	ApplyMinMarginRulesFlag   string    `bun:"apply_min_margin_rules_flag,notnull,default:N"`
	EnableLineProfitWarning   string    `bun:"enable_line_profit_warning,notnull,default:N"`
	OeSkipProfitCheckUnpriced string    `bun:"oe_skip_profit_check_unpriced,notnull,default:N"`
	OeSkipProfitCheckUncosted string    `bun:"oe_skip_profit_check_uncosted,notnull,default:N"`
	CompressorFlag            string    `bun:"compressor_flag,notnull,default:N"`
}
