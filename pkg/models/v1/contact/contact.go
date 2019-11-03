package contact

import "github.com/bungysheep/contact-management/pkg/models/v1/audit"

// Contact model
type Contact struct {
	ContactSystemCode string
	ContactID         int64
	FirstName         string
	LastName          string
	Status            string
	Audit             *audit.Audit
}

// NewContact creates Contact
func NewContact() *Contact {
	return &Contact{
		Audit: &audit.Audit{},
	}
}

// GetContactSystemCode returns Contact System Code
func (c *Contact) GetContactSystemCode() string {
	return c.ContactSystemCode
}

// GetContactID returns Contact ID
func (c *Contact) GetContactID() int64 {
	return c.ContactID
}

// GetFirstName returns First Name
func (c *Contact) GetFirstName() string {
	return c.FirstName
}

// GetLastName returns Last Name
func (c *Contact) GetLastName() string {
	return c.LastName
}

// GetStatus returns Status
func (c *Contact) GetStatus() string {
	return c.Status
}

// GetAudit returns Audit
func (c *Contact) GetAudit() *audit.Audit {
	return c.Audit
}
