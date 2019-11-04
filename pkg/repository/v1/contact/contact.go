package contact

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/bungysheep/contact-management/pkg/models/v1/audit"
	"github.com/bungysheep/contact-management/pkg/models/v1/contact"
	contactcommunicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contactcommunicationmethod"
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

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.ContactID,
		&result.FirstName,
		&result.LastName,
		&result.Status,
		&result.GetAudit().CreatedAt,
		&result.GetAudit().ModifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact", err))
	}

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
			&contact.ContactID,
			&contact.FirstName,
			&contact.LastName,
			&contact.Status,
			&contact.GetAudit().CreatedAt,
			&contact.GetAudit().ModifiedAt,
			&contact.GetAudit().Vers); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact", err))
		}

		result = append(result, contact)
	}

	return result, nil
}

func (cm *contactRepository) DoInsert(ctx context.Context, data *contact.Contact) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO contact (contact_system_code, first_name, last_name, status, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, 1)")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Contact", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetFirstName(), data.GetLastName(), data.GetStatus(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
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
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "UPDATE contact SET first_name=$3, last_name=$4, status=$5, modified_at=$6, vers=vers+1 WHERE contact_system_code=$1 AND contact_id=$2")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Contact", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactID(), data.GetFirstName(), data.GetLastName(), data.GetStatus(), data.GetAudit().GetModifiedAt())
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
