package contactsystem

import "github.com/bungysheep/contact-management/pkg/models/v1/audit"

// ContactSystem model
type ContactSystem struct {
	ContactSystemCode string
	Description       string
	Details           string
	Status            string
	Audit             *audit.Audit
}

// NewContactSystem creates Contact System
func NewContactSystem() *ContactSystem {
	return &ContactSystem{
		Audit: &audit.Audit{},
	}
}

// GetContactSystemCode returns Contact System Code
func (cs *ContactSystem) GetContactSystemCode() string {
	return cs.ContactSystemCode
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
	return cs.Status
}

// GetAudit returns Audit
func (cs *ContactSystem) GetAudit() *audit.Audit {
	return cs.Audit
}
