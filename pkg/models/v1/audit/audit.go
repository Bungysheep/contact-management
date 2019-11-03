package audit

import "time"

// Audit model
type Audit struct {
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
