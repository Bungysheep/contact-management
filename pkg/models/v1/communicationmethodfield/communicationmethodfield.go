package communicationmethodfield

import (
	"github.com/bungysheep/contact-management/pkg/models/v1/audit"
	"github.com/bungysheep/contact-management/pkg/models/v1/message"
	"github.com/bungysheep/contact-management/pkg/models/v1/modelbase"
)

// CommunicationMethodField model
type CommunicationMethodField struct {
	modelbase.ModelBase
	ContactSystemCode       string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	CommunicationMethodCode string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	FieldCode               string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	Caption                 string `mandatory:"true" max_length:"32"`
	Sequence                int64
	Audit                   *audit.Audit
}

// NewCommunicationMethodField creates Communication Method Field
func NewCommunicationMethodField() *CommunicationMethodField {
	return &CommunicationMethodField{
		Audit: &audit.Audit{},
	}
}

// GetContactSystemCode returns Contact System Code
func (cmf *CommunicationMethodField) GetContactSystemCode() string {
	return cmf.ContactSystemCode
}

// GetCommunicationMethodCode returns Communication Method Code
func (cmf *CommunicationMethodField) GetCommunicationMethodCode() string {
	return cmf.CommunicationMethodCode
}

// GetFieldCode returns Field Code
func (cmf *CommunicationMethodField) GetFieldCode() string {
	return cmf.FieldCode
}

// GetCaption returns Caption
func (cmf *CommunicationMethodField) GetCaption() string {
	return cmf.Caption
}

// GetSequence returns Sequence
func (cmf *CommunicationMethodField) GetSequence() int64 {
	return cmf.Sequence
}

// GetAudit returns Audit
func (cmf *CommunicationMethodField) GetAudit() *audit.Audit {
	return cmf.Audit
}

// DoValidate validates fields
func (cmf *CommunicationMethodField) DoValidate() message.IMessage {
	return cmf.DoValidateBase(*cmf)
}
