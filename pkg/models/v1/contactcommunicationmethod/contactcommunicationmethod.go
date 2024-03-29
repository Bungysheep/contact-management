package contactcommunicationmethod

import (
	"github.com/bungysheep/contact-management/pkg/models/v1/audit"
	"github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethodfield"
	"github.com/bungysheep/contact-management/pkg/models/v1/message"
	"github.com/bungysheep/contact-management/pkg/models/v1/modelbase"
)

// ContactCommunicationMethod model
type ContactCommunicationMethod struct {
	modelbase.ModelBase
	ContactSystemCode               string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	ContactID                       int64
	ContactCommunicationMethodID    int64
	CommunicationMethodCode         string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	CommunicationMethodLabelCode    string `mandatory:"true" max_length:"8" format:"UPPERCASE"`
	CommunicationMethodLabelCaption string `mandatory:"true" max_length:"16"`
	FormatValue                     string `mandatory:"false" max_length:"1024"`
	Status                          string `mandatory:"true" max_length:"1" valid_value:"A,I" format:"UPPERCASE"`
	IsDefault                       bool
	ContactCommunicationMethodField []*contactcommunicationmethodfield.ContactCommunicationMethodField
	Audit                           *audit.Audit
}

// NewContactCommunicationMethod creates Contact Communication Method
func NewContactCommunicationMethod() *ContactCommunicationMethod {
	return &ContactCommunicationMethod{
		ContactCommunicationMethodField: make([]*contactcommunicationmethodfield.ContactCommunicationMethodField, 0),
		Audit:                           &audit.Audit{},
	}
}

// GetContactSystemCode returns Contact System Code
func (ccm *ContactCommunicationMethod) GetContactSystemCode() string {
	return ccm.ContactSystemCode
}

// GetContactID returns Contact ID
func (ccm *ContactCommunicationMethod) GetContactID() int64 {
	return ccm.ContactID
}

// GetContactCommunicationMethodID returns Contact Communication Method ID
func (ccm *ContactCommunicationMethod) GetContactCommunicationMethodID() int64 {
	return ccm.ContactCommunicationMethodID
}

// GetCommunicationMethodCode returns Communication Method Code
func (ccm *ContactCommunicationMethod) GetCommunicationMethodCode() string {
	return ccm.CommunicationMethodCode
}

// GetCommunicationMethodLabelCode returns Communication Method Label Code
func (ccm *ContactCommunicationMethod) GetCommunicationMethodLabelCode() string {
	return ccm.CommunicationMethodLabelCode
}

// GetCommunicationMethodLabelCaption returns Communication Method Label Caption
func (ccm *ContactCommunicationMethod) GetCommunicationMethodLabelCaption() string {
	return ccm.CommunicationMethodLabelCaption
}

// GetFormatValue returns Format Value
func (ccm *ContactCommunicationMethod) GetFormatValue() string {
	return ccm.FormatValue
}

// GetStatus returns Status
func (ccm *ContactCommunicationMethod) GetStatus() string {
	return ccm.Status
}

// GetIsDefault returns Is Default
func (ccm *ContactCommunicationMethod) GetIsDefault() bool {
	return ccm.IsDefault
}

// GetContactCommunicationMethodField returns Contact Communication Method Field
func (ccm *ContactCommunicationMethod) GetContactCommunicationMethodField() []*contactcommunicationmethodfield.ContactCommunicationMethodField {
	return ccm.ContactCommunicationMethodField
}

// GetAudit returns Audit
func (ccm *ContactCommunicationMethod) GetAudit() *audit.Audit {
	return ccm.Audit
}

// DoValidate validates fields
func (ccm *ContactCommunicationMethod) DoValidate() message.IMessage {
	return ccm.DoValidateBase(*ccm)
}
