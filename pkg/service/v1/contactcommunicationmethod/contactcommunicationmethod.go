package contactcommunicationmethod

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	contactcommunicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethod"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	communicationmethodlabelrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodlabel"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
	contactcommunicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethodfield"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IContactCommunicationMethodService - Contact Communication Method service interface
type IContactCommunicationMethodService interface {
	DoRead(context.Context, string, int64, int64) (*contactcommunicationmethodmodel.ContactCommunicationMethod, error)
	DoReadAll(context.Context, string, int64) ([]*contactcommunicationmethodmodel.ContactCommunicationMethod, error)
	DoSave(context.Context, *contactcommunicationmethodmodel.ContactCommunicationMethod) error
	DoDelete(context.Context, string, int64, int64) error
}

type contactcommunicationmethodService struct {
	communicationMethodRepo             communicationmethodrepository.ICommunicationMethodRepository
	communicationMethodLabelRepo        communicationmethodlabelrepository.ICommunicationMethodLabelRepository
	contactRepo                         contactrepository.IContactRepository
	contactCommunicationMethodRepo      contactcommunicationmethodrepository.IContactCommunicationMethodRepository
	contactCommunicationMethodFieldRepo contactcommunicationmethodfieldrepository.IContactCommunicationMethodFieldRepository
}

// NewContactCommunicationMethodService - Contact Communication Method service implementation
func NewContactCommunicationMethodService(db *sql.DB) IContactCommunicationMethodService {
	return &contactcommunicationmethodService{
		communicationMethodRepo:             communicationmethodrepository.NewCommunicationMethodRepository(db),
		communicationMethodLabelRepo:        communicationmethodlabelrepository.NewCommunicationMethodLabelRepository(db),
		contactRepo:                         contactrepository.NewContactRepository(db),
		contactCommunicationMethodRepo:      contactcommunicationmethodrepository.NewContactCommunicationMethodRepository(db),
		contactCommunicationMethodFieldRepo: contactcommunicationmethodfieldrepository.NewContactCommunicationMethodFieldRepository(db),
	}
}

func (cmm *contactcommunicationmethodService) DoRead(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) (*contactcommunicationmethodmodel.ContactCommunicationMethod, error) {
	result, err := cmm.contactCommunicationMethodRepo.DoRead(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return nil, err
	}

	fieldResult, err := cmm.contactCommunicationMethodFieldRepo.DoRead(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return result, err
	}

	result.ContactCommunicationMethodField = fieldResult

	return result, nil
}

func (cmm *contactcommunicationmethodService) DoReadAll(ctx context.Context, contactSystemCode string, contactID int64) ([]*contactcommunicationmethodmodel.ContactCommunicationMethod, error) {
	result, err := cmm.contactCommunicationMethodRepo.DoReadAll(ctx, contactSystemCode, contactID)
	if err != nil {
		return result, err
	}

	for _, item := range result {
		fieldResult, err := cmm.contactCommunicationMethodFieldRepo.DoRead(ctx, contactSystemCode, contactID, item.GetContactCommunicationMethodID())
		if err != nil {
			return result, err
		}

		item.ContactCommunicationMethodField = fieldResult
	}

	return result, nil
}

func (cmm *contactcommunicationmethodService) DoSave(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) error {
	if err := cmm.DoValidate(ctx, data); err != nil {
		return err
	}

	if err := cmm.DoUpdateContactCommunicationMethod(ctx, data); err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return cmm.DoInsertContactCommunicationMethod(ctx, data)
			}
		}
	}

	return nil
}

func (cmm *contactcommunicationmethodService) DoUpdateContactCommunicationMethod(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) error {
	if err := cmm.contactCommunicationMethodRepo.DoUpdate(ctx, data); err != nil {
		return err
	}

	for _, item := range data.GetContactCommunicationMethodField() {
		if err := cmm.contactCommunicationMethodFieldRepo.DoUpdate(ctx, item); err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() == codes.NotFound {
					return cmm.contactCommunicationMethodFieldRepo.DoInsert(ctx, item)
				}
			}
		}
	}

	return nil
}

func (cmm *contactcommunicationmethodService) DoInsertContactCommunicationMethod(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) error {
	if err := cmm.contactCommunicationMethodRepo.DoInsert(ctx, data); err != nil {
		return err
	}

	for _, item := range data.GetContactCommunicationMethodField() {
		if err := cmm.contactCommunicationMethodFieldRepo.DoInsert(ctx, item); err != nil {
			return err
		}
	}

	return nil
}

func (cmm *contactcommunicationmethodService) DoDelete(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) error {
	data, err := cmm.contactCommunicationMethodRepo.DoRead(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return err
	} else if data.GetIsDefault() {
		return status.Errorf(codes.Unknown, message.UnableDeleteDefault("Contact Communication Method"))
	}

	if err := cmm.contactCommunicationMethodFieldRepo.DoDelete(ctx, contactSystemCode, contactID, contactCommunicationMethodID); err != nil {
		return err
	}

	return cmm.contactCommunicationMethodRepo.DoDelete(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
}

func (cmm *contactcommunicationmethodService) DoValidate(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) error {
	if _, err := cmm.contactRepo.DoRead(ctx, data.GetContactSystemCode(), data.GetContactID()); err != nil {
		return err
	}

	if _, err := cmm.communicationMethodRepo.DoRead(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode()); err != nil {
		return err
	}

	if _, err := cmm.communicationMethodLabelRepo.DoRead(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetCommunicationMethodLabelCode()); err != nil {
		return err
	}

	return nil
}
