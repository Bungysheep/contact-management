package communicationmethod

import (
	"context"

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
