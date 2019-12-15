package communicationmethodfield

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	communicationmethodfieldapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	communicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
	communicationmethodfieldservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodfield"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodFieldServiceServer struct {
	svc communicationmethodfieldservice.ICommunicationMethodFieldService
}

// NewCommunicationMethodFieldServiceServer - Communication Method Field service server implementation
func NewCommunicationMethodFieldServiceServer(svc communicationmethodfieldservice.ICommunicationMethodFieldService) communicationmethodfieldapi.CommunicationMethodFieldServiceServer {
	return &communicationMethodFieldServiceServer{svc: svc}
}

func (cmf *communicationMethodFieldServiceServer) DoRead(ctx context.Context, req *communicationmethodfieldapi.DoReadCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoReadCommunicationMethodFieldResponse, error) {
	result, err := cmf.svc.DoRead(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetFieldCode())

	return &communicationmethodfieldapi.DoReadCommunicationMethodFieldResponse{CommunicationMethodField: communicationMethodFieldModelToAPI(result)}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cmf *communicationMethodFieldServiceServer) DoReadAll(ctx context.Context, req *communicationmethodfieldapi.DoReadAllCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoReadAllCommunicationMethodFieldResponse, error) {
	result, err := cmf.svc.DoReadAll(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	resp := make([]*communicationmethodfieldapi.CommunicationMethodField, 0)

	for _, item := range result {
		resp = append(resp, communicationMethodFieldModelToAPI(item))
	}

	return &communicationmethodfieldapi.DoReadAllCommunicationMethodFieldResponse{CommunicationMethodField: resp}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cmf *communicationMethodFieldServiceServer) DoSave(ctx context.Context, req *communicationmethodfieldapi.DoSaveCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse, error) {
	err := cmf.svc.DoSave(ctx, communicationMethodFieldAPIToModel(req.GetCommunicationMethodField()))

	return &communicationmethodfieldapi.DoSaveCommunicationMethodFieldResponse{Result: err == nil}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cmf *communicationMethodFieldServiceServer) DoDelete(ctx context.Context, req *communicationmethodfieldapi.DoDeleteCommunicationMethodFieldRequest) (*communicationmethodfieldapi.DoDeleteCommunicationMethodFieldResponse, error) {
	err := cmf.svc.DoDelete(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetFieldCode())

	return &communicationmethodfieldapi.DoDeleteCommunicationMethodFieldResponse{Result: err == nil}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func communicationMethodFieldModelToAPI(dataModel *communicationmethodfieldmodel.CommunicationMethodField) *communicationmethodfieldapi.CommunicationMethodField {
	if dataModel == nil {
		return nil
	}

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
	if data == nil {
		return nil
	}

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
