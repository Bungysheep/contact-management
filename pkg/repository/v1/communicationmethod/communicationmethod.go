package communicationmethod

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
)

// ICommunicationMethodRepository - Communication Method repository interface
type ICommunicationMethodRepository interface {
	DoRead(context.Context, string, string) (*communicationmethod.CommunicationMethod, error)
	DoReadAll(context.Context) ([]*communicationmethod.CommunicationMethod, error)
	DoInsert(context.Context, *communicationmethod.CommunicationMethod) error
	DoUpdate(context.Context, *communicationmethod.CommunicationMethod) error
	DoDelete(context.Context, string, string) error
}

type communicationMethodRepository struct {
	db *sql.DB
}

// NewCommunicationMethodRepository - Communication Method repository implementation
func NewCommunicationMethodRepository(db *sql.DB) ICommunicationMethodRepository {
	return &communicationMethodRepository{db: db}
}

func (cm *communicationMethodRepository) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string) (*communicationmethod.CommunicationMethod, error) {
	return nil, nil
}

func (cm *communicationMethodRepository) DoReadAll(ctx context.Context) ([]*communicationmethod.CommunicationMethod, error) {
	return nil, nil
}

func (cm *communicationMethodRepository) DoInsert(ctx context.Context, data *communicationmethod.CommunicationMethod) error {
	return nil
}

func (cm *communicationMethodRepository) DoUpdate(ctx context.Context, data *communicationmethod.CommunicationMethod) error {
	return nil
}

func (cm *communicationMethodRepository) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string) error {
	return nil
}
