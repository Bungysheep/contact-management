package contactsystem

import (
	"reflect"
	"strings"

	"github.com/bungysheep/contact-management/pkg/common/constant/stringformat"
	"github.com/bungysheep/contact-management/pkg/models/v1/audit"
	"github.com/bungysheep/contact-management/pkg/models/v1/modelbase"
)

// ContactSystem model
type ContactSystem struct {
	modelbase.ModelBase
	ContactSystemCode string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	Description       string `mandatory:"true" max_length:"32"`
	Details           string `mandatory:"false" max_length:"255"`
	Status            string `mandatory:"true" max_length:"1" valid_value:"A,I" format:"UPPERCASE"`
	Audit             *audit.Audit
}

// NewContactSystem creates Contact System
func NewContactSystem() *ContactSystem {
	return &ContactSystem{
		Audit: audit.NewAudit(),
	}
}

// GetContactSystemCode returns Contact System Code
func (cs *ContactSystem) GetContactSystemCode() string {
	value := cs.ContactSystemCode
	field, _ := reflect.TypeOf(*cs).FieldByName("ContactSystemCode")

	if field.Tag.Get("format") == stringformat.Uppercase {
		value = strings.ToUpper(value)
	}

	return value
}

// GetDescription returns Description
func (cs *ContactSystem) GetDescription() string {
	return cs.Description
}

// GetDetails returns Details
func (cs *ContactSystem) GetDetails() string {
	return cs.Details
}

// GetStatus returns Status
func (cs *ContactSystem) GetStatus() string {
	value := cs.Status
	field, _ := reflect.TypeOf(*cs).FieldByName("Status")

	if field.Tag.Get("format") == stringformat.Uppercase {
		value = strings.ToUpper(value)
	}

	return value
}

// GetAudit returns Audit
func (cs *ContactSystem) GetAudit() *audit.Audit {
	return cs.Audit
}

// DoValidate validates fields
func (cs *ContactSystem) DoValidate() bool {
	return cs.DoValidateBase(*cs)
}
