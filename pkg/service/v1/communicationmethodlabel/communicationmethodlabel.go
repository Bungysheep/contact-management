package communicationmethodlabel

import (
	"context"

	communicationmethodlabelapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodlabel"
	communicationmethodlabelrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodLabelService struct {
	repo communicationmethodlabelrepository.ICommunicationMethodLabelRepository
}

// NewCommunicationMethodLabelService - Communication Method Label service implementation
func NewCommunicationMethodLabelService(repo communicationmethodlabelrepository.ICommunicationMethodLabelRepository) communicationmethodlabelapi.CommunicationMethodLabelServiceServer {
	return &communicationMethodLabelService{repo: repo}
}

func (cm *communicationMethodLabelService) DoRead(ctx context.Context, req *communicationmethodlabelapi.DoReadRequest) (*communicationmethodlabelapi.DoReadResponse, error) {
	result, err := cm.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetCommunicationMethodLabelCode())

	return &communicationmethodlabelapi.DoReadResponse{CommunicationMethodLabel: result}, err
}

func (cm *communicationMethodLabelService) DoReadAll(ctx context.Context, req *communicationmethodlabelapi.DoReadAllRequest) (*communicationmethodlabelapi.DoReadAllResponse, error) {
	result, err := cm.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	return &communicationmethodlabelapi.DoReadAllResponse{CommunicationMethodLabel: result}, err
}

func (cm *communicationMethodLabelService) DoSave(ctx context.Context, req *communicationmethodlabelapi.DoSaveRequest) (*communicationmethodlabelapi.DoSaveResponse, error) {
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

func (cm *communicationMethodLabelService) DoDelete(ctx context.Context, req *communicationmethodlabelapi.DoDeleteRequest) (*communicationmethodlabelapi.DoDeleteResponse, error) {
	err := cm.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetCommunicationMethodLabelCode())

	return &communicationmethodlabelapi.DoDeleteResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo communicationmethodlabelrepository.ICommunicationMethodLabelRepository, req *communicationmethodlabelapi.DoSaveRequest) (*communicationmethodlabelapi.DoSaveResponse, error) {
	err := repo.DoInsert(ctx, req.GetCommunicationMethodLabel())

	return &communicationmethodlabelapi.DoSaveResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodlabelrepository.ICommunicationMethodLabelRepository, req *communicationmethodlabelapi.DoSaveRequest) (*communicationmethodlabelapi.DoSaveResponse, error) {
	err := repo.DoUpdate(ctx, req.GetCommunicationMethodLabel())

	return &communicationmethodlabelapi.DoSaveResponse{Result: err == nil}, err
}
