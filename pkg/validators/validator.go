package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(data interface{}) []string {
	if err := validate.Struct(data); err != nil {
		var errs []string
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, "field '"+e.Field()+"' is "+e.Tag())
		}
		return errs
	}
	return nil
}
