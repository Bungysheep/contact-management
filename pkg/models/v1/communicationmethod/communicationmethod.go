package communicationmethod

import "github.com/bungysheep/contact-management/pkg/models/v1/audit"

// CommuniationMethod model
type CommuniationMethod struct {
	ContactSystemCode       string
	CommunicationMethodCode string
	Description             string
	Details                 string
	Status                  string
	FormatField             string
	Audit                   *audit.Audit
}

// NewCommunicationMethod creates Communication Method
func NewCommunicationMethod() *CommuniationMethod {
	return &CommuniationMethod{
		Audit: &audit.Audit{},
	}
}

// GetContactSystemCode returns Contact System Code
func (cmf *CommuniationMethod) GetContactSystemCode() string {
	return cmf.ContactSystemCode
}

// GetCommunicationMethodCode returns Communication Method Code
func (cmf *CommuniationMethod) GetCommunicationMethodCode() string {
	return cmf.CommunicationMethodCode
}

// GetDescription returns Description
func (cmf *CommuniationMethod) GetDescription() string {
	return cmf.Description
}

// GetDetails returns Details
func (cmf *CommuniationMethod) GetDetails() string {
	return cmf.Details
}

// GetStatus returns Status
func (cmf *CommuniationMethod) GetStatus() string {
	return cmf.Status
}

// GetFormatField returns FormatField
func (cmf *CommuniationMethod) GetFormatField() string {
	return cmf.FormatField
}

// GetAudit returns Audit
func (cmf *CommuniationMethod) GetAudit() *audit.Audit {
	return cmf.Audit
}
