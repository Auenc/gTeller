package filter

//Filter is an interface that allows Filter constructs to filter through records
type Filter interface {
	Filter(interface{}) []interface{}
	Valid(interface{}) bool
}
