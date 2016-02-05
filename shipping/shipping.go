package shipping

type ShippingType struct {
	ID    string
	Name  string
	Price float64
}

type ShippingDetails struct {
	ID       string
	Type     ShippingType
	Town     string
	Country  string
	Line1    string
	City     string
	PostCode string
	Name     string
	Phone    string
	Tracking string
}

type ByPrice []ShippingType

type ShippingRepository interface {
	AddType(st ShippingType) error
	AddDetail(sd ShippingDetails) error
	Type(id string) (ShippingType, error)
	Detail(id string) (ShippingDetails, error)
	Types() ([]ShippingType, error)
	Details() ([]ShippingDetails, error)
	RemoveType(id string) error
	RemoveDetail(id string) error
	UpdateDetail(ShippingDetails) error
	UpdateType(ShippingType) error
}

func (st ByPrice) Len() int {
	return len(st)
}

func (st ByPrice) Less(i, j int) bool {
	return st[i].Price < st[j].Price
}
func (st ByPrice) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
}
