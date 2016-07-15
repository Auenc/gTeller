package filter

const (
	ConditionEquals = 1
)

type Condition struct {
	Type  int
	Value interface{}
}

func (con *Condition) Valid(n interface{}) bool {
	result := false
	switch con.Type {
	case ConditionEquals:
		result = n == con.Value
		//fmt.Println(n, "==", con.Value, "=", result)
		break
	}
	return result
}
