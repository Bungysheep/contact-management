package communicationmethodfield

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
)

// ICommunicationMethodFieldRepository - Communication Method Field repository interface
type ICommunicationMethodFieldRepository interface {
	DoRead(context.Context, string, string, string) (*communicationmethodfield.CommunicationMethodField, error)
	DoReadAll(context.Context, string, string) ([]*communicationmethodfield.CommunicationMethodField, error)
	DoInsert(context.Context, *communicationmethodfield.CommunicationMethodField) error
	DoUpdate(context.Context, *communicationmethodfield.CommunicationMethodField) error
	DoDelete(context.Context, string, string, string) error
}

type communicationMethodFieldRepository struct {
	db *sql.DB
}

// NewCommunicationMethodFieldRepository - Communication Method Field repository implementation
func NewCommunicationMethodFieldRepository(db *sql.DB) ICommunicationMethodFieldRepository {
	return &communicationMethodFieldRepository{db: db}
}

func (cmf *communicationMethodFieldRepository) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) (*communicationmethodfield.CommunicationMethodField, error) {
	return nil, nil
}

func (cmf *communicationMethodFieldRepository) DoReadAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) ([]*communicationmethodfield.CommunicationMethodField, error) {
	return nil, nil
}

func (cmf *communicationMethodFieldRepository) DoInsert(ctx context.Context, data *communicationmethodfield.CommunicationMethodField) error {
	return nil
}

func (cmf *communicationMethodFieldRepository) DoUpdate(ctx context.Context, data *communicationmethodfield.CommunicationMethodField) error {
	return nil
}

func (cmf *communicationMethodFieldRepository) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) error {
	return nil
}
