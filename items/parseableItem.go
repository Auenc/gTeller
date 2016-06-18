package items

type ParseableItem struct {
	ID           string
	Name         string
	Price        string
	Discount     string
	Requirements []string
	ImageURI     string
	Hidden       bool
}
