package contactcommunicationmethod

import (
	"context"

	"github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactCommunicationMethodService struct {
}

// NewContactCommunicationMethodService - Communication Method service implementation
func NewContactCommunicationMethodService() contactcommunicationmethod.ContactCommunicationMethodServiceServer {
	return &contactCommunicationMethodService{}
}

func (ccm *contactCommunicationMethodService) DoRead(ctx context.Context, req *contactcommunicationmethod.DoReadRequest) (*contactcommunicationmethod.DoReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoRead"))
}

func (ccm *contactCommunicationMethodService) DoReadAll(ctx context.Context, req *contactcommunicationmethod.DoReadAllRequest) (*contactcommunicationmethod.DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoReadAll"))
}

func (ccm *contactCommunicationMethodService) DoSave(ctx context.Context, req *contactcommunicationmethod.DoSaveRequest) (*contactcommunicationmethod.DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoSave"))
}

func (ccm *contactCommunicationMethodService) DoDelete(ctx context.Context, req *contactcommunicationmethod.DoDeleteRequest) (*contactcommunicationmethod.DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoDelete"))
}
