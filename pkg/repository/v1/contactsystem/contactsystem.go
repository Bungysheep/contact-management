package contactsystem

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	contactsystemmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
)

// IContactSystemRepository - Contact System repository interface
type IContactSystemRepository interface {
	DoRead(context.Context, string) (*contactsystemmodel.ContactSystem, messagemodel.IMessage)
	DoReadAll(context.Context) ([]*contactsystemmodel.ContactSystem, messagemodel.IMessage)
	DoInsert(context.Context, *contactsystemmodel.ContactSystem) messagemodel.IMessage
	DoUpdate(context.Context, *contactsystemmodel.ContactSystem) messagemodel.IMessage
	DoDelete(context.Context, string) messagemodel.IMessage
}

type contactSystemRepository struct {
	db *sql.DB
}

// NewContactSystemRepository - Contact System repository implementation
func NewContactSystemRepository(db *sql.DB) IContactSystemRepository {
	return &contactSystemRepository{db: db}
}

func (cntsys *contactSystemRepository) DoRead(ctx context.Context, contactSystemCode string) (*contactsystemmodel.ContactSystem, messagemodel.IMessage) {
	result := contactsystemmodel.NewContactSystem()

	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return nil, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, description, details, status, 
			created_at, modified_at, vers 
		FROM contact_system 
		WHERE contact_system_code=$1`)
	if err != nil {
		return nil, message.FailedPrepareRead("Contact System", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return nil, message.FailedRead("Contact System", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, message.FailedRetrieveRow("Contact System", err)
		}
		return nil, message.DoesNotExist("Contact System")
	}

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.Description,
		&result.Details,
		&result.Status,
		&result.GetAudit().CreatedAt,
		&result.GetAudit().ModifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, message.FailedRetrieveValues("Contact System", err)
	}

	return result, nil
}

func (cntsys *contactSystemRepository) DoReadAll(ctx context.Context) ([]*contactsystemmodel.ContactSystem, messagemodel.IMessage) {
	result := make([]*contactsystemmodel.ContactSystem, 0)

	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return result, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, description, details, status, 
			created_at, modified_at, vers 
		FROM contact_system`)
	if err != nil {
		return result, message.FailedPrepareRead("Contact System", err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return result, message.FailedRead("Contact System", err)
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, message.FailedRetrieveRow("Contact System", err)
			}
			if len(result) == 0 {
				return result, message.DoesNotExist("Contact System")
			}
			break
		}

		contactSystem := contactsystemmodel.NewContactSystem()
		if err := rows.Scan(
			&contactSystem.ContactSystemCode,
			&contactSystem.Description,
			&contactSystem.Details,
			&contactSystem.Status,
			&contactSystem.GetAudit().CreatedAt,
			&contactSystem.GetAudit().ModifiedAt,
			&contactSystem.GetAudit().Vers); err != nil {
			return result, message.FailedRetrieveValues("Contact System", err)
		}

		result = append(result, contactSystem)
	}

	return result, nil
}

func (cntsys *contactSystemRepository) DoInsert(ctx context.Context, data *contactsystemmodel.ContactSystem) messagemodel.IMessage {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`INSERT INTO contact_system 
			(contact_system_code, description, details, status, 
			created_at, modified_at, vers) 
		VALUES ($1, $2, $3, $4, 
			$5, $6, 1)`)
	if err != nil {
		return message.FailedPrepareInsert("Contact System", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedInsert("Contact System", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.NoRowInserted()
	}

	return nil
}

func (cntsys *contactSystemRepository) DoUpdate(ctx context.Context, data *contactsystemmodel.ContactSystem) messagemodel.IMessage {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`UPDATE contact_system SET description=$2, details=$3, status=$4, 
			modified_at=$5, vers=vers+1 
		WHERE contact_system_code=$1`)
	if err != nil {
		return message.FailedPrepareUpdate("Contact System", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedUpdate("Contact System", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Contact System")
	}

	return nil
}

func (cntsys *contactSystemRepository) DoDelete(ctx context.Context, contactSystemCode string) messagemodel.IMessage {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM contact_system 
		WHERE contact_system_code=$1`)
	if err != nil {
		return message.FailedPrepareDelete("Contact System", err)
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode)
	if err != nil {
		return message.FailedDelete("Contact System", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Contact System")
	}

	return nil
}
