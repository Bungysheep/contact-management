package communicationmethodlabel

import (
	"context"

	communicationmethodlabelapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodlabel"
	communicationmethodlabelmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
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

func (cm *communicationMethodLabelService) DoRead(ctx context.Context, req *communicationmethodlabelapi.DoReadCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoReadCommunicationMethodLabelResponse, error) {
	result, err := cm.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetCommunicationMethodLabelCode())

	return &communicationmethodlabelapi.DoReadCommunicationMethodLabelResponse{CommunicationMethodLabel: communicationMethodLabelModelToAPI(result)}, err
}

func (cm *communicationMethodLabelService) DoReadAll(ctx context.Context, req *communicationmethodlabelapi.DoReadAllCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoReadAllCommunicationMethodLabelResponse, error) {
	result, err := cm.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	resp := make([]*communicationmethodlabelapi.CommunicationMethodLabel, 0)

	for _, item := range result {
		resp = append(resp, communicationMethodLabelModelToAPI(item))
	}

	return &communicationmethodlabelapi.DoReadAllCommunicationMethodLabelResponse{CommunicationMethodLabel: resp}, err
}

func (cm *communicationMethodLabelService) DoSave(ctx context.Context, req *communicationmethodlabelapi.DoSaveCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoSaveCommunicationMethodLabelResponse, error) {
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

func (cm *communicationMethodLabelService) DoDelete(ctx context.Context, req *communicationmethodlabelapi.DoDeleteCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoDeleteCommunicationMethodLabelResponse, error) {
	err := cm.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetCommunicationMethodLabelCode())

	return &communicationmethodlabelapi.DoDeleteCommunicationMethodLabelResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo communicationmethodlabelrepository.ICommunicationMethodLabelRepository, req *communicationmethodlabelapi.DoSaveCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoSaveCommunicationMethodLabelResponse, error) {
	err := repo.DoInsert(ctx, communicationMethodLabelAPIToModel(req.GetCommunicationMethodLabel()))

	return &communicationmethodlabelapi.DoSaveCommunicationMethodLabelResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodlabelrepository.ICommunicationMethodLabelRepository, req *communicationmethodlabelapi.DoSaveCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoSaveCommunicationMethodLabelResponse, error) {

	err := repo.DoUpdate(ctx, communicationMethodLabelAPIToModel(req.GetCommunicationMethodLabel()))

	return &communicationmethodlabelapi.DoSaveCommunicationMethodLabelResponse{Result: err == nil}, err
}

func communicationMethodLabelModelToAPI(dataModel *communicationmethodlabelmodel.CommunicationMethodLabel) *communicationmethodlabelapi.CommunicationMethodLabel {
	communicationMethodLabel := &communicationmethodlabelapi.CommunicationMethodLabel{}
	communicationMethodLabel.ContactSystemCode = dataModel.GetContactSystemCode()
	communicationMethodLabel.CommunicationMethodCode = dataModel.GetCommunicationMethodCode()
	communicationMethodLabel.CommunicationMethodLabelCode = dataModel.GetCommunicationMethodLabelCode()
	communicationMethodLabel.Caption = dataModel.GetCaption()
	return communicationMethodLabel
}

func communicationMethodLabelAPIToModel(data *communicationmethodlabelapi.CommunicationMethodLabel) *communicationmethodlabelmodel.CommunicationMethodLabel {
	communicationMethodLabel := communicationmethodlabelmodel.NewCommunicationMethodLabel()
	communicationMethodLabel.ContactSystemCode = data.GetContactSystemCode()
	communicationMethodLabel.CommunicationMethodCode = data.GetCommunicationMethodCode()
	communicationMethodLabel.CommunicationMethodLabelCode = data.GetCommunicationMethodLabelCode()
	communicationMethodLabel.Caption = data.GetCaption()
	return communicationMethodLabel
}
