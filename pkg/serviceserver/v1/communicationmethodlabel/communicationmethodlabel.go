package communicationmethodlabel

import (
	"context"

	communicationmethodlabelapi "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodlabel"
	communicationmethodlabelmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
	communicationmethodlabelservice "github.com/bungysheep/contact-management/pkg/service/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodLabelServiceServer struct {
	svc communicationmethodlabelservice.ICommunicationMethodLabelService
}

// NewCommunicationMethodLabelServiceServer - Communication Method Label service server implementation
func NewCommunicationMethodLabelServiceServer(svc communicationmethodlabelservice.ICommunicationMethodLabelService) communicationmethodlabelapi.CommunicationMethodLabelServiceServer {
	return &communicationMethodLabelServiceServer{svc: svc}
}

func (cm *communicationMethodLabelServiceServer) DoRead(ctx context.Context, req *communicationmethodlabelapi.DoReadCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoReadCommunicationMethodLabelResponse, error) {
	result, err := cm.svc.DoRead(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetCommunicationMethodLabelCode())

	return &communicationmethodlabelapi.DoReadCommunicationMethodLabelResponse{CommunicationMethodLabel: communicationMethodLabelModelToAPI(result)}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cm *communicationMethodLabelServiceServer) DoReadAll(ctx context.Context, req *communicationmethodlabelapi.DoReadAllCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoReadAllCommunicationMethodLabelResponse, error) {
	result, err := cm.svc.DoReadAll(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode())

	resp := make([]*communicationmethodlabelapi.CommunicationMethodLabel, 0)

	for _, item := range result {
		resp = append(resp, communicationMethodLabelModelToAPI(item))
	}

	return &communicationmethodlabelapi.DoReadAllCommunicationMethodLabelResponse{CommunicationMethodLabel: resp}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cm *communicationMethodLabelServiceServer) DoSave(ctx context.Context, req *communicationmethodlabelapi.DoSaveCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoSaveCommunicationMethodLabelResponse, error) {
	err := cm.svc.DoSave(ctx, communicationMethodLabelAPIToModel(req.GetCommunicationMethodLabel()))

	return &communicationmethodlabelapi.DoSaveCommunicationMethodLabelResponse{Result: err == nil}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func (cm *communicationMethodLabelServiceServer) DoDelete(ctx context.Context, req *communicationmethodlabelapi.DoDeleteCommunicationMethodLabelRequest) (*communicationmethodlabelapi.DoDeleteCommunicationMethodLabelResponse, error) {
	err := cm.svc.DoDelete(ctx, req.GetContactSystemCode(), req.GetCommunicationMethodCode(), req.GetCommunicationMethodLabelCode())

	return &communicationmethodlabelapi.DoDeleteCommunicationMethodLabelResponse{Result: err == nil}, status.Error(codes.OK, messagemodel.GetMessage(err))
}

func communicationMethodLabelModelToAPI(dataModel *communicationmethodlabelmodel.CommunicationMethodLabel) *communicationmethodlabelapi.CommunicationMethodLabel {
	if dataModel == nil {
		return nil
	}

	communicationMethodLabel := &communicationmethodlabelapi.CommunicationMethodLabel{}
	communicationMethodLabel.ContactSystemCode = dataModel.GetContactSystemCode()
	communicationMethodLabel.CommunicationMethodCode = dataModel.GetCommunicationMethodCode()
	communicationMethodLabel.CommunicationMethodLabelCode = dataModel.GetCommunicationMethodLabelCode()
	communicationMethodLabel.Caption = dataModel.GetCaption()
	return communicationMethodLabel
}

func communicationMethodLabelAPIToModel(data *communicationmethodlabelapi.CommunicationMethodLabel) *communicationmethodlabelmodel.CommunicationMethodLabel {
	if data == nil {
		return nil
	}

	communicationMethodLabel := communicationmethodlabelmodel.NewCommunicationMethodLabel()
	communicationMethodLabel.ContactSystemCode = data.GetContactSystemCode()
	communicationMethodLabel.CommunicationMethodCode = data.GetCommunicationMethodCode()
	communicationMethodLabel.CommunicationMethodLabelCode = data.GetCommunicationMethodLabelCode()
	communicationMethodLabel.Caption = data.GetCaption()
	return communicationMethodLabel
}
