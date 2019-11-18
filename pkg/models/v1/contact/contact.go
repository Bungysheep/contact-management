package contact

import (
	"github.com/bungysheep/contact-management/pkg/models/v1/audit"
	"github.com/bungysheep/contact-management/pkg/models/v1/message"
	"github.com/bungysheep/contact-management/pkg/models/v1/modelbase"
)

// Contact model
type Contact struct {
	modelbase.ModelBase
	ContactSystemCode string `mandatory:"true" max_length:"16" format:"UPPERCASE"`
	ContactID         int64
	FirstName         string `mandatory:"true" max_length:"16"`
	LastName          string `mandatory:"true" max_length:"16"`
	Status            string `mandatory:"true" max_length:"1" valid_value:"A,I" format:"UPPERCASE"`
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

// DoValidate validates fields
func (c *Contact) DoValidate() message.IMessage {
	return c.DoValidateBase(*c)
}
