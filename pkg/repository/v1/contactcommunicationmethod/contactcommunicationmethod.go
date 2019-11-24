package contactcommunicationmethod

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	contactcommunicationmethodmodel "github.com/bungysheep/contact-management/pkg/models/v1/contactcommunicationmethod"
	messagemodel "github.com/bungysheep/contact-management/pkg/models/v1/message"
)

// IContactCommunicationMethodRepository - Contact Communication Method repository interface
type IContactCommunicationMethodRepository interface {
	DoRead(context.Context, string, int64, int64) (*contactcommunicationmethodmodel.ContactCommunicationMethod, messagemodel.IMessage)
	DoReadAll(context.Context, string, int64) ([]*contactcommunicationmethodmodel.ContactCommunicationMethod, messagemodel.IMessage)
	DoInsert(context.Context, *contactcommunicationmethodmodel.ContactCommunicationMethod) messagemodel.IMessage
	DoUpdate(context.Context, *contactcommunicationmethodmodel.ContactCommunicationMethod) messagemodel.IMessage
	DoDelete(context.Context, string, int64, int64) messagemodel.IMessage
	DoDeleteAll(context.Context, string, int64) messagemodel.IMessage
}

type contactCommunicationMethodRepository struct {
	db *sql.DB
}

// NewContactCommunicationMethodRepository - Contact Communication Method repository implementation
func NewContactCommunicationMethodRepository(db *sql.DB) IContactCommunicationMethodRepository {
	return &contactCommunicationMethodRepository{db: db}
}

func (cm *contactCommunicationMethodRepository) DoRead(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) (*contactcommunicationmethodmodel.ContactCommunicationMethod, messagemodel.IMessage) {
	result := contactcommunicationmethodmodel.NewContactCommunicationMethod()

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return nil, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
			ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
			ccm.created_at, ccm.modified_at, ccm.vers 
		FROM contact_communication_method ccm
		INNER JOIN communication_method_label cml ON ccm.contact_system_code=cml.contact_system_code
			AND ccm.communication_method_code=cml.communication_method_code
			AND ccm.communication_method_label_code=cml.communication_method_label_code
		WHERE ccm.contact_system_code=$1 
			AND ccm.contact_id=$2 
			AND ccm.contact_communication_method_id=$3`)
	if err != nil {
		return nil, message.FailedPrepareRead("Contact Communication Method", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return nil, message.FailedRead("Contact Communication Method", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, message.FailedRetrieveRow("Contact Communication Method", err)
		}
		return nil, message.DoesNotExist("Contact Communication Method")
	}

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.ContactID,
		&result.ContactCommunicationMethodID,
		&result.CommunicationMethodCode,
		&result.CommunicationMethodLabelCode,
		&result.CommunicationMethodLabelCaption,
		&result.FormatValue,
		&result.Status,
		&result.IsDefault,
		&result.GetAudit().CreatedAt,
		&result.GetAudit().ModifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, message.FailedRetrieveValues("Contact Communication Method", err)
	}

	return result, nil
}

func (cm *contactCommunicationMethodRepository) DoReadAll(ctx context.Context, contactSystemCode string, contactID int64) ([]*contactcommunicationmethodmodel.ContactCommunicationMethod, messagemodel.IMessage) {
	result := make([]*contactcommunicationmethodmodel.ContactCommunicationMethod, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT ccm.contact_system_code, ccm.contact_id, ccm.contact_communication_method_id, 
			ccm.communication_method_code, ccm.communication_method_label_code, cml.caption, ccm.format_value, ccm.status, ccm.is_default, 
			ccm.created_at, ccm.modified_at, ccm.vers 
		FROM contact_communication_method ccm
		INNER JOIN communication_method_label cml ON ccm.contact_system_code=cml.contact_system_code
			AND ccm.communication_method_code=cml.communication_method_code
			AND ccm.communication_method_label_code=cml.communication_method_label_code
		WHERE ccm.contact_system_code=$1 
			AND ccm.contact_id=$2`)
	if err != nil {
		return result, message.FailedPrepareRead("Contact Communication Method", err)
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, contactID)
	if err != nil {
		return result, message.FailedRead("Contact Communication Method", err)
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, message.FailedRetrieveRow("Contact Communication Method", err)
			}
			if len(result) == 0 {
				return result, message.DoesNotExist("Contact Communication Method")
			}
			break
		}

		contactCommunicationMethod := contactcommunicationmethodmodel.NewContactCommunicationMethod()
		if err := rows.Scan(
			&contactCommunicationMethod.ContactSystemCode,
			&contactCommunicationMethod.ContactID,
			&contactCommunicationMethod.ContactCommunicationMethodID,
			&contactCommunicationMethod.CommunicationMethodCode,
			&contactCommunicationMethod.CommunicationMethodLabelCode,
			&contactCommunicationMethod.CommunicationMethodLabelCaption,
			&contactCommunicationMethod.FormatValue,
			&contactCommunicationMethod.Status,
			&contactCommunicationMethod.IsDefault,
			&contactCommunicationMethod.GetAudit().CreatedAt,
			&contactCommunicationMethod.GetAudit().ModifiedAt,
			&contactCommunicationMethod.GetAudit().Vers); err != nil {
			return result, message.FailedRetrieveValues("Contact Communication Method", err)
		}

		result = append(result, contactCommunicationMethod)
	}

	return result, nil
}

