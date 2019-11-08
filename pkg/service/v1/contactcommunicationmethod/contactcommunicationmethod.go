package contactcommunicationmethod

import (
	"context"
	"database/sql"

	contactcommunicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethod"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
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
	contactRepo                    contactrepository.IContactRepository
	contactCommunicationMethodRepo contactcommunicationmethodrepository.IContactCommunicationMethodRepository
}

// NewContactCommunicationMethodService - Contact Communication Method service implementation
func NewContactCommunicationMethodService(db *sql.DB) IContactCommunicationMethodService {
	return &contactcommunicationmethodService{
		contactRepo:                    contactrepository.NewContactRepository(db),
		contactCommunicationMethodRepo: contactcommunicationmethodrepository.NewContactCommunicationMethodRepository(db),
	}
}

func (cmm *contactcommunicationmethodService) DoRead(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) (*contactcommunicationmethodmodel.ContactCommunicationMethod, error) {
	return cmm.contactCommunicationMethodRepo.DoRead(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
}

func (cmm *contactcommunicationmethodService) DoReadAll(ctx context.Context, contactSystemCode string, contactID int64) ([]*contactcommunicationmethodmodel.ContactCommunicationMethod, error) {
	return cmm.contactCommunicationMethodRepo.DoReadAll(ctx, contactSystemCode, contactID)
}

func (cmm *contactcommunicationmethodService) DoSave(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) error {
	if err := cmm.DoValidate(ctx, data); err != nil {
		return err
	}

	if err := cmm.contactCommunicationMethodRepo.DoUpdate(ctx, data); err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return cmm.contactCommunicationMethodRepo.DoInsert(ctx, data)
			}
		}
	}

	return nil
}

func (cmm *contactcommunicationmethodService) DoDelete(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) error {
	return cmm.contactCommunicationMethodRepo.DoDelete(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
}

func (cmm *contactcommunicationmethodService) DoValidate(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) error {
	if _, err := cmm.contactRepo.DoRead(ctx, data.GetContactSystemCode(), data.GetContactID()); err != nil {
		return err
	}

	return nil
}
