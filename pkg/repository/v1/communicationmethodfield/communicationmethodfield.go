package communicationmethodfield

import (
	"context"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
)

// ICommunicationMethodFieldRepository - Communication Method Field repository interface
type ICommunicationMethodFieldRepository interface {
	DoRead(context.Context, string, string, string) (*communicationmethodfield.CommunicationMethodField, error)
	DoReadAll(context.Context) ([]*communicationmethodfield.CommunicationMethodField, error)
	DoInsert(context.Context, *communicationmethodfield.CommunicationMethodField) error
	DoUpdate(context.Context, *communicationmethodfield.CommunicationMethodField) error
	DoDelete(context.Context, string, string, string) error
}