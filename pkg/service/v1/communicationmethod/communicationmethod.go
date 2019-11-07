package communicationmethod

import (
	"context"
	"database/sql"

	communicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	communicationmethodlabelrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ICommunicationMethodService - Communication Method service interface
type ICommunicationMethodService interface {
	DoRead(context.Context, string, string) (*communicationmethodmodel.CommunicationMethod, error)
	DoReadAll(context.Context, string) ([]*communicationmethodmodel.CommunicationMethod, error)
	DoSave(context.Context, *communicationmethodmodel.CommunicationMethod) error
	DoDelete(context.Context, string, string) error
}

type communicationMethodService struct {
	communicationMethodRepo      communicationmethodrepository.ICommunicationMethodRepository
	communicationMethodFieldRepo communicationmethodfieldrepository.ICommunicationMethodFieldRepository
	communicationMethodLabelRepo communicationmethodlabelrepository.ICommunicationMethodLabelRepository
}

// NewCommunicationMethodService - Communication Method service implementation
func NewCommunicationMethodService(db *sql.DB) ICommunicationMethodService {
	return &communicationMethodService{
		communicationMethodRepo:      communicationmethodrepository.NewCommunicationMethodRepository(db),
		communicationMethodFieldRepo: communicationmethodfieldrepository.NewCommunicationMethodFieldRepository(db),
		communicationMethodLabelRepo: communicationmethodlabelrepository.NewCommunicationMethodLabelRepository(db),
	}
}

func (cm *communicationMethodService) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string) (*communicationmethodmodel.CommunicationMethod, error) {
	return cm.communicationMethodRepo.DoRead(ctx, contactSystemCode, communicationMethodCode)
}

func (cm *communicationMethodService) DoReadAll(ctx context.Context, contactSystemCode string) ([]*communicationmethodmodel.CommunicationMethod, error) {
	return cm.communicationMethodRepo.DoReadAll(ctx, contactSystemCode)
}

func (cm *communicationMethodService) DoSave(ctx context.Context, data *communicationmethodmodel.CommunicationMethod) error {
	if err := cm.communicationMethodRepo.DoUpdate(ctx, data); err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return cm.communicationMethodRepo.DoInsert(ctx, data)
			}
		}
	}

	return nil
}

func (cm *communicationMethodService) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string) error {
	if err := cm.communicationMethodFieldRepo.DoDeleteAll(ctx, contactSystemCode, communicationMethodCode); err != nil {
		return err
	}

	if err := cm.communicationMethodLabelRepo.DoDeleteAll(ctx, contactSystemCode, communicationMethodCode); err != nil {
		return err
	}

	return cm.communicationMethodRepo.DoDelete(ctx, contactSystemCode, communicationMethodCode)
}
