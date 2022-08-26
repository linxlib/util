package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

func Translate(err error) string {

	var result = make([]string, 0)

	errors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, err := range errors {
			result = append(result, err.Translate(trans))
		}
		return strings.Join(result, ";")
	}
	return err.Error()

}
