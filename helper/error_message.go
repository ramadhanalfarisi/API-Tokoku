package helper

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(model interface{}) []string {
	en := en.New()
	uni := ut.New(en,en)
	trans,_ := uni.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	// validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
	// 	return ut.Add("required", "{0} is required, please input a value", true)
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("required", fe.Field())

	// 	return t
	// })

	err := validate.Struct(model)
	var resultError []string
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			resultError = append(resultError, err.Translate(trans))
		}
	}
	return resultError
}
