package contact

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/api/v1/contact"
)

// IContactRepository - Contact repository interface
type IContactRepository interface {
	DoRead(context.Context, string, int64) (*contact.Contact, error)
	DoReadAll(context.Context, string) ([]*contact.Contact, error)
	DoInsert(context.Context, *contact.Contact) error
	DoUpdate(context.Context, *contact.Contact) error
	DoDelete(context.Context, string, int64) error
}

type contactRepository struct {
	db *sql.DB
}

// NewContactRepository - Contact repository implementation
func NewContactRepository(db *sql.DB) IContactRepository {
	return &contactRepository{db: db}
}

func (cm *contactRepository) DoRead(ctx context.Context, contactSystemCode string, contactID int64) (*contact.Contact, error) {
	return nil, nil
}

func (cm *contactRepository) DoReadAll(ctx context.Context, contactSystemCode string) ([]*contact.Contact, error) {
	return nil, nil
}

func (cm *contactRepository) DoInsert(ctx context.Context, data *contact.Contact) error {
	return nil
}

func (cm *contactRepository) DoUpdate(ctx context.Context, data *contact.Contact) error {
	return nil
}

func (cm *contactRepository) DoDelete(ctx context.Context, contactSystemCode string, contactID int64) error {
	return nil
}
