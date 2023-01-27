package catalog

type ProductResults struct {
	Cursor   *ProductResultsCursor
	Products []*ProductResultsProduct
}

type ProductResultsCursor struct {
	Start int32
	End   int32
}

type ProductResultsProduct struct {
	InvMastUid     int32
	ItemId         string
	ItemDesc       string
	ExtendedDesc   string
	ProductGroupId string
}
