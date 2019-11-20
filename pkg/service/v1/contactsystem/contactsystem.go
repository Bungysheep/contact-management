package contactsystem

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/constant/messagecode"
	"github.com/bungysheep/contact-management/pkg/common/message"
	contactsystemmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	contactsystemrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem"
)

// IContactSystemService - Contact System service interface
type IContactSystemService interface {
	DoRead(context.Context, string) (*contactsystemmodel.ContactSystem, messagemodel.IMessage)
	DoReadAll(context.Context) ([]*contactsystemmodel.ContactSystem, messagemodel.IMessage)
	DoSave(context.Context, *contactsystemmodel.ContactSystem) messagemodel.IMessage
	DoDelete(context.Context, string) messagemodel.IMessage
}

type contactSystemService struct {
	communicationMethodRepo communicationmethodrepository.ICommunicationMethodRepository
	contactRepo             contactrepository.IContactRepository
	contactSystemRepo       contactsystemrepository.IContactSystemRepository
}

// NewContactSystemService - Contact System service implementation
func NewContactSystemService(db *sql.DB) IContactSystemService {
	return &contactSystemService{
		communicationMethodRepo: communicationmethodrepository.NewCommunicationMethodRepository(db),
		contactRepo:             contactrepository.NewContactRepository(db),
		contactSystemRepo:       contactsystemrepository.NewContactSystemRepository(db),
	}
}

func (cntsys *contactSystemService) DoRead(ctx context.Context, contactSystemCode string) (*contactsystemmodel.ContactSystem, messagemodel.IMessage) {
	return cntsys.contactSystemRepo.DoRead(ctx, contactSystemCode)
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context) ([]*contactsystemmodel.ContactSystem, messagemodel.IMessage) {
	return cntsys.contactSystemRepo.DoReadAll(ctx)
}

func (cntsys *contactSystemService) DoSave(ctx context.Context, data *contactsystemmodel.ContactSystem) messagemodel.IMessage {
	if err := data.DoValidate(); err != nil {
		return err
	}

	if err := cntsys.contactSystemRepo.DoUpdate(ctx, data); err != nil {
		if err.Code() == messagecode.NotFound {
			return cntsys.contactSystemRepo.DoInsert(ctx, data)
		}
	}

	return nil
}

func (cntsys *contactSystemService) DoDelete(ctx context.Context, contactSystemCode string) messagemodel.IMessage {
	anyRef, err := cntsys.communicationMethodRepo.AnyReference(ctx, contactSystemCode)
	if err != nil {
		return err
	} else if anyRef {
		return message.FailedDeleteAsReferenceExist("Communication Method")
	}

	anyRef, err = cntsys.contactRepo.AnyReference(ctx, contactSystemCode)
	if err != nil {
		return err
	} else if anyRef {
		return message.FailedDeleteAsReferenceExist("Contact")
	}

	return cntsys.contactSystemRepo.DoDelete(ctx, contactSystemCode)
}
