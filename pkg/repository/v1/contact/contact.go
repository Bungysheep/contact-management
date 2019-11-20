package contact

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	contactmodel "github.com/bungysheep/contact-management/pkg/models/v1/contact"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
)

// IContactRepository - Contact repository interface
type IContactRepository interface {
	DoRead(context.Context, string, int64) (*contactmodel.Contact, messagemodel.IMessage)
	DoReadAll(context.Context, string) ([]*contactmodel.Contact, messagemodel.IMessage)
	DoInsert(context.Context, *contactmodel.Contact) messagemodel.IMessage
	DoUpdate(context.Context, *contactmodel.Contact) messagemodel.IMessage
	DoDelete(context.Context, string, int64) messagemodel.IMessage
	AnyReference(context.Context, string) (bool, messagemodel.IMessage)
}

type contactRepository struct {
	db *sql.DB
}

// NewContactRepository - Contact repository implementation
func NewContactRepository(db *sql.DB) IContactRepository {
	return &contactRepository{db: db}
}

func (cm *contactRepository) DoRead(ctx context.Context, contactSystemCode string, contactID int64) (*contactmodel.Contact, messagemodel.IMessage) {
	result := contactmodel.NewContact()

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return nil, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, contact_id, first_name, last_name, status, 
			created_at, modified_at, vers 
		FROM contact 
		WHERE contact_system_code=$1 
			AND contact_id=$2`)
	if err != nil {
		return nil, message.FailedPrepareRead("Contact", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, contactID)
	if err != nil {
		return nil, message.FailedRead("Contact", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, message.FailedRetrieveRow("Contact", err)
		}
		return nil, message.DoesNotExist("Contact")
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
		return nil, message.FailedRetrieveValues("Contact", err)
	}

	return result, nil
}

func (cm *contactRepository) DoReadAll(ctx context.Context, contactSystemCode string) ([]*contactmodel.Contact, messagemodel.IMessage) {
	result := make([]*contactmodel.Contact, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, contact_id, first_name, last_name, status, 
			created_at, modified_at, vers 
		FROM contact 
		WHERE contact_system_code=$1`)
	if err != nil {
		return result, message.FailedPrepareRead("Contact", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return result, message.FailedRead("Contact", err)
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, message.FailedRetrieveRow("Contact", err)
			}
			if len(result) == 0 {
				return result, message.DoesNotExist("Contact")
			}
			break
		}

		contact := contactmodel.NewContact()
		if err := rows.Scan(
			&contact.ContactSystemCode,
			&contact.ContactID,
			&contact.FirstName,
			&contact.LastName,
			&contact.Status,
			&contact.GetAudit().CreatedAt,
			&contact.GetAudit().ModifiedAt,
			&contact.GetAudit().Vers); err != nil {
			return result, message.FailedRetrieveValues("Contact", err)
		}

		result = append(result, contact)
	}

	return result, nil
}

func (cm *contactRepository) DoInsert(ctx context.Context, data *contactmodel.Contact) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`INSERT INTO contact 
			(contact_system_code, first_name, last_name, status, 
			created_at, modified_at, vers) 
		VALUES ($1, $2, $3, $4, 
			$5, $6, 1)`)
	if err != nil {
		return message.FailedPrepareInsert("Contact", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetFirstName(), data.GetLastName(), data.GetStatus(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedInsert("Contact", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.NoRowInserted()
	}

	return nil
}

func (cm *contactRepository) DoUpdate(ctx context.Context, data *contactmodel.Contact) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`UPDATE contact SET first_name=$3, last_name=$4, status=$5, 
			modified_at=$6, vers=vers+1 
		WHERE contact_system_code=$1 
			AND contact_id=$2`)
	if err != nil {
		return message.FailedPrepareUpdate("Contact", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactID(), data.GetFirstName(), data.GetLastName(), data.GetStatus(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedUpdate("Contact", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Contact")
	}

	return nil
}

func (cm *contactRepository) DoDelete(ctx context.Context, contactSystemCode string, contactID int64) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM contact 
		WHERE contact_system_code=$1 
			AND contact_id=$2`)
	if err != nil {
		return message.FailedPrepareDelete("Contact", err)
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, contactID)
	if err != nil {
		return message.FailedDelete("Contact", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Contact")
	}

	return nil
}

func (cm *contactRepository) AnyReference(ctx context.Context, contactSystemCode string) (bool, messagemodel.IMessage) {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return false, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT 1 
		FROM contact 
		WHERE contact_system_code=$1`)
	if err != nil {
		return false, message.FailedPrepareRead("Contact", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return false, message.FailedRead("Contact", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return false, message.FailedRetrieveRow("Contact", err)
		}
		return false, nil
	}

	return true, nil
}