func (cm *contactCommunicationMethodRepository) DoInsert(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`INSERT INTO contact_communication_method 
			(contact_system_code, contact_id, communication_method_code, communication_method_label_code, format_value, status, is_default, 
			created_at, modified_at, vers) 
		VALUES ($1, $2, $3, $4, $5, $6, 
			$7, $8, $9, 1)`)
	if err != nil {
		return message.FailedPrepareInsert("Contact Communication Method", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactID(), data.GetCommunicationMethodCode(), data.GetCommunicationMethodLabelCode(), data.GetFormatValue(), data.GetStatus(), data.GetIsDefault(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedInsert("Contact Communication Method", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.NoRowInserted()
	}

	return nil
}

func (cm *contactCommunicationMethodRepository) DoUpdate(ctx context.Context, data *contactcommunicationmethodmodel.ContactCommunicationMethod) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`UPDATE contact_communication_method 
		SET communication_method_code=$4, communication_method_label_code=$5, format_value=$6, status=$7, is_default=$8, 
			modified_at=$9, vers=vers+1 
		WHERE contact_system_code=$1 
			AND contact_id=$2 
			AND contact_communication_method_id=$3`)
	if err != nil {
		return message.FailedPrepareUpdate("Contact Communication Method", err)
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetContactID(), data.GetContactCommunicationMethodID(), data.GetCommunicationMethodCode(), data.GetCommunicationMethodLabelCode(), data.GetFormatValue(), data.GetStatus(), data.GetIsDefault(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return message.FailedUpdate("Contact Communication Method", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Contact Communication Method")
	}

	return nil
}

func (cm *contactCommunicationMethodRepository) DoDelete(ctx context.Context, contactSystemCode string, contactID int64, contactCommunicationMethodID int64) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM contact_communication_method 
		WHERE contact_system_code=$1 
			AND contact_id=$2 
			AND contact_communication_method_id=$3`)
	if err != nil {
		return message.FailedPrepareDelete("Contact Communication Method", err)
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, contactID, contactCommunicationMethodID)
	if err != nil {
		return message.FailedDelete("Contact Communication Method", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return message.DoesNotExist("Contact Communication Method")
	}

	return nil
}

func (cm *contactCommunicationMethodRepository) DoDeleteAll(ctx context.Context, contactSystemCode string, contactID int64) messagemodel.IMessage {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return message.FailedConnectToDatabase(err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM contact_communication_method 
		WHERE contact_system_code=$1 
			AND contact_id=$2`)
	if err != nil {
		return message.FailedPrepareDelete("All Contact Communication Methods", err)
	}

	_, err = stmt.ExecContext(ctx, contactSystemCode, contactID)
	if err != nil {
		return message.FailedDelete("All Contact Communication Methods", err)
	}

	return nil
}
