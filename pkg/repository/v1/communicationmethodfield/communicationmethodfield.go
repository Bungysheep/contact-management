package communicationmethodfield

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	communicationmethodfieldmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
)

// ICommunicationMethodFieldRepository - Communication Method Field repository interface
type ICommunicationMethodFieldRepository interface {
	DoRead(context.Context, string, string, string) (*communicationmethodfieldmodel.CommunicationMethodField, messagemodel.IMessage)
	DoReadAll(context.Context, string, string) ([]*communicationmethodfieldmodel.CommunicationMethodField, messagemodel.IMessage)
	DoInsert(context.Context, *communicationmethodfieldmodel.CommunicationMethodField) messagemodel.IMessage
	DoUpdate(context.Context, *communicationmethodfieldmodel.CommunicationMethodField) messagemodel.IMessage
	DoDelete(context.Context, string, string, string) messagemodel.IMessage
	DoDeleteAll(context.Context, string, string) messagemodel.IMessage
}

type communicationMethodFieldRepository struct {
	db *sql.DB
}

// NewCommunicationMethodFieldRepository - Communication Method Field repository implementation
func NewCommunicationMethodFieldRepository(db *sql.DB) ICommunicationMethodFieldRepository {
	return &communicationMethodFieldRepository{db: db}
}

func (cmf *communicationMethodFieldRepository) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) (*communicationmethodfieldmodel.CommunicationMethodField, messagemodel.IMessage) {
	result := communicationmethodfieldmodel.NewCommunicationMethodField()

	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return nil, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
			created_at, modified_at, vers 
		FROM communication_method_field 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2 
			AND field_code=$3`)
	if err != nil {
		return nil, message.FailedPrepareRead("Communication Method Field", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicationMethodCode, fieldCode)
	if err != nil {
		return nil, message.FailedRead("Communication Method Field", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, message.FailedRetrieveRow("Communication Method Field", err)
		}
		return nil, message.DoesNotExist("Communication Method Field")
	}

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.CommunicationMethodCode,
		&result.FieldCode,
		&result.Caption,
		&result.Sequence,
		&result.GetAudit().CreatedAt,
		&result.GetAudit().ModifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, message.FailedRetrieveValues("Communication Method Field", err)
	}

	return result, nil
}

func (cmf *communicationMethodFieldRepository) DoReadAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) ([]*communicationmethodfieldmodel.CommunicationMethodField, messagemodel.IMessage) {
	result := make([]*communicationmethodfieldmodel.CommunicationMethodField, 0)

	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return result, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, communication_method_code, field_code, caption, sequence, 
			created_at, modified_at, vers 
		FROM communication_method_field 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2`)
	if err != nil {
		return result, message.FailedPrepareRead("Communication Method Field", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return result, message.FailedRead("Communication Method Field", err)
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, message.FailedRetrieveRow("Communication Method Field", err)
			}
			if len(result) == 0 {
				return result, message.DoesNotExist("Communication Method Field")
			}
			break
		}

		communicationMethodField := communicationmethodfieldmodel.NewCommunicationMethodField()
		if err := rows.Scan(
			&communicationMethodField.ContactSystemCode,
			&communicationMethodField.CommunicationMethodCode,
			&communicationMethodField.FieldCode,
			&communicationMethodField.Caption,
			&communicationMethodField.Sequence,
			&communicationMethodField.GetAudit().CreatedAt,
			&communicationMethodField.GetAudit().ModifiedAt,
			&communicationMethodField.GetAudit().Vers); err != nil {
			return result, message.FailedRetrieveValues("Communication Method Field", err)
		}

		result = append(result, communicationMethodField)
	}

	return result, nil
}

func (cmf *communicationMethodFieldRepository) DoInsert(ctx context.Context, data *communicationmethodfieldmodel.CommunicationMethodField) messagemodel.IMessage {
	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`INSERT INTO communication_method_field 
			(contact_system_code, communication_method_code, field_code, caption, sequence, 
			created_at, modified_at, vers) 
		VALUES ($1, $2, $3, $4, $5, 
			$6, $7, 1)`)
	if err != nil {
		return message.FailedPrepareInsert("Communication Method Field", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetFieldCode(), data.GetCaption(), data.GetSequence(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedInsert("Communication Method Field", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.NoRowInserted()
	}

	return nil
}

func (cmf *communicationMethodFieldRepository) DoUpdate(ctx context.Context, data *communicationmethodfieldmodel.CommunicationMethodField) messagemodel.IMessage {
	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`UPDATE communication_method_field 
		SET caption=$4, sequence=$5, 
			modified_at=$6, vers=vers+1 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2 
			AND field_code=$3`)
	if err != nil {
		return message.FailedPrepareUpdate("Communication Method Field", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetFieldCode(), data.GetCaption(), data.GetSequence(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedUpdate("Communication Method Field", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Communication Method Field")
	}

	return nil
}

func (cmf *communicationMethodFieldRepository) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string, fieldCode string) messagemodel.IMessage {
	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM communication_method_field 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2 
			AND field_code=$3`)
	if err != nil {
		return message.FailedPrepareDelete("Communication Method Field", err)
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode, fieldCode)
	if err != nil {
		return message.FailedDelete("Communication Method Field", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Communication Method Field")
	}

	return nil
}

func (cmf *communicationMethodFieldRepository) DoDeleteAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) messagemodel.IMessage {
	conn, err := cmf.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM communication_method_field 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2`)
	if err != nil {
		return message.FailedPrepareDelete("All Communication Method Fields", err)
	}

	_, err = stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return message.FailedDelete("All Communication Method Fields", err)
	}

	return nil
}
