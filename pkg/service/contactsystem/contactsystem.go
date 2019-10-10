package contactsystem

import "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"

type contactSystemService struct {
}

// NewContactSystemService - Contact System service implementation
func NewContactSystemService() contactsystem.ContactSystemServiceServer {
	return &contactSystemService{}
}
