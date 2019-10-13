package contactcommunicationmethod

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
)

// IContactCommunicationMethodRepository - Contact Communication Method repository interface
type IContactCommunicationMethodRepository interface {
	DoRead(context.Context, string, int64, int64) (*contactcommunicationmethod.ContactCommunicationMethod, error)
	DoReadAll(context.Context) ([]*contactcommunicationmethod.ContactCommunicationMethod, error)
	DoInsert(context.Context, *contactcommunicationmethod.ContactCommunicationMethod) error
	DoUpdate(context.Context, *contactcommunicationmethod.ContactCommunicationMethod) error
	DoDelete(context.Context, string, int64, int64) error
}

type communicationMethodRepository struct {
	db *sql.DB
}

// NewContactCommunicationMethodRepository - Contact Communication Method repository implementation
func NewContactCommunicationMethodRepository(db *sql.DB) IContactCommunicationMethodRepository {
	return &communicationMethodRepository{db: db}
}

func (cm *communicationMethodRepository) DoRead(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) (*contactcommunicationmethod.ContactCommunicationMethod, error) {
	return nil, nil
}

func (cm *communicationMethodRepository) DoReadAll(ctx context.Context) ([]*contactcommunicationmethod.ContactCommunicationMethod, error) {
	return nil, nil
}

func (cm *communicationMethodRepository) DoInsert(ctx context.Context, data *contactcommunicationmethod.ContactCommunicationMethod) error {
	return nil
}

func (cm *communicationMethodRepository) DoUpdate(ctx context.Context, data *contactcommunicationmethod.ContactCommunicationMethod) error {
	return nil
}

func (cm *communicationMethodRepository) DoDelete(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) error {
	return nil
}
