package contactcommunicationmethodfield

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	contactcommunicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethodfield"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IContactCommunicationMethodFieldRepository - Contact Communication Method Field repository interface
type IContactCommunicationMethodFieldRepository interface {
	DoRead(context.Context, string, int64, int64) ([]*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField, error)
	DoInsert(context.Context, *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) error
	DoUpdate(context.Context, *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) error
	DoDelete(context.Context, string, int64, int64) error
}

type contactCommunicationMethodFieldRepository struct {
	db *sql.DB
}

// NewContactCommunicationMethodFieldRepository - Contact Communication Method Field repository implementation
func NewContactCommunicationMethodFieldRepository(db *sql.DB) IContactCommunicationMethodFieldRepository {
	return &contactCommunicationMethodFieldRepository{db: db}
}

func (cm *contactCommunicationMethodFieldRepository) DoRead(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) ([]*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField, error) {
	result := make([]*contactcommunicationmethodfieldmodel.ContactCommunicationMethodField, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, contact_id, contact_communication_method_id, field_code, field_value 
		FROM contact_communication_method_field
		WHERE contact_system_code=$1 
			AND contact_id=$2
			AND contact_communication_method_id=$3`)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact Communication Method Fields", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Contact Communication Method Fields", err))
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact Communication Method Field", err))
			}
			if len(result) == 0 {
				return result, status.Errorf(codes.NotFound, message.DoesNotExist("Contact Communication Method Field"))
			}
			break
		}

		contactCommunicationMethodField := contactcommunicationmethodfieldmodel.NewContactCommunicationMethodField()
		if err := rows.Scan(
			&contactCommunicationMethodField.ContactSystemCode,
			&contactCommunicationMethodField.ContactID,
			&contactCommunicationMethodField.ContactCommunicationMethodID,
			&contactCommunicationMethodField.FieldCode,
			&contactCommunicationMethodField.FieldValue); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact Communication Method Field", err))
		}

		result = append(result, contactCommunicationMethodField)
	}

	return result, nil
}

func (cm *contactCommunicationMethodFieldRepository) DoInsert(ctx context.Context, data *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`INSERT INTO contact_communication_method_field 
			(contact_system_code, contact_id, contact_communication_method_id, field_code, field_value) 
		VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Contact Communication Method Field", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactID(), data.GetContactCommunicationMethodID(), data.GetFieldCode(), data.GetFieldValue())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedInsert("Contact Communication Method Field", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.Unknown, message.NoRowInserted())
	}

	return nil
}

func (cm *contactCommunicationMethodFieldRepository) DoUpdate(ctx context.Context, data *contactcommunicationmethodfieldmodel.ContactCommunicationMethodField) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`UPDATE contact_communication_method_field 
		SET field_value=$5
		WHERE contact_system_code=$1 
			AND contact_id=$2 
			AND contact_communication_method_id=$3
			AND field_code=$4`)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Contact Communication Method Field", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactID(), data.GetContactCommunicationMethodID(), data.GetFieldCode(), data.GetFieldValue())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedUpdate("Contact Communication Method", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Contact Communication Method Field"))
	}

	return nil
}

func (cm *contactCommunicationMethodFieldRepository) DoDelete(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM contact_communication_method_field 
		WHERE contact_system_code=$1 
			AND contact_id=$2
			AND contact_communication_method_id=$3`)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("All Contact Communication Method Fields", err))
	}

	_, err = stmt.ExecContext(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("All Contact Communication Method Fields", err))
	}

	return nil
}
