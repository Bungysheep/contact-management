package communicationmethodlabel

import (
	"context"
	"database/sql"

	communicationmethodlabelmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
	communicationmethodlabelrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ICommunicationMethodLabelService - Communication Method Label service interface
type ICommunicationMethodLabelService interface {
	DoRead(context.Context, string, string, string) (*communicationmethodlabelmodel.CommunicationMethodLabel, error)
	DoReadAll(context.Context, string, string) ([]*communicationmethodlabelmodel.CommunicationMethodLabel, error)
	DoSave(context.Context, *communicationmethodlabelmodel.CommunicationMethodLabel) error
	DoDelete(context.Context, string, string, string) error
}

type communicationMethodLabelService struct {
	communicationMethodLabelRepo communicationmethodlabelrepository.ICommunicationMethodLabelRepository
}

// NewCommunicationMethodLabelService - Communication Method Label service implementation
func NewCommunicationMethodLabelService(db *sql.DB) ICommunicationMethodLabelService {
	return &communicationMethodLabelService{
		communicationMethodLabelRepo: communicationmethodlabelrepository.NewCommunicationMethodLabelRepository(db),
	}
}

func (cm *communicationMethodLabelService) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string, communicationMethodLabelCode string) (*communicationmethodlabelmodel.CommunicationMethodLabel, error) {
	return cm.communicationMethodLabelRepo.DoRead(ctx, contactSystemCode, communicationMethodCode, communicationMethodLabelCode)
}

func (cm *communicationMethodLabelService) DoReadAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) ([]*communicationmethodlabelmodel.CommunicationMethodLabel, error) {
	return cm.communicationMethodLabelRepo.DoReadAll(ctx, contactSystemCode, communicationMethodCode)
}

func (cm *communicationMethodLabelService) DoSave(ctx context.Context, data *communicationmethodlabelmodel.CommunicationMethodLabel) error {
	if err := cm.communicationMethodLabelRepo.DoUpdate(ctx, data); err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return cm.communicationMethodLabelRepo.DoInsert(ctx, data)
			}
		}
	}

	return nil
}

func (cm *communicationMethodLabelService) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string, communicationMethodLabelCode string) error {
	return cm.communicationMethodLabelRepo.DoDelete(ctx, contactSystemCode, communicationMethodCode, communicationMethodLabelCode)
}
