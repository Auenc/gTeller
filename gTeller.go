package gTeller

const (
	ItemURI            = "item/"
	OrderURI           = "order/"
	ShippingTypeURI    = "shipping/type/"
	ShippingDetailsURI = "shipping/details/"
	DiscountURI        = "discount/"
	StatusURI          = "status/"
	EmailTemplateURI   = "email/template/"
	UserURI            = "user/"
	RequirementURI     = "requirement/"
)

//GTeller is an object that holds config values and grants access to methods to
//use the GTeller api
type GTeller struct {
	Host   string
	APIKey string
}

func New(host, apiKey string) GTeller {
	return GTeller{APIKey: apiKey, Host: host}
}

//Items returns a request object that is configured to the items endpoint
func (gt *GTeller) Items() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + ItemURI
	req.method = "POST"
	return req
}

//ShippingTypes returns a request object that is configured to the ShippingTypes endpoint
func (gt *GTeller) ShippingTypes() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + ShippingTypeURI
	req.method = "POST"
	return req
}

//ShippingDetails returns a request object that is configured to the ShippingDetails endpoint
func (gt *GTeller) ShippingDetails() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + ShippingDetailsURI
	req.method = "POST"
	return req
}

//Orders returns a request object that is configured to the Orders endpoint
func (gt *GTeller) Orders() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + OrderURI
	req.method = "POST"
	return req
}

//Discount returns a request object that is configured to the Discount endpoint
func (gt *GTeller) Discount() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + DiscountURI
	req.method = "POST"
	return req
}

//Statuses returns a request object that is configured to the Statuses endpoint
func (gt *GTeller) Statuses() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + StatusURI
	req.method = "POST"
	return req
}

//EmailTemplates returns a request object that is configured to the EmailTemplates endpoint
func (gt *GTeller) EmailTemplates() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + EmailTemplateURI
	req.method = "POST"
	return req
}

//USer returns a Request object that is configured to the User endpoint
func (gt *GTeller) User() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + UserURI
	req.method = "POST"
	return req
}

//Requirements returns a Request object that is configured to the
func (gt *GTeller) Requirements() Request {
	req := Request{Host: gt.Host}
	req.URL = req.URL + RequirementURI
	req.method = "POST"
	return req
}
