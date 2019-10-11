package contactsystem

import (
	"context"

	"github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactSystemService struct {
}

// NewContactSystemService - Contact System service implementation
func NewContactSystemService() contactsystem.ContactSystemServiceServer {
	return &contactSystemService{}
}

func (cntsys *contactSystemService) DoRead(ctx context.Context, req *contactsystem.DoReadRequest) (*contactsystem.DoReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoRead"))
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context, req *contactsystem.DoReadAllRequest) (*contactsystem.DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoReadAll"))
}

func (cntsys *contactSystemService) DoSave(ctx context.Context, req *contactsystem.DoSaveRequest) (*contactsystem.DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoSave"))
}

func (cntsys *contactSystemService) DoDelete(ctx context.Context, req *contactsystem.DoDeleteRequest) (*contactsystem.DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoDelete"))
}
