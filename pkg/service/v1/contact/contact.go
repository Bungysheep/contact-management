package contact

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/constant/messagecode"
	contactmodel "github.com/bungysheep/contact-management/pkg/models/v1/contact"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
	contactsystemrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem"
)

// IContactService - Contact service interface
type IContactService interface {
	DoRead(context.Context, string, int64) (*contactmodel.Contact, messagemodel.IMessage)
	DoReadAll(context.Context, string) ([]*contactmodel.Contact, messagemodel.IMessage)
	DoSave(context.Context, *contactmodel.Contact) messagemodel.IMessage
	DoDelete(context.Context, string, int64) messagemodel.IMessage
}

type contactService struct {
	contactRepo           contactrepository.IContactRepository
	contactCommMethodRepo contactcommunicationmethodrepository.IContactCommunicationMethodRepository
	contactSystemRepo     contactsystemrepository.IContactSystemRepository
}

// NewContactService - Contact service implementation
func NewContactService(db *sql.DB) IContactService {
	return &contactService{
		contactRepo:           contactrepository.NewContactRepository(db),
		contactCommMethodRepo: contactcommunicationmethodrepository.NewContactCommunicationMethodRepository(db),
		contactSystemRepo:     contactsystemrepository.NewContactSystemRepository(db),
	}
}

func (cnt *contactService) DoRead(ctx context.Context, contactSystemCode string, contactID int64) (*contactmodel.Contact, messagemodel.IMessage) {
	return cnt.contactRepo.DoRead(ctx, contactSystemCode, contactID)
}

func (cnt *contactService) DoReadAll(ctx context.Context, contactSystemCode string) ([]*contactmodel.Contact, messagemodel.IMessage) {
	return cnt.contactRepo.DoReadAll(ctx, contactSystemCode)
}

func (cnt *contactService) DoSave(ctx context.Context, data *contactmodel.Contact) messagemodel.IMessage {
	if err := data.DoValidate(); err != nil {
		return err
	}

	if err := cnt.DoValidate(ctx, data); err != nil {
		return err
	}

	if err := cnt.contactRepo.DoUpdate(ctx, data); err != nil {
		if err.Code() == messagecode.NotFound {
			return cnt.contactRepo.DoInsert(ctx, data)
		}
	}

	return nil
}

func (cnt *contactService) DoDelete(ctx context.Context, contactSystemCode string, contactID int64) messagemodel.IMessage {
	if err := cnt.contactCommMethodRepo.DoDeleteAll(ctx, contactSystemCode, contactID); err != nil {
		return err
	}

	return cnt.contactRepo.DoDelete(ctx, contactSystemCode, contactID)
}

func (cnt *contactService) DoValidate(ctx context.Context, data *contactmodel.Contact) messagemodel.IMessage {
	if _, err := cnt.contactSystemRepo.DoRead(ctx, data.GetContactSystemCode()); err != nil {
		return err
	}

	return nil
}
