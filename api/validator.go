package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/raphaeldiscky/simple-bank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if the currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
