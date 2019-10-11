package contact

import (
	"context"

	"github.com/bungysheep/contact-management/pkg/api/v1/contact"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactService struct {
}

// NewContactService - Contact service implementation
func NewContactService() contact.ContactServiceServer {
	return &contactService{}
}

func (cnt *contactService) DoRead(ctx context.Context, req *contact.DoReadRequest) (*contact.DoReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoRead"))
}

func (cnt *contactService) DoReadAll(ctx context.Context, req *contact.DoReadAllRequest) (*contact.DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoReadAll"))
}

func (cnt *contactService) DoSave(ctx context.Context, req *contact.DoSaveRequest) (*contact.DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoSave"))
}

func (cnt *contactService) DoDelete(ctx context.Context, req *contact.DoDeleteRequest) (*contact.DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, message.Unimplemented("DoDelete"))
}
