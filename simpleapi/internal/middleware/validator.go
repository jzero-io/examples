package middleware

import (
	"net/http"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(r *http.Request, data any) error {
	validate := validator.New()

	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Tag.Get("json")
			if label == "" {
				label = field.Tag.Get("form")
				if label == "" {
					label = field.Tag.Get("path")
				}
			}
		}
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, ve := range err.(validator.ValidationErrors) {
			if trans != nil {
				return errors.Errorf(ve.Translate(trans))
			}
			return ve
		}
	}
	return nil
}
