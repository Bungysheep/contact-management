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

	return &communicationmethodfieldapi.DoReadCommunicationMethodFieldResponse{CommunicationMethodField: communicationMethodFieldModelToAPI(result)}, err
}

func (cmf *communicationMethodFieldService) DoReadAll(ctx context.Context, req *communicationmethodfieldapi.DoReadAllCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoReadAllCommunicationMethodFieldResponse, error) {
	result, err := cmf.repo.DoReadAll(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	resp := make([]*communicationmethodfieldapi.CommunicationMethodField, 0)

	for _, item := range result {
		resp = append(resp, communicationMethodFieldModelToAPI(item))
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
	err := repo.DoInsert(ctx, communicationMethodFieldAPIToModel(req.GetCommunicationMethodField()))

	return &communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo communicationmethodfieldrepository.ICommunicationMethodFieldRepository, req *communicationmethodfieldapi.DoSaveCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse, error) {
	err := repo.DoUpdate(ctx, communicationMethodFieldAPIToModel(req.GetCommunicationMethodField()))

	return &communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse{Result: err == nil}, err
}

func communicationMethodFieldModelToAPI(dataModel *communicationmethodfieldmodel.CommunicationMethodField) *communicationmethodfieldapi.CommunicationMethodField {
	communicationMethodField := &communicationmethodfieldapi.CommunicationMethodField{Audit: &auditapi.Audit{}}
	communicationMethodField.ContactSystemCode = dataModel.GetContactSystemCode()
	communicationMethodField.CommunicationMethodCode = dataModel.GetCommunicationMethodCode()
	communicationMethodField.FieldCode = dataModel.GetFieldCode()
	communicationMethodField.Caption = dataModel.GetCaption()
	communicationMethodField.Sequence = dataModel.GetSequence()
	communicationMethodField.GetAudit().CreatedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetCreatedAt())
	communicationMethodField.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetModifiedAt())
	communicationMethodField.GetAudit().Vers = dataModel.GetAudit().GetVers()
	return communicationMethodField
}

func communicationMethodFieldAPIToModel(data *communicationmethodfieldapi.CommunicationMethodField) *communicationmethodfieldmodel.CommunicationMethodField {
	communicationMethodField := communicationmethodfieldmodel.NewCommunicationMethodField()
	communicationMethodField.ContactSystemCode = data.GetContactSystemCode()
	communicationMethodField.CommunicationMethodCode = data.GetCommunicationMethodCode()
	communicationMethodField.FieldCode = data.GetFieldCode()
	communicationMethodField.Caption = data.GetCaption()
	communicationMethodField.Sequence = data.GetSequence()
	communicationMethodField.GetAudit().CreatedAt, _ = ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	communicationMethodField.GetAudit().ModifiedAt, _ = ptypes.Timestamp(data.GetAudit().GetModifiedAt())
	communicationMethodField.GetAudit().Vers = data.GetAudit().GetVers()
	return communicationMethodField
}
