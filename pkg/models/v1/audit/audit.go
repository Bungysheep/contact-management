package audit

import (
	"time"

	"github.com/bungysheep/contact-management/pkg/models/v1/message"
	"github.com/bungysheep/contact-management/pkg/models/v1/modelbase"
)

// Audit model
type Audit struct {
	modelbase.ModelBase
	CreatedAt  time.Time
	ModifiedAt time.Time
	Vers       int64
}

// NewAudit creates Audit
func NewAudit() *Audit {
	return &Audit{}
}

// GetCreatedAt return Created At
func (a *Audit) GetCreatedAt() time.Time {
	return a.CreatedAt
}

// GetModifiedAt return Modified At
func (a *Audit) GetModifiedAt() time.Time {
	return a.ModifiedAt
}

// GetVers return Vers
func (a *Audit) GetVers() int64 {
	return a.Vers
}

// DoValidate validates fields
func (a *Audit) DoValidate() message.IMessage {
	return nil
}
