package communicationmethodfield

import (
	"context"
	"database/sql"
	"time"

	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ICommunicationMethodFieldRepository - Communication Method Field repository interface
type ICommunicationMethodFieldRepository interface {
	DoRead(context.Context, string, string, string) (*communicationmethodfield.CommunicationMethodField, error)
	DoReadAll(context.Context, string, string) ([]*communicationmethodfield.CommunicationMethodField, error)
	DoInsert(context.Context, *communicationmethodfield.CommunicationMethodField) error
	DoUpdate(context.Context, *communicationmethodfield.CommunicationMethodField) error
	DoDelete(context.Context, string, string, string) error
	DoDeleteAll(context.Context, string, string) error
}

type communicationMethodFieldRepository struct {
	db *sql.DB
}

// NewCommunicationMethodFieldRepository - Communication Method Field repository implementation
func NewCommunicationMethodFieldRepository(db *sql.DB) ICommunicationMethodFieldRepository {
	return &communicationMethodFieldRepository{db: db}
}

func (cmf *communicationMethodFieldRepository) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) (*communicationmethodfield.CommunicationMethodField, error) {
	result := &communicationmethodfield.CommunicationMethodField{Audit: &audit.Audit{}}

	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, communication_method_code, field_code, caption, sequence, created_at, modified_at, vers FROM communication_method_field WHERE contact_system_code=$1 AND communication_method_code=$2 AND field_code=$3")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method Field", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicationMethodCode, fieldCode)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRead("Communication Method Field", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Communication Method Field", err))
		}
		return nil, status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Field"))
	}

	var createdAt, modifiedAt time.Time

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.CommunicationMethodCode,
		&result.FieldCode,
		&result.Caption,
		&result.Sequence,
		&createdAt,
		&modifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method Field", err))
	}

	result.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
	result.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

	return result, nil
}

func (cmf *communicationMethodFieldRepository) DoReadAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) ([]*communicationmethodfield.CommunicationMethodField, error) {
	result := make([]*communicationmethodfield.CommunicationMethodField, 0)

	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, communication_method_code, field_code, caption, sequence, created_at, modified_at, vers FROM communication_method_field WHERE contact_system_code=$1 AND communication_method_code=$2")
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method Field", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Communication Method Field", err))
	}
	defer rows.Close()

	var createdAt, modifiedAt time.Time

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Communication Method Field", err))
			}
			if len(result) == 0 {
				return result, status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Field"))
			}
			break
		}

		communicationMethodField := &communicationmethodfield.CommunicationMethodField{Audit: &audit.Audit{}}
		if err := rows.Scan(
			&communicationMethodField.ContactSystemCode,
			&communicationMethodField.CommunicationMethodCode,
			&communicationMethodField.FieldCode,
			&communicationMethodField.Caption,
			&communicationMethodField.Sequence,
			&createdAt,
			&modifiedAt,
			&communicationMethodField.GetAudit().Vers); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method Field", err))
		}

		communicationMethodField.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
		communicationMethodField.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

		result = append(result, communicationMethodField)
	}

	return result, nil
}

func (cmf *communicationMethodFieldRepository) DoInsert(ctx context.Context, data *communicationmethodfield.CommunicationMethodField) error {
	createdAt, _ := ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO communication_method_field (contact_system_code, communication_method_code, field_code, caption, sequence, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, $7, 1)")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Communication Method Field", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetFieldCode(), data.GetCaption(), data.GetSequence(), createdAt, modifiedAt)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedInsert("Communication Method Field", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.Unknown, message.NoRowInserted())
	}

	return nil
}

func (cmf *communicationMethodFieldRepository) DoUpdate(ctx context.Context, data *communicationmethodfield.CommunicationMethodField) error {
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "UPDATE communication_method_field SET caption=$4, sequence=$5, modified_at=$6, vers=vers+1 WHERE contact_system_code=$1 AND communication_method_code=$2 AND field_code=$3")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Communication Method Field", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetFieldCode(), data.GetCaption(), data.GetSequence(), modifiedAt)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedUpdate("Communication Method Field", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Field"))
	}

	return nil
}

func (cmf *communicationMethodFieldRepository) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) error {
	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "DELETE FROM communication_method_field WHERE contact_system_code=$1 AND communication_method_code=$2 AND field_code=$3")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("Communication Method Field", err))
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode, fieldCode)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("Communication Method Field", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Field"))
	}

	return nil
}

func (cmf *communicationMethodFieldRepository) DoDeleteAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) error {
	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "DELETE FROM communication_method_field WHERE contact_system_code=$1 AND communication_method_code=$2")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("All Communication Method Fields", err))
	}

	_, err = stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("All Communication Method Fields", err))
	}

	return nil
}
