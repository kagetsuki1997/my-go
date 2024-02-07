package validator

import (
	"my-go/internal/doge/model"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func New() *Validator {
	v := validator.New()

	_ = v.RegisterValidation("type", validateDogeType)
	_ = v.RegisterValidation("name", validateDogeName)

	return &Validator{
		validator: v,
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func validateDogeType(fl validator.FieldLevel) bool {
	allowValues := []model.DogeType{model.DogeTypeDoge, model.DogeTypeDogeBuffed, model.DogeTypeDogeCry, model.DogeTypeDogeKachitoritai}

	fieldValue := fl.Field().String()
	for _, allowValue := range allowValues {
		if fieldValue == string(allowValue) {
			return true
		}
	}

	return false
}

func validateDogeName(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	if len(str) < 2 || len(str) > 50 {
		return false
	}

	return true
}
