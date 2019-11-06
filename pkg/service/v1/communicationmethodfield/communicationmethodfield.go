package communicationmethodfield

import (
	"context"
	"database/sql"

	communicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ICommunicationMethodFieldService - Communication Method Field service interface
type ICommunicationMethodFieldService interface {
	DoRead(context.Context, string, string, string) (*communicationmethodfieldmodel.CommunicationMethodField, error)
	DoReadAll(context.Context, string, string) ([]*communicationmethodfieldmodel.CommunicationMethodField, error)
	DoSave(context.Context, *communicationmethodfieldmodel.CommunicationMethodField) error
	DoDelete(context.Context, string, string, string) error
}

type communicationMethodFieldService struct {
	communicationMethodFieldRepo communicationmethodfieldrepository.ICommunicationMethodFieldRepository
}

// NewCommunicationMethodFieldService - Communication Method Field service implementation
func NewCommunicationMethodFieldService(db *sql.DB) ICommunicationMethodFieldService {
	return &communicationMethodFieldService{
		communicationMethodFieldRepo: communicationmethodfieldrepository.NewCommunicationMethodFieldRepository(db),
	}
}

func (cmf *communicationMethodFieldService) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) (*communicationmethodfieldmodel.CommunicationMethodField, error) {
	return cmf.communicationMethodFieldRepo.DoRead(ctx, contactSystemCode, communicationMethodCode, fieldCode)
}

func (cmf *communicationMethodFieldService) DoReadAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) ([]*communicationmethodfieldmodel.CommunicationMethodField, error) {
	return cmf.communicationMethodFieldRepo.DoReadAll(ctx, contactSystemCode, communicationMethodCode)
}

func (cmf *communicationMethodFieldService) DoSave(ctx context.Context, data *communicationmethodfieldmodel.CommunicationMethodField) error {
	if err := cmf.communicationMethodFieldRepo.DoUpdate(ctx, data); err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return cmf.communicationMethodFieldRepo.DoInsert(ctx, data)
			}
		}
	}

	return nil
}

func (cmf *communicationMethodFieldService) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) error {
	return cmf.communicationMethodFieldRepo.DoDelete(ctx, contactSystemCode, communicationMethodCode, fieldCode)
}
