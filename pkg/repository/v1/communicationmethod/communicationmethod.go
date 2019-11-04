package communicationmethod

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/bungysheep/contact-management/pkg/models/v1/audit"
	"github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	communicationmethodfieldrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodfield"
	communicationmethodlabelrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ICommunicationMethodRepository - Communication Method repository interface
type ICommunicationMethodRepository interface {
	DoRead(context.Context, string, string) (*communicationmethod.CommunicationMethod, error)
	DoReadAll(context.Context, string) ([]*communicationmethod.CommunicationMethod, error)
	DoInsert(context.Context, *communicationmethod.CommunicationMethod) error
	DoUpdate(context.Context, *communicationmethod.CommunicationMethod) error
	DoDelete(context.Context, string, string) error
	AnyReference(context.Context, string) (bool, error)
}

type communicationMethodRepository struct {
	db *sql.DB
}

// NewCommunicationMethodRepository - Communication Method repository implementation
func NewCommunicationMethodRepository(db *sql.DB) ICommunicationMethodRepository {
	return &communicationMethodRepository{db: db}
}

func (cm *communicationMethodRepository) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string) (*communicationmethod.CommunicationMethod, error) {
	result := &communicationmethod.CommunicationMethod{Audit: &audit.Audit{}}

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method WHERE contact_system_code=$1 AND communication_method_code=$2")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRead("Communication Method", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Communication Method", err))
		}
		return nil, status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method"))
	}

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.CommunicationMethodCode,
		&result.Description,
		&result.Details,
		&result.Status,
		&result.FormatField,
		&result.GetAudit().CreatedAt,
		&result.GetAudit().ModifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method", err))
	}

	return result, nil
}

func (cm *communicationMethodRepository) DoReadAll(ctx context.Context, contactSystemCode string) ([]*communicationmethod.CommunicationMethod, error) {
	result := make([]*communicationmethod.CommunicationMethod, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method WHERE contact_system_code=$1")
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Communication Method", err))
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Communication Method", err))
			}
			if len(result) == 0 {
				return result, status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method"))
			}
			break
		}

		communicationMethod := &communicationmethod.CommunicationMethod{Audit: &audit.Audit{}}
		if err := rows.Scan(
			&communicationMethod.ContactSystemCode,
			&communicationMethod.CommunicationMethodCode,
			&communicationMethod.Description,
			&communicationMethod.Details,
			&communicationMethod.Status,
			&communicationMethod.FormatField,
			&communicationMethod.GetAudit().CreatedAt,
			&communicationMethod.GetAudit().ModifiedAt,
			&communicationMethod.GetAudit().Vers); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method", err))
		}

		result = append(result, communicationMethod)
	}

	return result, nil
}

func (cm *communicationMethodRepository) DoInsert(ctx context.Context, data *communicationmethod.CommunicationMethod) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO communication_method (contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 1)")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetFormatField(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedInsert("Communication Method", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.Unknown, message.NoRowInserted())
	}

	return nil
}

func (cm *communicationMethodRepository) DoUpdate(ctx context.Context, data *communicationmethod.CommunicationMethod) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "UPDATE communication_method SET description=$3, details=$4, status=$5, format_field=$6, modified_at=$7, vers=vers+1 WHERE contact_system_code=$1 AND communication_method_code=$2")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetFormatField(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedUpdate("Communication Method", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method"))
	}

	return nil
}

func (cm *communicationMethodRepository) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "DELETE FROM communication_method WHERE contact_system_code=$1 AND communication_method_code=$2")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("Communication Method", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method"))
	}

	// Delete all related Communication Method Fields
	cmf := communicationmethodfieldrepository.NewCommunicationMethodFieldRepository(cm.db)
	if err := cmf.DoDeleteAll(ctx, contactSystemCode, communicationMethodCode); err != nil {
		return err
	}

	// Delete all related Communication Method Labels
	cml := communicationmethodlabelrepository.NewCommunicationMethodLabelRepository(cm.db)
	if err := cml.DoDeleteAll(ctx, contactSystemCode, communicationMethodCode); err != nil {
		return err
	}

	return nil
}

func (cm *communicationMethodRepository) AnyReference(ctx context.Context, contactSystemCode string) (bool, error) {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return false, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT 1 FROM communication_method WHERE contact_system_code=$1")
	if err != nil {
		return false, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return false, status.Errorf(codes.Unknown, message.FailedRead("Communication Method", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return false, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Communication Method", err))
		}
		return false, nil
	}

	return true, nil
}
