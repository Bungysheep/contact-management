package communicationmethodlabel

import "github.com/bungysheep/contact-management/pkg/models/v1/modelbase"

// CommunicationMethodLabel model
type CommunicationMethodLabel struct {
	modelbase.ModelBase
	ContactSystemCode            string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	CommunicationMethodCode      string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	CommunicationMethodLabelCode string `mandatory:"true" max_length:"8" format:"UPPERCASE"`
	Caption                      string `mandatory:"true" max_length:"16"`
}

// NewCommunicationMethodLabel creates Communication Method Label
func NewCommunicationMethodLabel() *CommunicationMethodLabel {
	return &CommunicationMethodLabel{}
}

// GetContactSystemCode returns Contact System Code
func (cml *CommunicationMethodLabel) GetContactSystemCode() string {
	return cml.ContactSystemCode
}

// GetCommunicationMethodCode returns Communication Method Code
func (cml *CommunicationMethodLabel) GetCommunicationMethodCode() string {
	return cml.CommunicationMethodCode
}

// GetCommunicationMethodLabelCode returns Communication Method Label Code
func (cml *CommunicationMethodLabel) GetCommunicationMethodLabelCode() string {
	return cml.CommunicationMethodLabelCode
}

// GetCaption returns Caption
func (cml *CommunicationMethodLabel) GetCaption() string {
	return cml.Caption
}

// DoValidate validates fields
func (cml *CommunicationMethodLabel) DoValidate() error {
	return cml.DoValidateBase(*cml)
}
