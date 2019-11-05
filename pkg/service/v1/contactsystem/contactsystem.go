package contactsystem

import (
	"context"
	"database/sql"

	contactsystemmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	contactsystemrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IContactSystemService - Contact System service interface
type IContactSystemService interface {
	DoRead(context.Context, string) (*contactsystemmodel.ContactSystem, error)
	DoReadAll(context.Context) ([]*contactsystemmodel.ContactSystem, error)
	DoSave(context.Context, *contactsystemmodel.ContactSystem) error
	DoDelete(context.Context, string) error
}

type contactSystemService struct {
	contactSystemRepo contactsystemrepository.IContactSystemRepository
}

// NewContactSystemService - Contact System service implementation
func NewContactSystemService(db *sql.DB) IContactSystemService {
	return &contactSystemService{
		contactSystemRepo: contactsystemrepository.NewContactSystemRepository(db),
	}
}

func (cntsys *contactSystemService) DoRead(ctx context.Context, contactSystemCode string) (*contactsystemmodel.ContactSystem, error) {
	return cntsys.contactSystemRepo.DoRead(ctx, contactSystemCode)
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context) ([]*contactsystemmodel.ContactSystem, error) {
	return cntsys.contactSystemRepo.DoReadAll(ctx)
}

func (cntsys *contactSystemService) DoSave(ctx context.Context, data *contactsystemmodel.ContactSystem) error {
	if err := cntsys.contactSystemRepo.DoUpdate(ctx, data); err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return cntsys.contactSystemRepo.DoInsert(ctx, data)
			}
		}
	}

	return nil
}

func (cntsys *contactSystemService) DoDelete(ctx context.Context, contactSystemCode string) error {
	return cntsys.contactSystemRepo.DoDelete(ctx, contactSystemCode)
}
