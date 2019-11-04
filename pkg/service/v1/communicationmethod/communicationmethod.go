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

	resp := &communicationmethodapi.CommunicationMethod{Audit: &auditapi.Audit{}}
	resp.ContactSystemCode = result.GetContactSystemCode()
	resp.CommunicationMethodCode = result.GetCommunicationMethodCode()
	resp.Description = result.GetDescription()
	resp.Details = result.GetDetails()
	resp.Status = result.GetStatus()
	resp.FormatField = result.GetFormatField()
	resp.GetAudit().CreatedAt, _ = ptypes.TimestampProto(result.GetAudit().GetCreatedAt())
	resp.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(result.GetAudit().GetModifiedAt())
	resp.GetAudit().Vers = result.GetAudit().GetVers()

	return &communicationmethodapi.DoReadCommunicationMethodResponse{CommunicationMethod: resp}, err
}

func (cm *communicationMethodService) DoReadAll(ctx context.Context, req *communicationmethodapi.DoReadAllCommunicationMethodRequest) (*communicationmethodapi.DoReadAllCommunicationMethodResponse, error) {
	result, err := cm.repo.DoReadAll(ctx, req.GetContactSystemCode())

	resp := make([]*communicationmethodapi.CommunicationMethod, 0)

	for _, item := range result {
		contactSystem := &communicationmethodapi.CommunicationMethod{Audit: &auditapi.Audit{}}
		contactSystem.ContactSystemCode = item.GetContactSystemCode()
		contactSystem.CommunicationMethodCode = item.GetCommunicationMethodCode()
		contactSystem.Description = item.GetDescription()
		contactSystem.Details = item.GetDetails()
		contactSystem.Status = item.GetStatus()
		contactSystem.FormatField = item.GetFormatField()
		contactSystem.GetAudit().CreatedAt, _ = ptypes.TimestampProto(item.GetAudit().GetCreatedAt())
		contactSystem.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(item.GetAudit().GetModifiedAt())
		contactSystem.GetAudit().Vers = item.GetAudit().GetVers()

		resp = append(resp, contactSystem)
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
	communicationMethod := communicationmethodmodel.NewCommunicationMethod()
	communicationMethod.ContactSystemCode = req.GetCommunicationMethod().GetContactSystemCode()
	communicationMethod.CommunicationMethodCode = req.GetCommunicationMethod().GetCommunicationMethodCode()
	communicationMethod.Description = req.GetCommunicationMethod().GetDescription()
	communicationMethod.Details = req.GetCommunicationMethod().GetDetails()
	communicationMethod.Status = req.GetCommunicationMethod().GetStatus()
	communicationMethod.FormatField = req.GetCommunicationMethod().GetFormatField()
	communicationMethod.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetCommunicationMethod().GetAudit().GetCreatedAt())
	communicationMethod.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetCommunicationMethod().GetAudit().GetModifiedAt())
	communicationMethod.GetAudit().Vers = req.GetCommunicationMethod().GetAudit().GetVers()

	err := repo.DoInsert(ctx, communicationMethod)

	return &communicationmethodapi.DoSaveCommunicationMethodResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodrepository.ICommunicationMethodRepository, req *communicationmethodapi.DoSaveCommunicationMethodRequest) (*communicationmethodapi.DoSaveCommunicationMethodResponse, error) {
	communicationMethod := communicationmethodmodel.NewCommunicationMethod()
	communicationMethod.ContactSystemCode = req.GetCommunicationMethod().GetContactSystemCode()
	communicationMethod.CommunicationMethodCode = req.GetCommunicationMethod().GetCommunicationMethodCode()
	communicationMethod.Description = req.GetCommunicationMethod().GetDescription()
	communicationMethod.Details = req.GetCommunicationMethod().GetDetails()
	communicationMethod.Status = req.GetCommunicationMethod().GetStatus()
	communicationMethod.FormatField = req.GetCommunicationMethod().GetFormatField()
	communicationMethod.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetCommunicationMethod().GetAudit().GetCreatedAt())
	communicationMethod.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetCommunicationMethod().GetAudit().GetModifiedAt())
	communicationMethod.GetAudit().Vers = req.GetCommunicationMethod().GetAudit().GetVers()

	err := repo.DoUpdate(ctx, communicationMethod)

	return &communicationmethodapi.DoSaveCommunicationMethodResponse{Result: err == nil}, err
}
