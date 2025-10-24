package request

import (
	"gbase/src/core/helper"
	"strings"

	"github.com/go-playground/validator/v10"

	locale "github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

var Validate *validator.Validate

var uni *ut.UniversalTranslator
var trans ut.Translator

type ErrorBag struct {
	Name    string `json:"name" example:"code"`     // 錯誤欄位名稱
	Message string `json:"message" example:"為必填欄位"` // 錯誤訊息
} //@name ErrorBag

func init() {
	lang := locale.New()
	uni = ut.New(lang, lang)

	trans, _ = uni.GetTranslator("zh_tw")

	Validate = validator.New(validator.WithRequiredStructEnabled())
	translations.RegisterDefaultTranslations(Validate, trans)

}

func FormateErrorBag(err error) ([]ErrorBag, error) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return nil, err
	}

	var ErrorBags []ErrorBag
	for _, e := range err.(validator.ValidationErrors) {
		ErrorBags = append(ErrorBags, ErrorBag{
			Name:    helper.LowerFirst(e.Field()),
			Message: strings.TrimPrefix(e.Translate(trans), e.Field()),
		})

	}
	return ErrorBags, nil

}
