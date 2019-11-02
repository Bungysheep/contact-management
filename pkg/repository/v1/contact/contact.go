package contact

import (
	"context"
	"database/sql"
	"time"

	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/contact"
	"github.com/bungysheep/contact-management/pkg/common/message"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IContactRepository - Contact repository interface
type IContactRepository interface {
	DoRead(context.Context, string, int64) (*contact.Contact, error)
	DoReadAll(context.Context, string) ([]*contact.Contact, error)
	DoInsert(context.Context, *contact.Contact) error
	DoUpdate(context.Context, *contact.Contact) error
	DoDelete(context.Context, string, int64) error
	AnyReference(context.Context, string) (bool, error)
}

type contactRepository struct {
	db *sql.DB
}

// NewContactRepository - Contact repository implementation
func NewContactRepository(db *sql.DB) IContactRepository {
	return &contactRepository{db: db}
}

func (cm *contactRepository) DoRead(ctx context.Context, contactSystemCode string, contactID int64) (*contact.Contact, error) {
	result := &contact.Contact{Audit: &audit.Audit{}}

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact WHERE contact_system_code=$1 AND contact_id=$2")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, contactID)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRead("Contact", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact", err))
		}
		return nil, status.Errorf(codes.NotFound, message.DoesNotExist("Contact"))
	}

	var createdAt, modifiedAt time.Time

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.ContactId,
		&result.FirstName,
		&result.LastName,
		&result.Status,
		&createdAt,
		&modifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact", err))
	}

	result.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
	result.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

	return result, nil
}

func (cm *contactRepository) DoReadAll(ctx context.Context, contactSystemCode string) ([]*contact.Contact, error) {
	result := make([]*contact.Contact, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact WHERE contact_system_code=$1")
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Contact", err))
	}
	defer rows.Close()

	var createdAt, modifiedAt time.Time

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact", err))
			}
			if len(result) == 0 {
				return result, status.Errorf(codes.NotFound, message.DoesNotExist("Contact"))
			}
			break
		}

		contact := &contact.Contact{Audit: &audit.Audit{}}
		if err := rows.Scan(
			&contact.ContactSystemCode,
			&contact.ContactId,
			&contact.FirstName,
			&contact.LastName,
			&contact.Status,
			&createdAt,
			&modifiedAt,
			&contact.GetAudit().Vers); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact", err))
		}

		contact.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
		contact.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

		result = append(result, contact)
	}

	return result, nil
}

func (cm *contactRepository) DoInsert(ctx context.Context, data *contact.Contact) error {
	createdAt, _ := ptypes.Timestamp(data.GetAudit().GetCreatedAt())
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO contact (contact_system_code, first_name, last_name, status, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, 1)")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Contact", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetFirstName(), data.GetLastName(), data.GetStatus(), createdAt, modifiedAt)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedInsert("Contact", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.Unknown, message.NoRowInserted())
	}

	return nil
}

func (cm *contactRepository) DoUpdate(ctx context.Context, data *contact.Contact) error {
	modifiedAt, _ := ptypes.Timestamp(data.GetAudit().GetModifiedAt())

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "UPDATE contact SET first_name=$3, last_name=$4, status=$5, modified_at=$6, vers=vers+1 WHERE contact_system_code=$1 AND contact_id=$2")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Contact", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactId(), data.GetFirstName(), data.GetLastName(), data.GetStatus(), modifiedAt)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedUpdate("Contact", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Contact"))
	}

	return nil
}

func (cm *contactRepository) DoDelete(ctx context.Context, contactSystemCode string, contactID int64) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "DELETE FROM contact WHERE contact_system_code=$1 AND contact_id=$2")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("Contact", err))
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, contactID)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("Contact", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Contact"))
	}

	// Delete all related Contact Communication Methods
	cmf := contactcommunicationmethodrepository.NewContactCommunicationMethodRepository(cm.db)
	if err := cmf.DoDeleteAll(ctx, contactSystemCode, contactID); err != nil {
		return err
	}

	return nil
}

func (cm *contactRepository) AnyReference(ctx context.Context, contactSystemCode string) (bool, error) {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return false, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT 1 FROM contact WHERE contact_system_code=$1")
	if err != nil {
		return false, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return false, status.Errorf(codes.Unknown, message.FailedRead("Contact", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return false, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact", err))
		}
		return false, nil
	}

	return true, nil
}
