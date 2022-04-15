package Validations

import (
	"regexp"

	"github.com/bykovme/gotrans"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

/*
	Set your own validation rules here to accessable from your controller
*/

func RequiredRule() validation.Rule {
	return validation.Required.Error(gotrans.T("required"))
}

func MaxCharacter(maxCharNo int) validation.Rule {
	return validation.Length(0, maxCharNo).Error(gotrans.T("max_character"))
}

func Email() validation.Rule {
	return is.Email.Error(gotrans.T("email"))
}

func NumericValue() validation.Rule {
	return validation.Match(regexp.MustCompile("^[0-9]")).Error(gotrans.T("numeric"))
}

func EnumValues(key string, values ...interface{}) validation.Rule {
	return validation.In(values...).Error(gotrans.T(key))
}

func MaxValue(value int) validation.Rule {
	return validation.Max(value).Error(gotrans.T("max_value"))
}
