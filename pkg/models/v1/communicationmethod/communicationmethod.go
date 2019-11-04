package communicationmethod

import "github.com/bungysheep/contact-management/pkg/models/v1/audit"

// CommunicationMethod model
type CommunicationMethod struct {
	ContactSystemCode       string
	CommunicationMethodCode string
	Description             string
	Details                 string
	Status                  string
	FormatField             string
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
