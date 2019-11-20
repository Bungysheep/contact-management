package communicationmethod

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/constant/messagecode"
	communicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	communicationmethodlabelrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodlabel"
	contactsystemrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem"
)

// ICommunicationMethodService - Communication Method service interface
type ICommunicationMethodService interface {
	DoRead(context.Context, string, string) (*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage)
	DoReadAll(context.Context, string) ([]*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage)
	DoSave(context.Context, *communicationmethodmodel.CommunicationMethod) messagemodel.IMessage
	DoDelete(context.Context, string, string) messagemodel.IMessage
}

type communicationMethodService struct {
	communicationMethodRepo      communicationmethodrepository.ICommunicationMethodRepository
	communicationMethodFieldRepo communicationmethodfieldrepository.ICommunicationMethodFieldRepository
	communicationMethodLabelRepo communicationmethodlabelrepository.ICommunicationMethodLabelRepository
	contactSystemRepo            contactsystemrepository.IContactSystemRepository
}

// NewCommunicationMethodService - Communication Method service implementation
func NewCommunicationMethodService(db *sql.DB) ICommunicationMethodService {
	return &communicationMethodService{
		communicationMethodRepo:      communicationmethodrepository.NewCommunicationMethodRepository(db),
		communicationMethodFieldRepo: communicationmethodfieldrepository.NewCommunicationMethodFieldRepository(db),
		communicationMethodLabelRepo: communicationmethodlabelrepository.NewCommunicationMethodLabelRepository(db),
		contactSystemRepo:            contactsystemrepository.NewContactSystemRepository(db),
	}
}

func (cm *communicationMethodService) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string) (*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage) {
	return cm.communicationMethodRepo.DoRead(ctx, contactSystemCode, communicationMethodCode)
}

func (cm *communicationMethodService) DoReadAll(ctx context.Context, contactSystemCode string) ([]*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage) {
	return cm.communicationMethodRepo.DoReadAll(ctx, contactSystemCode)
}

func (cm *communicationMethodService) DoSave(ctx context.Context, data *communicationmethodmodel.CommunicationMethod) messagemodel.IMessage {
	if err := data.DoValidate(); err != nil {
		return err
	}

	if err := cm.DoValidate(ctx, data); err != nil {
		return err
	}

	if err := cm.communicationMethodRepo.DoUpdate(ctx, data); err != nil {
		if err.Code() == messagecode.NotFound {
			return cm.communicationMethodRepo.DoInsert(ctx, data)
		}
	}

	return nil
}

func (cm *communicationMethodService) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string) messagemodel.IMessage {
	if err := cm.communicationMethodFieldRepo.DoDeleteAll(ctx, contactSystemCode, communicationMethodCode); err != nil {
		return err
	}

	if err := cm.communicationMethodLabelRepo.DoDeleteAll(ctx, contactSystemCode, communicationMethodCode); err != nil {
		return err
	}

	return cm.communicationMethodRepo.DoDelete(ctx, contactSystemCode, communicationMethodCode)
}

func (cm *communicationMethodService) DoValidate(ctx context.Context, data *communicationmethodmodel.CommunicationMethod) messagemodel.IMessage {
	if _, err := cm.contactSystemRepo.DoRead(ctx, data.GetContactSystemCode()); err != nil {
		return err
	}

	return nil
}
