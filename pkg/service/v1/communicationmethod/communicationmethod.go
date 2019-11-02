package communicationmethod

import (
	"context"

	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodService struct {
	repo communicationmethodrepository.ICommunicationMethodRepository
}

// NewCommunicationMethodService - Communication Method service implementation
func NewCommunicationMethodService(repo communicationmethodrepository.ICommunicationMethodRepository) communicationmethodapi.CommunicationMethodServiceServer {
	return &communicationMethodService{repo: repo}
}

func (cm *communicationMethodService) DoRead(ctx context.Context, req *communicationmethodapi.DoReadCommunicationMethodRequest) (*communicationmethodapi.DoReadCommunicationMethodResponse, error) {
	result, err := cm.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	return &communicationmethodapi.DoReadCommunicationMethodResponse{CommunicationMethod: result}, err
}

func (cm *communicationMethodService) DoReadAll(ctx context.Context, req *communicationmethodapi.DoReadAllCommunicationMethodRequest) (*communicationmethodapi.DoReadAllCommunicationMethodResponse, error) {
	result, err := cm.repo.DoReadAll(ctx, req.GetContactSystemCode())

	return &communicationmethodapi.DoReadAllCommunicationMethodResponse{CommunicationMethod: result}, err
}

func (cm *communicationMethodService) DoSave(ctx context.Context, req *communicationmethodapi.DoSaveCommunicationMethodRequest) (*communicationmethodapi.DoSaveCommunicationMethodResponse, error) {
	res, err := doUpdate(ctx, cm.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cm.repo, req)
			}
		}
	}

	return res, err
}

func (cm *communicationMethodService) DoDelete(ctx context.Context, req *communicationmethodapi.DoDeleteCommunicationMethodRequest) (*communicationmethodapi.DoDeleteCommunicationMethodResponse, error) {
	err := cm.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	return &communicationmethodapi.DoDeleteCommunicationMethodResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo communicationmethodrepository.ICommunicationMethodRepository, req *communicationmethodapi.DoSaveCommunicationMethodRequest) (*communicationmethodapi.DoSaveCommunicationMethodResponse, error) {
	err := repo.DoInsert(ctx, req.GetCommunicationMethod())

	return &communicationmethodapi.DoSaveCommunicationMethodResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodrepository.ICommunicationMethodRepository, req *communicationmethodapi.DoSaveCommunicationMethodRequest) (*communicationmethodapi.DoSaveCommunicationMethodResponse, error) {
	err := repo.DoUpdate(ctx, req.GetCommunicationMethod())

	return &communicationmethodapi.DoSaveCommunicationMethodResponse{Result: err == nil}, err
}
