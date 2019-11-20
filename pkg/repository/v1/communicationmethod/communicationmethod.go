package communicationmethod

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	communicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethod"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
)

// ICommunicationMethodRepository - Communication Method repository interface
type ICommunicationMethodRepository interface {
	DoRead(context.Context, string, string) (*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage)
	DoReadAll(context.Context, string) ([]*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage)
	DoInsert(context.Context, *communicationmethodmodel.CommunicationMethod) messagemodel.IMessage
	DoUpdate(context.Context, *communicationmethodmodel.CommunicationMethod) messagemodel.IMessage
	DoDelete(context.Context, string, string) messagemodel.IMessage
	AnyReference(context.Context, string) (bool, messagemodel.IMessage)
}

type communicationMethodRepository struct {
	db *sql.DB
}

// NewCommunicationMethodRepository - Communication Method repository implementation
func NewCommunicationMethodRepository(db *sql.DB) ICommunicationMethodRepository {
	return &communicationMethodRepository{db: db}
}

func (cm *communicationMethodRepository) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string) (*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage) {
	result := communicationmethodmodel.NewCommunicationMethod()

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return nil, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
			created_at, modified_at, vers 
		FROM communication_method 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2`)
	if err != nil {
		return nil, message.FailedPrepareRead("Communication Method", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return nil, message.FailedRead("Communication Method", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, message.FailedRetrieveRow("Communication Method", err)
		}
		return nil, message.DoesNotExist("Communication Method")
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
		return nil, message.FailedRetrieveValues("Communication Method", err)
	}

	return result, nil
}

func (cm *communicationMethodRepository) DoReadAll(ctx context.Context, contactSystemCode string) ([]*communicationmethodmodel.CommunicationMethod, messagemodel.IMessage) {
	result := make([]*communicationmethodmodel.CommunicationMethod, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, communication_method_code, description, details, status, format_field, 
			created_at, modified_at, vers 
		FROM communication_method 
		WHERE contact_system_code=$1`)
	if err != nil {
		return result, message.FailedPrepareRead("Communication Method", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return result, message.FailedRead("Communication Method", err)
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, message.FailedRetrieveRow("Communication Method", err)
			}
			if len(result) == 0 {
				return result, message.DoesNotExist("Communication Method")
			}
			break
		}

		communicationMethod := communicationmethodmodel.NewCommunicationMethod()
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
			return result, message.FailedRetrieveValues("Communication Method", err)
		}

		result = append(result, communicationMethod)
	}

	return result, nil
}

func (cm *communicationMethodRepository) DoInsert(ctx context.Context, data *communicationmethodmodel.CommunicationMethod) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`INSERT INTO communication_method 
			(contact_system_code, communication_method_code, description, details, status, format_field, 
			created_at, modified_at, vers) 
		VALUES ($1, $2, $3, $4, $5, $6, 
			$7, $8, 1)`)
	if err != nil {
		return message.FailedPrepareInsert("Communication Method", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetFormatField(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedInsert("Communication Method", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.NoRowInserted()
	}

	return nil
}

func (cm *communicationMethodRepository) DoUpdate(ctx context.Context, data *communicationmethodmodel.CommunicationMethod) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`UPDATE communication_method 
		SET description=$3, details=$4, status=$5, format_field=$6, 
			modified_at=$7, vers=vers+1 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2`)
	if err != nil {
		return message.FailedPrepareUpdate("Communication Method", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetFormatField(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedUpdate("Communication Method", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Communication Method")
	}

	return nil
}

func (cm *communicationMethodRepository) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM communication_method 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2`)
	if err != nil {
		return message.FailedPrepareDelete("Communication Method", err)
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return message.FailedDelete("Communication Method", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Communication Method")
	}

	return nil
}

func (cm *communicationMethodRepository) AnyReference(ctx context.Context, contactSystemCode string) (bool, messagemodel.IMessage) {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return false, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT 1 
		FROM communication_method 
		WHERE contact_system_code=$1`)
	if err != nil {
		return false, message.FailedPrepareRead("Communication Method", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return false, message.FailedRead("Communication Method", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return false, message.FailedRetrieveRow("Communication Method", err)
		}
		return false, nil
	}

	return true, nil
}
