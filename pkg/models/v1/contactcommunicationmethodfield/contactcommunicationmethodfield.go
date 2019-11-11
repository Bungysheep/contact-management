package contactcommunicationmethodfield

import "github.com/bungysheep/contact-management/pkg/models/v1/modelbase"

// ContactCommunicationMethodField model
type ContactCommunicationMethodField struct {
	modelbase.ModelBase
	ContactSystemCode            string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	ContactID                    int64
	ContactCommunicationMethodID int64
	FieldCode                    string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	FieldValue                   string `mandatory:"false" max_length:"32"`
}

// NewContactCommunicationMethodField creates Contact Communication Method Field
func NewContactCommunicationMethodField() *ContactCommunicationMethodField {
	return &ContactCommunicationMethodField{}
}

// GetContactSystemCode returns Contact System Code
func (ccmf *ContactCommunicationMethodField) GetContactSystemCode() string {
	return ccmf.ContactSystemCode
}

// GetContactID returns Contact ID
func (ccmf *ContactCommunicationMethodField) GetContactID() int64 {
	return ccmf.ContactID
}

// GetContactCommunicationMethodID returns Contact Communication Method ID
func (ccmf *ContactCommunicationMethodField) GetContactCommunicationMethodID() int64 {
	return ccmf.ContactCommunicationMethodID
}

// GetFieldCode returns Field Code
func (ccmf *ContactCommunicationMethodField) GetFieldCode() string {
	return ccmf.FieldCode
}

// GetFieldValue returns Field Value
func (ccmf *ContactCommunicationMethodField) GetFieldValue() string {
	return ccmf.FieldValue
}

// DoValidate validates fields
func (ccmf *ContactCommunicationMethodField) DoValidate() bool {
	return ccmf.DoValidateBase(*ccmf)
}
