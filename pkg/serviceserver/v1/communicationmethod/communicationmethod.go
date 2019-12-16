package communicationmethod

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	communicationmethodapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	messageapi "github.com/bungysheep/contact-management/pkg/api/v1/message"
	communicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
	communicationmethodservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethod"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodServiceServer struct {
	svc communicationmethodservice.ICommunicationMethodService
}

// NewCommunicationMethodServiceServer - Communication Method service server implementation
func NewCommunicationMethodServiceServer(svc communicationmethodservice.ICommunicationMethodService) communicationmethodapi.CommunicationMethodServiceServer {
	return &communicationMethodServiceServer{svc: svc}
}

func (cm *communicationMethodServiceServer) DoRead(ctx context.Context, req *communicationmethodapi.DoReadCommunicationMethodRequest) (*communicationmethodapi.DoReadCommunicationMethodResponse, error) {
	result, err := cm.svc.DoRead(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	return &communicationmethodapi.DoReadCommunicationMethodResponse{CommunicationMethod: communicationMethodModelToAPI(result), Message: messageModelToAPI(err)}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cm *communicationMethodServiceServer) DoReadAll(ctx context.Context, req *communicationmethodapi.DoReadAllCommunicationMethodRequest) (*communicationmethodapi.DoReadAllCommunicationMethodResponse, error) {
	result, err := cm.svc.DoReadAll(ctx, req.GetContactSystemCode())

	resp := make([]*communicationmethodapi.CommunicationMethod, 0)

	for _, item := range result {
		resp = append(resp, communicationMethodModelToAPI(item))
	}

	return &communicationmethodapi.DoReadAllCommunicationMethodResponse{CommunicationMethod: resp, Message: messageModelToAPI(err)}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cm *communicationMethodServiceServer) DoSave(ctx context.Context, req *communicationmethodapi.DoSaveCommunicationMethodRequest) (*communicationmethodapi.DoSaveCommunicationMethodResponse, error) {
	err := cm.svc.DoSave(ctx, communicationMethodAPIToModel(req.GetCommunicationMethod()))

	return &communicationmethodapi.DoSaveCommunicationMethodResponse{Result: err == nil, Message: messageModelToAPI(err)}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cm *communicationMethodServiceServer) DoDelete(ctx context.Context, req *communicationmethodapi.DoDeleteCommunicationMethodRequest) (*communicationmethodapi.DoDeleteCommunicationMethodResponse, error) {
	err := cm.svc.DoDelete(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	return &communicationmethodapi.DoDeleteCommunicationMethodResponse{Result: err == nil, Message: messageModelToAPI(err)}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func messageModelToAPI(messageModel messagemodel.IMessage) *messageapi.Message {
	if messageModel == nil {
		return nil
	}

	message := &messageapi.Message{}
	message.Code = messageModel.Code()
	message.Type = messageapi.Message_MessageType(messageModel.Type())
	message.IsError = messageModel.IsError()
	message.IsWarning = messageModel.IsWarning()
	message.ShortDescription = messageModel.ShortDescription()
	message.LongDescription = messageModel.LongDescription()
	return message
}

func communicationMethodModelToAPI(dataModel *communicationmethodmodel.CommunicationMethod) *communicationmethodapi.CommunicationMethod {
	if dataModel == nil {
		return nil
	}

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
	if data == nil {
		return nil
	}

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
