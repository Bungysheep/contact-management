package modelbase

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bungysheep/contact-management/pkg/common/message"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
)

// ModelBase model
type ModelBase struct{}

// NewModelBase creates Model Base
func NewModelBase() *ModelBase {
	return &ModelBase{}
}

// DoValidateBase validates fields
func (mb *ModelBase) DoValidateBase(model interface{}) messagemodel.IMessage {
	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		switch field.Type.Kind() {
		case reflect.String:
			value := modelValue.Field(i).String()

			if !mb.IsSpecified(value, field.Tag.Get("mandatory")) {
				return message.MustBeSpecified(field.Name)
			}

			if !mb.IsValidMaxLength(value, field.Tag.Get("max_length")) {
				return message.CannotMoreThanNChars(field.Name, field.Tag.Get("max_length"))
			}

			if !mb.IsValidValue(value, field.Tag.Get("valid_value")) {
				return message.NotValid(field.Name, field.Tag.Get("valid_value"))
			}

		case reflect.Ptr:
			if err := modelValue.Field(i).MethodByName("DoValidate").Call([]reflect.Value{})[0].Interface(); err != nil {
				return err.(messagemodel.IMessage)
			}

		case reflect.Struct:
			switch field.Type.Name() {
			case reflect.TypeOf(time.Time{}).Name():

			}

		case reflect.Slice:
			for j := 0; j < modelValue.Field(i).Len(); j++ {
				switch modelValue.Field(i).Index(j).Kind() {
				case reflect.Ptr:
					if err := modelValue.Field(i).Index(j).MethodByName("DoValidate").Call([]reflect.Value{})[0].Interface(); err != nil {
						return err.(messagemodel.IMessage)
					}
				}
			}
		}
	}

	return nil
}

// IsSpecified validates mandatory
func (mb *ModelBase) IsSpecified(value string, isMandatory string) bool {
	if isMandatory != "" {
		isMandatoryBool, _ := strconv.ParseBool(isMandatory)
		if isMandatoryBool && value == "" {
			return false
		}
	}
	return true
}

// IsValidMaxLength validates max length
func (mb *ModelBase) IsValidMaxLength(value string, maxLength string) bool {
	if maxLength != "" {
		maxLengthInt, _ := strconv.Atoi(maxLength)
		if len(value) > maxLengthInt {
			return false
		}
	}
	return true
}

// IsValidValue validates valid values
func (mb *ModelBase) IsValidValue(value string, validValue string) bool {
	if validValue != "" {
		values := strings.Split(validValue, ",")

		for _, item := range values {
			if item == value {
				return true
			}
		}
		return false
	}
	return true
}
