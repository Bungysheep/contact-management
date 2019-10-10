package contact

import (
	"context"

	"github.com/bungysheep/contact-management/pkg/api/v1/contact"
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
	return nil, status.Errorf(codes.Unimplemented, "DoRead has not been implemented.")
}

func (cnt *contactService) DoReadAll(ctx context.Context, req *contact.DoReadAllRequest) (*contact.DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "DoReadAll has not been implemented.")
}

func (cnt *contactService) DoSave(ctx context.Context, req *contact.DoSaveRequest) (*contact.DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "DoSave has not been implemented.")
}

func (cnt *contactService) DoDelete(ctx context.Context, req *contact.DoDeleteRequest) (*contact.DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "DoDelete has not been implemented.")
}
