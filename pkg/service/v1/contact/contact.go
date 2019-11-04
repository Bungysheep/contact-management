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

	return &contactapi.DoReadContactResponse{Contact: contactModelToAPI(result)}, err
}

func (cnt *contactService) DoReadAll(ctx context.Context, req *contactapi.DoReadAllContactRequest) (*contactapi.DoReadAllContactResponse, error) {
	result, err := cnt.repo.DoReadAll(ctx, req.GetContactSystemCode())

	resp := make([]*contactapi.Contact, 0)

	for _, item := range result {
		resp = append(resp, contactModelToAPI(item))
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
	err := repo.DoInsert(ctx, contactAPIToModel(req.GetContact()))

	return &contactapi.DoSaveContactResponse{Result: err == nil}, err
}

func doUpdate(ctx context.Context, repo contactrepository.IContactRepository, req *contactapi.DoSaveContactRequest) (*contactapi.DoSaveContactResponse, error) {
	err := repo.DoUpdate(ctx, contactAPIToModel(req.GetContact()))

	return &contactapi.DoSaveContactResponse{Result: err == nil}, err
}

func contactModelToAPI(dataModel *contactmodel.Contact) *contactapi.Contact {
	contact := &contactapi.Contact{Audit: &auditapi.Audit{}}
	contact.ContactSystemCode = dataModel.GetContactSystemCode()
	contact.ContactId = dataModel.GetContactID()
	contact.FirstName = dataModel.GetFirstName()
	contact.LastName = dataModel.GetLastName()
	contact.Status = dataModel.GetStatus()
	contact.GetAudit().CreatedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetCreatedAt())
	contact.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(dataModel.GetAudit().GetModifiedAt())
	contact.GetAudit().Vers = dataModel.GetAudit().GetVers()
	return contact
}

func contactAPIToModel(data *contactapi.Contact) *contactmodel.Contact {
	contact := contactmodel.NewContact()
	contact.ContactSystemCode = data.GetContactSystemCode()
	contact.ContactID = data.GetContactId()
	contact.FirstName = data.GetFirstName()
	contact.LastName = data.GetLastName()
	contact.Status = data.GetStatus()
	contact.GetAudit().CreatedAt, _ = ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	contact.GetAudit().ModifiedAt, _ = ptypes.Timestamp(data.GetAudit().GetModifiedAt())
	contact.GetAudit().Vers = data.GetAudit().GetVers()
	return contact
}
