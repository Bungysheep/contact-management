package communicationmethodfield

import "github.com/bungysheep/contact-management/pkg/models/v1/audit"

// CommuniationMethodField model
type CommuniationMethodField struct {
	ContactSystemCode       string
	CommunicationMethodCode string
	FieldCode               string
	Caption                 string
	Sequence                int64
	Audit                   *audit.Audit
}

// NewCommunicationMethodField creates Communication Method Field
func NewCommunicationMethodField() *CommuniationMethodField {
	return &CommuniationMethodField{
		Audit: &audit.Audit{},
	}
}

// GetContactSystemCode returns Contact System Code
func (cmf *CommuniationMethodField) GetContactSystemCode() string {
	return cmf.ContactSystemCode
}

// GetCommunicationMethodCode returns Communication Method Code
func (cmf *CommuniationMethodField) GetCommunicationMethodCode() string {
	return cmf.CommunicationMethodCode
}

// GetFieldCode returns Field Code
func (cmf *CommuniationMethodField) GetFieldCode() string {
	return cmf.FieldCode
}

// GetCaption returns Caption
func (cmf *CommuniationMethodField) GetCaption() string {
	return cmf.Caption
}

// GetSequence returns Sequence
func (cmf *CommuniationMethodField) GetSequence() int64 {
	return cmf.Sequence
}

// GetAudit returns Audit
func (cmf *CommuniationMethodField) GetAudit() *audit.Audit {
	return cmf.Audit
}
