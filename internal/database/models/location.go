package models

import "time"

type Location struct {
	CompanyId               string  `bun:"company_id"`
	LocationId              float32 `bun:"location_id,pk"`
	DefaultBranchId         string
	DeleteFlag              string
	DateCreated             time.Time
	DateLastModified        time.Time
	LastMaintainedBy        string
	TaxGroupId              string
	LocationName            string
	LotBinIntegration       string
	LocationType            int32
	UseTagsFlag             string
	RfEnabledFlag           string
	WwmsShowBinAllocWarning string
}
