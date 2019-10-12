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

func (cntsys *contactSystemService) DoRead(ctx context.Context, req *contactsystemapi.DoReadRequest) (*contactsystemapi.DoReadResponse, error) {
	result, err := cntsys.repo.DoRead(ctx, req.GetContactSystemCode())

	return &contactsystemapi.DoReadResponse{ContactSystem: result}, err
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context, req *contactsystemapi.DoReadAllRequest) (*contactsystemapi.DoReadAllResponse, error) {
	result, err := cntsys.repo.DoReadAll(ctx)

	return &contactsystemapi.DoReadAllResponse{ContactSystems: result}, err
}

func (cntsys *contactSystemService) DoSave(ctx context.Context, req *contactsystemapi.DoSaveRequest) (*contactsystemapi.DoSaveResponse, error) {
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

func (cntsys *contactSystemService) DoDelete(ctx context.Context, req *contactsystemapi.DoDeleteRequest) (*contactsystemapi.DoDeleteResponse, error) {
	err := cntsys.repo.DoDelete(ctx, req.GetContactSystemCode())

	return &contactsystemapi.DoDeleteResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactsystemrepository.IContactSystemRepository, req *contactsystemapi.DoSaveRequest) (*contactsystemapi.DoSaveResponse, error) {
	err := repo.DoInsert(ctx, req.GetContactSystem())

	return &contactsystemapi.DoSaveResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactsystemrepository.IContactSystemRepository, req *contactsystemapi.DoSaveRequest) (*contactsystemapi.DoSaveResponse, error) {
	err := repo.DoUpdate(ctx, req.GetContactSystem())

	return &contactsystemapi.DoSaveResponse{Result: err == nil}, err
}
