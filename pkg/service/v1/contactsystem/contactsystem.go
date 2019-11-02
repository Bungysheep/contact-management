package contactsystem

import (
	"context"

	contactsystemapi "github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	contactsystemrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactsystem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactSystemService struct {
	repo contactsystemrepository.IContactSystemRepository
}

// NewContactSystemService - Contact System service implementation
func NewContactSystemService(repo contactsystemrepository.IContactSystemRepository) contactsystemapi.ContactSystemServiceServer {
	return &contactSystemService{repo: repo}
}

func (cntsys *contactSystemService) DoRead(ctx context.Context, req *contactsystemapi.DoReadContactSystemRequest) (*contactsystemapi.DoReadContactSystemResponse, error) {
	result, err := cntsys.repo.DoRead(ctx, req.GetContactSystemCode())

	return &contactsystemapi.DoReadContactSystemResponse{ContactSystem: result}, err
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context, req *contactsystemapi.DoReadAllContactSystemRequest) (*contactsystemapi.DoReadAllContactSystemResponse, error) {
	result, err := cntsys.repo.DoReadAll(ctx)

	return &contactsystemapi.DoReadAllContactSystemResponse{ContactSystems: result}, err
}

func (cntsys *contactSystemService) DoSave(ctx context.Context, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	res, err := doUpdate(ctx, cntsys.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cntsys.repo, req)
			}
		}
	}

	return res, err
}

func (cntsys *contactSystemService) DoDelete(ctx context.Context, req *contactsystemapi.DoDeleteContactSystemRequest) (*contactsystemapi.DoDeleteContactSystemResponse, error) {
	err := cntsys.repo.DoDelete(ctx, req.GetContactSystemCode())

	return &contactsystemapi.DoDeleteContactSystemResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactsystemrepository.IContactSystemRepository, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	err := repo.DoInsert(ctx, req.GetContactSystem())

	return &contactsystemapi.DoSaveContactSystemResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactsystemrepository.IContactSystemRepository, req *contactsystemapi.DoSaveContactSystemRequest) (*contactsystemapi.DoSaveContactSystemResponse, error) {
	err := repo.DoUpdate(ctx, req.GetContactSystem())

	return &contactsystemapi.DoSaveContactSystemResponse{Result: err == nil}, err
}
