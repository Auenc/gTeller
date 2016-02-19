package discounts

import (
	"errors"
)

const (
	DISCOUNT_TYPE_PERCENTAGE = 1
	DISCOUNT_TYPE_NUMERICAL  = 2
)

type ParseableDiscount struct {
	ID              string
	Type            int
	Percent         int
	Numerical       float64
	Code            string
	RequirementData string
	RequirementType string
}

//Discount is an object that can be applied to an item, reducing it's price.
type Discount struct {
	ID          string
	Type        int
	Percent     int
	Numerical   float64
	Requirement Condition
	Code        string
}

//Parse parses data held within ParseableDiscount and returns a discount object
func (parse *ParseableDiscount) Parse() Discount {
	req := LoadCondition(parse.RequirementType, parse.RequirementData)
	return Discount{ID: parse.ID, Type: parse.Type, Percent: parse.Percent,
		Numerical: parse.Numerical, Requirement: req, Code: parse.Code}
}

//Parsable returns an instance of ParseableDiscount derived from the Discount object
func (discount *Discount) Parsable() (ParseableDiscount, error) {
	var p ParseableDiscount
	reqData, err := discount.Requirement.Data()
	if err != nil {
		return p, err
	}
	return ParseableDiscount{ID: discount.ID, Type: discount.Type,
		Percent: discount.Percent, Code: discount.Code,
		RequirementData: reqData, RequirementType: discount.Requirement.Type()}, nil
}

func (discount *Discount) Value() interface{} {
	switch discount.Type {
	case DISCOUNT_TYPE_NUMERICAL:
		return discount.Numerical
	case DISCOUNT_TYPE_PERCENTAGE:
		return discount.Percent
	}
	return nil
}

func (discount *Discount) TypeName() string {
	switch discount.Type {
	case DISCOUNT_TYPE_NUMERICAL:
		return "Flat rate"
	case DISCOUNT_TYPE_PERCENTAGE:
		return "Percent"
	}
	return "Unknown type"
}

/**
*	A method that will calculate the discount of a specified price and return the outcome.
*
 */
func (discount *Discount) Calculate(price float64) (float64, error) {
	var calculatedPrice float64
	var err error

	//Checking if condition is met

	switch discount.Type {
	case DISCOUNT_TYPE_NUMERICAL:
		/*	Checking to see if there is a valid numerical value	*/
		if discount.Numerical > 0 {
			/*	Calculate new price	*/
			calculatedPrice = price - discount.Numerical
		} else {
			/*	log the issue and create an error	*/
			err = errors.New("Discount numerical is less than 0!")
		}
		break
	case DISCOUNT_TYPE_PERCENTAGE:
		/*	Checking to see if there is a valid percent value*/
		if discount.Percent > 0 {
			/*	Calculate numeric to subtract */
			val1 := float64(discount.Percent) / 100.00
			val2 := float64(val1) * price
			/*	Calculate new price */
			calculatedPrice = price - val2
		} else {
			/*	Log isse and create an error */
			err = errors.New("Discount percentage is less than 0!")
		}
		break
	default:
		err = errors.New("Unknown discount type: " + string(discount.Type))
		break
	}

	return calculatedPrice, err
}
