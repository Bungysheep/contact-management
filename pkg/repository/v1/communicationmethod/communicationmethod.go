package communicationmethod

import (
	"context"
	"database/sql"
	"time"

	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/golang/protobuf/ptypes"
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

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, communication_method_code, description, details, status, created_at, modified_at, vers FROM communication_method WHERE contact_system_code=$1 AND communication_method_code=$2")
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

	var createdAt, modifiedAt time.Time

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.CommunicationMethodCode,
		&result.Description,
		&result.Details,
		&result.Status,
		&createdAt,
		&modifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method", err))
	}

	result.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
	result.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

	return result, nil
}

func (cm *communicationMethodRepository) DoReadAll(ctx context.Context, contactSystemCode string) ([]*communicationmethod.CommunicationMethod, error) {
	result := make([]*communicationmethod.CommunicationMethod, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, communication_method_code, description, details, status, created_at, modified_at, vers FROM communication_method WHERE contact_system_code=$1")
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Communication Method", err))
	}
	defer rows.Close()

	var createdAt, modifiedAt time.Time

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

		contactSystem := &communicationmethod.CommunicationMethod{Audit: &audit.Audit{}}
		if err := rows.Scan(
			&contactSystem.ContactSystemCode,
			&contactSystem.CommunicationMethodCode,
			&contactSystem.Description,
			&contactSystem.Details,
			&contactSystem.Status,
			&createdAt,
			&modifiedAt,
			&contactSystem.GetAudit().Vers); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method", err))
		}

		contactSystem.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
		contactSystem.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

		result = append(result, contactSystem)
	}

	return result, nil
}

func (cm *communicationMethodRepository) DoInsert(ctx context.Context, data *communicationmethod.CommunicationMethod) error {
	createdAt, _ := ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO communication_method (contact_system_code, communication_method_code, description, details, status, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, $7, 1)")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), createdAt, modifiedAt)
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
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "UPDATE communication_method SET description=$3, details=$4, status=$5, modified_at=$6, vers=vers+1 WHERE contact_system_code=$1 AND communication_method_code=$2")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), modifiedAt)
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

	return nil
}
