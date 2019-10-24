package contactcommunicationmethod

import (
	"context"
	"database/sql"
	"time"

	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IContactCommunicationMethodRepository - Contact Communication Method repository interface
type IContactCommunicationMethodRepository interface {
	DoRead(context.Context, string, int64, int64) (*contactcommunicationmethod.ContactCommunicationMethod, error)
	DoReadAll(context.Context, string, int64) ([]*contactcommunicationmethod.ContactCommunicationMethod, error)
	DoInsert(context.Context, *contactcommunicationmethod.ContactCommunicationMethod) error
	DoUpdate(context.Context, *contactcommunicationmethod.ContactCommunicationMethod) error
	DoDelete(context.Context, string, int64, int64) error
}

type communicationMethodRepository struct {
	db *sql.DB
}

// NewContactCommunicationMethodRepository - Contact Communication Method repository implementation
func NewContactCommunicationMethodRepository(db *sql.DB) IContactCommunicationMethodRepository {
	return &communicationMethodRepository{db: db}
}

func (cm *communicationMethodRepository) DoRead(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) (*contactcommunicationmethod.ContactCommunicationMethod, error) {
	result := &contactcommunicationmethod.ContactCommunicationMethod{Audit: &audit.Audit{}}

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method WHERE contact_system_code=$1 and contact_id=$2 and contact_communication_method_id=$3")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact Communication Method", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRead("Contact Communication Method", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact Communication Method", err))
		}
		return nil, status.Errorf(codes.NotFound, message.DoesNotExist("Contact Communication Method"))
	}

	var createdAt, modifiedAt time.Time

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.ContactId,
		&result.ContactCommunicationMethodId,
		&result.CommunicationMethodCode,
		&result.CommunicationMethodLabelCode,
		&result.FormatValue,
		&result.IsDefault,
		&createdAt,
		&modifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact Communication Method", err))
	}

	result.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
	result.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

	return result, nil
}

func (cm *communicationMethodRepository) DoReadAll(ctx context.Context, contactSystemCode string, contactID int64) ([]*contactcommunicationmethod.ContactCommunicationMethod, error) {
	result := make([]*contactcommunicationmethod.ContactCommunicationMethod, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers FROM contact_communication_method WHERE contact_system_code=$1 AND contact_id=$2")
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact Communication Method", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, contactID)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Contact Communication Method", err))
	}
	defer rows.Close()

	var createdAt, modifiedAt time.Time

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact Communication Method", err))
			}
			if len(result) == 0 {
				return result, status.Errorf(codes.NotFound, message.DoesNotExist("Contact Communication Method"))
			}
			break
		}

		contactCommunicationMethod := &contactcommunicationmethod.ContactCommunicationMethod{Audit: &audit.Audit{}}
		if err := rows.Scan(
			&contactCommunicationMethod.ContactSystemCode,
			&contactCommunicationMethod.ContactId,
			&contactCommunicationMethod.ContactCommunicationMethodId,
			&contactCommunicationMethod.CommunicationMethodCode,
			&contactCommunicationMethod.CommunicationMethodLabelCode,
			&contactCommunicationMethod.FormatValue,
			&contactCommunicationMethod.IsDefault,
			&createdAt,
			&modifiedAt,
			&contactCommunicationMethod.GetAudit().Vers); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact Communication Method", err))
		}

		contactCommunicationMethod.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
		contactCommunicationMethod.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

		result = append(result, contactCommunicationMethod)
	}

	return result, nil
}

func (cm *communicationMethodRepository) DoInsert(ctx context.Context, data *contactcommunicationmethod.ContactCommunicationMethod) error {
	createdAt, _ := ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO contact_communication_method (contact_system_code, contact_id, communication_method_code, communication_method_label_code, format_value, is_default, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 1)")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Contact Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactId(), data.GetCommunicationMethodCode(), data.GetCommunicationMethodLabelCode(), data.GetFormatValue(), data.GetIsDefault(), createdAt, modifiedAt)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedInsert("Contact Communication Method", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.Unknown, message.NoRowInserted())
	}

	return nil
}

func (cm *communicationMethodRepository) DoUpdate(ctx context.Context, data *contactcommunicationmethod.ContactCommunicationMethod) error {
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "UPDATE contact_communication_method SET communication_method_code=$4, communication_method_label_code=$5, format_value=$6, is_default=$7, modified_at=$8, vers=vers+1 WHERE contact_system_code=$1 AND contact_id=$2 AND contact_communication_method_id=$3")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Contact Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactId(), data.GetContactCommunicationMethodId(), data.GetCommunicationMethodCode(), data.GetCommunicationMethodLabelCode(), data.GetFormatValue(), data.GetIsDefault(), modifiedAt)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedUpdate("Contact Communication Method", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Contact Communication Method"))
	}

	return nil
}

func (cm *communicationMethodRepository) DoDelete(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "DELETE FROM contact_communication_method WHERE contact_system_code=$1 AND contact_id=$2 AND contact_communication_method_id=$3")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("Contact Communication Method", err))
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("Contact Communication Method", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Contact Communication Method"))
	}

	return nil
}
