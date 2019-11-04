package communicationmethodfield

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	communicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	"github.com/golang/protobuf/ptypes"
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

	resp := &communicationmethodfieldapi.CommunicationMethodField{Audit: &auditapi.Audit{}}
	resp.ContactSystemCode = result.GetContactSystemCode()
	resp.CommunicationMethodCode = result.GetCommunicationMethodCode()
	resp.FieldCode = result.GetFieldCode()
	resp.Caption = result.GetCaption()
	resp.Sequence = result.GetSequence()
	resp.GetAudit().CreatedAt, _ = ptypes.TimestampProto(result.GetAudit().GetCreatedAt())
	resp.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(result.GetAudit().GetModifiedAt())
	resp.GetAudit().Vers = result.GetAudit().GetVers()

	return &communicationmethodfieldapi.DoReadCommunicationMethodFieldResponse{CommunicationMethodField: resp}, err
}

func (cmf *communicationMethodFieldService) DoReadAll(ctx context.Context, req *communicationmethodfieldapi.DoReadAllCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoReadAllCommunicationMethodFieldResponse, error) {
	result, err := cmf.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	resp := make([]*communicationmethodfieldapi.CommunicationMethodField, 0)

	for _, item := range result {
		communicationMethodField := &communicationmethodfieldapi.CommunicationMethodField{Audit: &auditapi.Audit{}}
		communicationMethodField.ContactSystemCode = item.GetContactSystemCode()
		communicationMethodField.CommunicationMethodCode = item.GetCommunicationMethodCode()
		communicationMethodField.FieldCode = item.GetFieldCode()
		communicationMethodField.Caption = item.GetCaption()
		communicationMethodField.Sequence = item.GetSequence()
		communicationMethodField.GetAudit().CreatedAt, _ = ptypes.TimestampProto(item.GetAudit().GetCreatedAt())
		communicationMethodField.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(item.GetAudit().GetModifiedAt())
		communicationMethodField.GetAudit().Vers = item.GetAudit().GetVers()

		resp = append(resp, communicationMethodField)
	}

	return &communicationmethodfieldapi.DoReadAllCommunicationMethodFieldResponse{CommunicationMethodField: resp}, err
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
	communicationMethodField := communicationmethodfieldmodel.NewCommunicationMethodField()
	communicationMethodField.ContactSystemCode = req.GetCommunicationMethodField().GetContactSystemCode()
	communicationMethodField.CommunicationMethodCode = req.GetCommunicationMethodField().GetCommunicationMethodCode()
	communicationMethodField.FieldCode = req.GetCommunicationMethodField().GetFieldCode()
	communicationMethodField.Caption = req.GetCommunicationMethodField().GetCaption()
	communicationMethodField.Sequence = req.GetCommunicationMethodField().GetSequence()
	communicationMethodField.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetCommunicationMethodField().GetAudit().GetCreatedAt())
	communicationMethodField.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetCommunicationMethodField().GetAudit().GetModifiedAt())
	communicationMethodField.GetAudit().Vers = req.GetCommunicationMethodField().GetAudit().GetVers()

	err := repo.DoInsert(ctx, communicationMethodField)

	return &communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodfieldrepository.ICommunicationMethodFieldRepository, req *communicationmethodfieldapi.DoSaveCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse, error) {
	communicationMethodField := communicationmethodfieldmodel.NewCommunicationMethodField()
	communicationMethodField.ContactSystemCode = req.GetCommunicationMethodField().GetContactSystemCode()
	communicationMethodField.CommunicationMethodCode = req.GetCommunicationMethodField().GetCommunicationMethodCode()
	communicationMethodField.FieldCode = req.GetCommunicationMethodField().GetFieldCode()
	communicationMethodField.Caption = req.GetCommunicationMethodField().GetCaption()
	communicationMethodField.Sequence = req.GetCommunicationMethodField().GetSequence()
	communicationMethodField.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetCommunicationMethodField().GetAudit().GetCreatedAt())
	communicationMethodField.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetCommunicationMethodField().GetAudit().GetModifiedAt())
	communicationMethodField.GetAudit().Vers = req.GetCommunicationMethodField().GetAudit().GetVers()

	err := repo.DoUpdate(ctx, communicationMethodField)

	return &communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse{Result: err == nil}, err
}
