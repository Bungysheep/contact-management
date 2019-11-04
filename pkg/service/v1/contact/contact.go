package contact

import (
	"context"

	auditapi "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactapi "github.com/bungysheep/contact-management/pkg/api/v1/contact"
	contactmodel "github.com/bungysheep/contact-management/pkg/models/v1/contact"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactService struct {
	repo contactrepository.IContactRepository
}

// NewContactService - Contact service implementation
func NewContactService(repo contactrepository.IContactRepository) contactapi.ContactServiceServer {
	return &contactService{repo: repo}
}

func (cnt *contactService) DoRead(ctx context.Context, req *contactapi.DoReadContactRequest) (*contactapi.DoReadContactResponse, error) {
	result, err := cnt.repo.DoRead(ctx, req.GetContactSystemCode(), req.GetContactId())

	resp := &contactapi.Contact{Audit: &auditapi.Audit{}}
	resp.ContactSystemCode = result.GetContactSystemCode()
	resp.ContactId = result.GetContactID()
	resp.FirstName = result.GetFirstName()
	resp.LastName = result.GetLastName()
	resp.Status = result.GetStatus()
	resp.GetAudit().CreatedAt, _ = ptypes.TimestampProto(result.GetAudit().GetCreatedAt())
	resp.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(result.GetAudit().GetModifiedAt())
	resp.GetAudit().Vers = result.GetAudit().GetVers()

	return &contactapi.DoReadContactResponse{Contact: resp}, err
}

func (cnt *contactService) DoReadAll(ctx context.Context, req *contactapi.DoReadAllContactRequest) (*contactapi.DoReadAllContactResponse, error) {
	result, err := cnt.repo.DoReadAll(ctx, req.GetContactSystemCode())

	resp := make([]*contactapi.Contact, 0)

	for _, item := range result {
		contact := &contactapi.Contact{Audit: &auditapi.Audit{}}
		contact.ContactSystemCode = item.GetContactSystemCode()
		contact.ContactId = item.GetContactID()
		contact.FirstName = item.GetFirstName()
		contact.LastName = item.GetLastName()
		contact.Status = item.GetStatus()
		contact.GetAudit().CreatedAt, _ = ptypes.TimestampProto(item.GetAudit().GetCreatedAt())
		contact.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(item.GetAudit().GetModifiedAt())
		contact.GetAudit().Vers = item.GetAudit().GetVers()

		resp = append(resp, contact)
	}

	return &contactapi.DoReadAllContactResponse{Contact: resp}, err
}

func (cnt *contactService) DoSave(ctx context.Context, req *contactapi.DoSaveContactRequest) (*contactapi.DoSaveContactResponse, error) {
	res, err := doUpdate(ctx, cnt.repo, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(ctx, cnt.repo, req)
			}
		}
	}

	return res, err
}

func (cnt *contactService) DoDelete(ctx context.Context, req *contactapi.DoDeleteContactRequest) (*contactapi.DoDeleteContactResponse, error) {
	err := cnt.repo.DoDelete(ctx, req.GetContactSystemCode(), req.GetContactId())

	return &contactapi.DoDeleteContactResponse{Result: err == nil}, err
}

func doInsert(ctx context.Context, repo contactrepository.IContactRepository, req *contactapi.DoSaveContactRequest) (*contactapi.DoSaveContactResponse, error) {
	contact := contactmodel.NewContact()
	contact.ContactSystemCode = req.GetContact().GetContactSystemCode()
	contact.ContactID = req.GetContact().GetContactId()
	contact.FirstName = req.GetContact().GetFirstName()
	contact.LastName = req.GetContact().GetLastName()
	contact.Status = req.GetContact().GetStatus()
	contact.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetContact().GetAudit().GetCreatedAt())
	contact.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetContact().GetAudit().GetModifiedAt())
	contact.GetAudit().Vers = req.GetContact().GetAudit().GetVers()

	err := repo.DoInsert(ctx, contact)

	return &contactapi.DoSaveContactResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactrepository.IContactRepository, req *contactapi.DoSaveContactRequest) (*contactapi.DoSaveContactResponse, error) {
	contact := contactmodel.NewContact()
	contact.ContactSystemCode = req.GetContact().GetContactSystemCode()
	contact.ContactID = req.GetContact().GetContactId()
	contact.FirstName = req.GetContact().GetFirstName()
	contact.LastName = req.GetContact().GetLastName()
	contact.Status = req.GetContact().GetStatus()
	contact.GetAudit().CreatedAt, _ = ptypes.Timestamp(req.GetContact().GetAudit().GetCreatedAt())
	contact.GetAudit().ModifiedAt, _ = ptypes.Timestamp(req.GetContact().GetAudit().GetModifiedAt())
	contact.GetAudit().Vers = req.GetContact().GetAudit().GetVers()

	err := repo.DoUpdate(ctx, contact)

	return &contactapi.DoSaveContactResponse{Result: err == nil}, err
}
