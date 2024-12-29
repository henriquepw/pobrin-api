package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/henriquepw/pobrin-api/pkg/errors"
)

var validate *validator.Validate

func GetValidate() *validator.Validate {
	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())

		validate.RegisterAlias("cnpj", "numeric,len=14")
		validate.RegisterAlias("cpf", "numeric,len=11")
		validate.RegisterAlias("phone", "numeric,len=11")
		validate.RegisterAlias("id", "uppercase,alphanum")
		validate.RegisterAlias("profile", "eq=TEST|eq=DEVELOPMENT|eq=STAGING|eq=PRODUCTION")
		validate.RegisterAlias("logFormat", "eq=DEFAULT|eq=JSON")

		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return validate
}

func getTagError(tag, param string) string {
	switch tag {
	case "required":
		return "Campo obrigátorio"
	case "email":
		return "Email inválido"
	case "cpf":
		return "CPF inválido"
	case "cnpj":
		return "CNPJ inválido"
	case "date":
		return "Data inválida"
	case "phone":
		return "Telefone inválido"
	case "len":
		return "Deve ter exatamente" + param + " caracteres"
	case "max":
		return "Deve ter no máximo " + param + " caracteres"
	case "min":
		return "Deve ter no minímo " + param + " caracteres"
	case "lte":
		return "Deve ser menor ou igual a " + param
	case "lt":
		return "Deve ser menor que " + param
	case "gte":
		return "Deve ser maior ou igual a " + param
	case "gt":
		return "Deve ser maior que " + param
	case "id":
		return "id inválido"
	case "profile":
		return "profile inválido"
	case "logFormat":
		return "log format inválido"
	}

	return "Campo inválido"
}

func Check[T any](val T) error {
	v := GetValidate()
	err := v.Struct(val)

	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	e := make(map[string]string)
	for _, field := range err.(validator.ValidationErrors) {
		name := strings.Join(strings.Split(field.Namespace(), ".")[1:], ".")
		e[name] = getTagError(field.Tag(), field.Param())
	}

	return errors.InvalidRequestData(e)
}
