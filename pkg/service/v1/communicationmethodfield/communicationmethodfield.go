package communicationmethodfield

import (
	"context"

	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodFieldService struct {
	repo communicationmethodfieldrepository.ICommunicationMethodFieldRepository
}

// NewCommunicationMethodFieldService - Communication Method Field service implementation
func NewCommunicationMethodFieldService(repo communicationmethodfieldrepository.ICommunicationMethodFieldRepository) communicationmethodfieldapi.CommunicationMethodFieldServiceServer {
	return &communicationMethodFieldService{repo: repo}
}

func (cmf *communicationMethodFieldService) DoRead(ctx context.Context, req *communicationmethodfieldapi.DoReadCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoReadCommunicationMethodFieldResponse, error) {
	result, err := cmf.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetFieldCode())

	return &communicationmethodfieldapi.DoReadCommunicationMethodFieldResponse{CommunicationMethodField: result}, err
}

func (cmf *communicationMethodFieldService) DoReadAll(ctx context.Context, req *communicationmethodfieldapi.DoReadAllCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoReadAllCommunicationMethodFieldResponse, error) {
	result, err := cmf.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	return &communicationmethodfieldapi.DoReadAllCommunicationMethodFieldResponse{CommunicationMethodField: result}, err
}

func (cmf *communicationMethodFieldService) DoSave(ctx context.Context, req *communicationmethodfieldapi.DoSaveCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse, error) {
	res, err := doUpdate(ctx, cmf.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cmf.repo, req)
			}
		}
	}

	return res, err
}

func (cmf *communicationMethodFieldService) DoDelete(ctx context.Context, req *communicationmethodfieldapi.DoDeleteCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoDeleteCommunicationMethodFieldResponse, error) {
	err := cmf.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetFieldCode())

	return &communicationmethodfieldapi.DoDeleteCommunicationMethodFieldResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo communicationmethodfieldrepository.ICommunicationMethodFieldRepository, req *communicationmethodfieldapi.DoSaveCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse, error) {
	err := repo.DoInsert(ctx, req.GetCommunicationMethodField())

	return &communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodfieldrepository.ICommunicationMethodFieldRepository, req *communicationmethodfieldapi.DoSaveCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse, error) {
	err := repo.DoUpdate(ctx, req.GetCommunicationMethodField())

	return &communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse{Result: err == nil}, err
}
