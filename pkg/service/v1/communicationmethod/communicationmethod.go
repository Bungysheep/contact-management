package communicationmethod

import (
	"context"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodService struct {
}

// NewCommunicationMethodService - Communication Method service implementation
func NewCommunicationMethodService() communicationmethod.CommunicationMethodServiceServer {
	return &communicationMethodService{}
}

func (cm *communicationMethodService) DoRead(ctx context.Context, req *communicationmethod.DoReadRequest) (*communicationmethod.DoReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoRead"))
}

func (cm *communicationMethodService) DoReadAll(ctx context.Context, req *communicationmethod.DoReadAllRequest) (*communicationmethod.DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoReadAll"))
}

func (cm *communicationMethodService) DoSave(ctx context.Context, req *communicationmethod.DoSaveRequest) (*communicationmethod.DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoSave"))
}

func (cm *communicationMethodService) DoDelete(ctx context.Context, req *communicationmethod.DoDeleteRequest) (*communicationmethod.DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoDelete"))
}
