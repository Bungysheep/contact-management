package communicationmethodfield

import (
	"context"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type communicationMethodFieldService struct {
}

// NewCommunicationMethodFieldService - Communication Method service implementation
func NewCommunicationMethodFieldService() communicationmethodfield.CommunicationMethodFieldServiceServer {
	return &communicationMethodFieldService{}
}

func (cmf *communicationMethodFieldService) DoRead(ctx context.Context, req *communicationmethodfield.DoReadRequest) (*communicationmethodfield.DoReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoRead"))
}

func (cmf *communicationMethodFieldService) DoReadAll(ctx context.Context, req *communicationmethodfield.DoReadAllRequest) (*communicationmethodfield.DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoReadAll"))
}

func (cmf *communicationMethodFieldService) DoSave(ctx context.Context, req *communicationmethodfield.DoSaveRequest) (*communicationmethodfield.DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoSave"))
}

func (cmf *communicationMethodFieldService) DoDelete(ctx context.Context, req *communicationmethodfield.DoDeleteRequest) (*communicationmethodfield.DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoDelete"))
}
