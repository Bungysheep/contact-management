package contactcommunicationmethod

import (
	"context"

	contactcommunicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactcommunicationmethodService struct {
	repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository
}

// NewContactCommunicationMethodService - Contact Communication Method service implementation
func NewContactCommunicationMethodService(repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository) contactcommunicationmethodapi.ContactCommunicationMethodServiceServer {
	return &contactcommunicationmethodService{repo: repo}
}

func (cmm *contactcommunicationmethodService) DoRead(ctx context.Context, req *contactcommunicationmethodapi.DoReadRequest) (*contactcommunicationmethodapi.DoReadResponse, error) {
	result, err := cmm.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoReadResponse{ContactCommunicationMethod: result}, err
}

func (cmm *contactcommunicationmethodService) DoReadAll(ctx context.Context, req *contactcommunicationmethodapi.DoReadAllRequest) (*contactcommunicationmethodapi.DoReadAllResponse, error) {
	result, err := cmm.repo.DoReadAll(ctx, req.GetContactSystemCode())

	return &contactcommunicationmethodapi.DoReadAllResponse{ContactCommunicationMethod: result}, err
}

func (cmm *contactcommunicationmethodService) DoSave(ctx context.Context, req *contactcommunicationmethodapi.DoSaveRequest) (*contactcommunicationmethodapi.DoSaveResponse, error) {
	res, err := doUpdate(ctx, cmm.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cmm.repo, req)
			}
		}
	}

	return res, err
}

func (cmm *contactcommunicationmethodService) DoDelete(ctx context.Context, req *contactcommunicationmethodapi.DoDeleteRequest) (*contactcommunicationmethodapi.DoDeleteResponse, error) {
	err := cmm.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoDeleteResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository, req *contactcommunicationmethodapi.DoSaveRequest) (*contactcommunicationmethodapi.DoSaveResponse, error) {
	err := repo.DoInsert(ctx, req.GetContactCommunicationMethod())

	return &contactcommunicationmethodapi.DoSaveResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository, req *contactcommunicationmethodapi.DoSaveRequest) (*contactcommunicationmethodapi.DoSaveResponse, error) {
	err := repo.DoUpdate(ctx, req.GetContactCommunicationMethod())

	return &contactcommunicationmethodapi.DoSaveResponse{Result: err == nil}, err
}
