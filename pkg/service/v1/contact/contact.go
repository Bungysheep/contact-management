package contact

import (
	"context"

	contactapi "github.com/bungysheep/contact-management/pkg/api/v1/contact"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactService struct {
	repo contactrepository.IContactRepository
}

// NewContactService - Contact service implementation
func NewContactService(repo contactrepository.IContactRepository) contactapi.ContactServiceServer {
	return &contactService{repo: repo}
}

func (cnt *contactService) DoRead(ctx context.Context, req *contactapi.DoReadContactRequest) (*contactapi.DoReadContactResponse, error) {
	result, err := cnt.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetContactId())

	return &contactapi.DoReadContactResponse{Contact: result}, err
}

func (cnt *contactService) DoReadAll(ctx context.Context, req *contactapi.DoReadAllContactRequest) (*contactapi.DoReadAllContactResponse, error) {
	result, err := cnt.repo.DoReadAll(ctx, req.GetContactSystemCode())

	return &contactapi.DoReadAllContactResponse{Contact: result}, err
}

func (cnt *contactService) DoSave(ctx context.Context, req *contactapi.DoSaveContactRequest) (*contactapi.DoSaveContactResponse, error) {
	res, err := doUpdate(ctx, cnt.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cnt.repo, req)
			}
		}
	}

	return res, err
}

func (cnt *contactService) DoDelete(ctx context.Context, req *contactapi.DoDeleteContactRequest) (*contactapi.DoDeleteContactResponse, error) {
	err := cnt.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetContactId())

	return &contactapi.DoDeleteContactResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactrepository.IContactRepository, req *contactapi.DoSaveContactRequest) (*contactapi.DoSaveContactResponse, error) {
	err := repo.DoInsert(ctx, req.GetContact())

	return &contactapi.DoSaveContactResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactrepository.IContactRepository, req *contactapi.DoSaveContactRequest) (*contactapi.DoSaveContactResponse, error) {
	err := repo.DoUpdate(ctx, req.GetContact())

	return &contactapi.DoSaveContactResponse{Result: err == nil}, err
}
