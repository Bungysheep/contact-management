package communicationmethod

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	communicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	"github.com/golang/protobuf/ptypes"
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

	return &communicationmethodapi.DoReadCommunicationMethodResponse{CommunicationMethod: communicationMethodModelToAPI(result)}, err
}

func (cm *communicationMethodService) DoReadAll(ctx context.Context, req *communicationmethodapi.DoReadAllCommunicationMethodRequest) (*communicationmethodapi.DoReadAllCommunicationMethodResponse, error) {
	result, err := cm.repo.DoReadAll(ctx, req.GetContactSystemCode())

	resp := make([]*communicationmethodapi.CommunicationMethod, 0)

	for _, item := range result {
		resp = append(resp, communicationMethodModelToAPI(item))
	}

	return &communicationmethodapi.DoReadAllCommunicationMethodResponse{CommunicationMethod: resp}, err
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
	err := repo.DoInsert(ctx, communicationMethodAPIToModel(req.GetCommunicationMethod()))

	return &communicationmethodapi.DoSaveCommunicationMethodResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodrepository.ICommunicationMethodRepository, req *communicationmethodapi.DoSaveCommunicationMethodRequest) (*communicationmethodapi.DoSaveCommunicationMethodResponse, error) {
	err := repo.DoUpdate(ctx, communicationMethodAPIToModel(req.GetCommunicationMethod()))

	return &communicationmethodapi.DoSaveCommunicationMethodResponse{Result: err == nil}, err
}

func communicationMethodModelToAPI(dataModel *communicationmethodmodel.CommunicationMethod) *communicationmethodapi.CommunicationMethod {
	communicationMethod := &communicationmethodapi.CommunicationMethod{Audit: &auditapi.Audit{}}
	communicationMethod.ContactSystemCode = dataModel.GetContactSystemCode()
	communicationMethod.CommunicationMethodCode = dataModel.GetCommunicationMethodCode()
	communicationMethod.Description = dataModel.GetDescription()
	communicationMethod.Details = dataModel.GetDetails()
	communicationMethod.Status = dataModel.GetStatus()
	communicationMethod.FormatField = dataModel.GetFormatField()
	communicationMethod.GetAudit().CreatedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetCreatedAt())
	communicationMethod.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetModifiedAt())
	communicationMethod.GetAudit().Vers = dataModel.GetAudit().GetVers()
	return communicationMethod
}

func communicationMethodAPIToModel(data *communicationmethodapi.CommunicationMethod) *communicationmethodmodel.CommunicationMethod {
	communicationMethod := communicationmethodmodel.NewCommunicationMethod()
	communicationMethod.ContactSystemCode = data.GetContactSystemCode()
	communicationMethod.CommunicationMethodCode = data.GetCommunicationMethodCode()
	communicationMethod.Description = data.GetDescription()
	communicationMethod.Details = data.GetDetails()
	communicationMethod.Status = data.GetStatus()
	communicationMethod.FormatField = data.GetFormatField()
	communicationMethod.GetAudit().CreatedAt, _ = ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	communicationMethod.GetAudit().ModifiedAt, _ = ptypes.Timestamp(data.GetAudit().GetModifiedAt())
	communicationMethod.GetAudit().Vers = data.GetAudit().GetVers()
	return communicationMethod
}
