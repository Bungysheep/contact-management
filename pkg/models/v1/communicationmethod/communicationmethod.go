package communicationmethod

import (
	"github.com/bungysheep/contact-management/pkg/models/v1/audit"
	"github.com/bungysheep/contact-management/pkg/models/v1/modelbase"
)

// CommunicationMethod model
type CommunicationMethod struct {
	modelbase.ModelBase
	ContactSystemCode       string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	CommunicationMethodCode string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	Description             string `mandatory:"true" max_length:"32"`
	Details                 string `mandatory:"false" max_length:"255"`
	Status                  string `mandatory:"true" max_length:"1" valid_value:"A,I" format:"UPPERCASE"`
	FormatField             string `mandatory:"true" max_length:"1024"`
	Audit                   *audit.Audit
}

// NewCommunicationMethod creates Communication Method
func NewCommunicationMethod() *CommunicationMethod {
	return &CommunicationMethod{
		Audit: &audit.Audit{},
	}
}

// GetContactSystemCode returns Contact System Code
func (cmf *CommunicationMethod) GetContactSystemCode() string {
	return cmf.ContactSystemCode
}

// GetCommunicationMethodCode returns Communication Method Code
func (cmf *CommunicationMethod) GetCommunicationMethodCode() string {
	return cmf.CommunicationMethodCode
}

// GetDescription returns Description
func (cmf *CommunicationMethod) GetDescription() string {
	return cmf.Description
}

// GetDetails returns Details
func (cmf *CommunicationMethod) GetDetails() string {
	return cmf.Details
}

// GetStatus returns Status
func (cmf *CommunicationMethod) GetStatus() string {
	return cmf.Status
}

// GetFormatField returns FormatField
func (cmf *CommunicationMethod) GetFormatField() string {
	return cmf.FormatField
}

// GetAudit returns Audit
func (cmf *CommunicationMethod) GetAudit() *audit.Audit {
	return cmf.Audit
}

// DoValidate validates fields
func (cmf *CommunicationMethod) DoValidate() bool {
	return cmf.DoValidateBase(*cmf)
}
