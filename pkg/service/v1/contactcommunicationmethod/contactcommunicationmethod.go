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

func (cmm *contactcommunicationmethodService) DoRead(ctx context.Context, req *contactcommunicationmethodapi.DoReadContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoReadContactCommunicationMethodResponse, error) {
	result, err := cmm.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoReadContactCommunicationMethodResponse{ContactCommunicationMethod: result}, err
}

func (cmm *contactcommunicationmethodService) DoReadAll(ctx context.Context, req *contactcommunicationmethodapi.DoReadAllContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoReadAllContactCommunicationMethodResponse, error) {
	result, err := cmm.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetContactId())

	return &contactcommunicationmethodapi.DoReadAllContactCommunicationMethodResponse{ContactCommunicationMethod: result}, err
}

func (cmm *contactcommunicationmethodService) DoSave(ctx context.Context, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
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

func (cmm *contactcommunicationmethodService) DoDelete(ctx context.Context, req *contactcommunicationmethodapi.DoDeleteContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoDeleteContactCommunicationMethodResponse, error) {
	err := cmm.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetContactId(), req.GetContactCommunicationMethodId())

	return &contactcommunicationmethodapi.DoDeleteContactCommunicationMethodResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
	err := repo.DoInsert(ctx, req.GetContactCommunicationMethod())

	return &contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactcommunicationmethodrepository.IContactCommunicationMethodRepository, req *contactcommunicationmethodapi.DoSaveContactCommunicationMethodRequest) (*contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse, error) {
	err := repo.DoUpdate(ctx, req.GetContactCommunicationMethod())

	return &contactcommunicationmethodapi.DoSaveContactCommunicationMethodResponse{Result: err == nil}, err
}
