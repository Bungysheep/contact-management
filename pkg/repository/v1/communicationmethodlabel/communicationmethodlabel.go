package communicationmethodlabel

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	communicationmethodlabelmodel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ICommunicationMethodLabelRepository - Communication Method Label repository interface
type ICommunicationMethodLabelRepository interface {
	DoRead(context.Context, string, string, string) (*communicationmethodlabelmodel.CommunicationMethodLabel, error)
	DoReadAll(context.Context, string, string) ([]*communicationmethodlabelmodel.CommunicationMethodLabel, error)
	DoInsert(context.Context, *communicationmethodlabelmodel.CommunicationMethodLabel) error
	DoUpdate(context.Context, *communicationmethodlabelmodel.CommunicationMethodLabel) error
	DoDelete(context.Context, string, string, string) error
	DoDeleteAll(context.Context, string, string) error
}

type communicationMethodLabelRepository struct {
	db *sql.DB
}

// NewCommunicationMethodLabelRepository - Communication Method Label repository implementation
func NewCommunicationMethodLabelRepository(db *sql.DB) ICommunicationMethodLabelRepository {
	return &communicationMethodLabelRepository{db: db}
}

func (cm *communicationMethodLabelRepository) DoRead(ctx context.Context, contactSystemCode string, communicationMethodCode string, communicationMethodLabelCode string) (*communicationmethodlabelmodel.CommunicationMethodLabel, error) {
	result := &communicationmethodlabelmodel.CommunicationMethodLabel{}

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
		FROM communication_method_label 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2 
			AND communication_method_label_code=$3`)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method Label", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicationMethodCode, communicationMethodLabelCode)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRead("Communication Method Label", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Communication Method Label", err))
		}
		return nil, status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Label"))
	}

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.CommunicationMethodCode,
		&result.CommunicationMethodLabelCode,
		&result.Caption); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method Label", err))
	}

	return result, nil
}

func (cm *communicationMethodLabelRepository) DoReadAll(ctx context.Context, contactSystemCode string, communicatonMethodCode string) ([]*communicationmethodlabelmodel.CommunicationMethodLabel, error) {
	result := make([]*communicationmethodlabelmodel.CommunicationMethodLabel, 0)

	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT contact_system_code, communication_method_code, communication_method_label_code, caption 
		FROM communication_method_label 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2`)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Communication Method Label", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode, communicatonMethodCode)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Communication Method Label", err))
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Communication Method Label", err))
			}
			if len(result) == 0 {
				return result, status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Label"))
			}
			break
		}

		communicationMethodLabel := &communicationmethodlabelmodel.CommunicationMethodLabel{}
		if err := rows.Scan(
			&communicationMethodLabel.ContactSystemCode,
			&communicationMethodLabel.CommunicationMethodCode,
			&communicationMethodLabel.CommunicationMethodLabelCode,
			&communicationMethodLabel.Caption); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Communication Method Label", err))
		}

		result = append(result, communicationMethodLabel)
	}

	return result, nil
}

func (cm *communicationMethodLabelRepository) DoInsert(ctx context.Context, data *communicationmethodlabelmodel.CommunicationMethodLabel) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`INSERT INTO communication_method_label 
			(contact_system_code, communication_method_code, communication_method_label_code, caption) 
		VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Communication Method Label", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetCommunicationMethodLabelCode(), data.GetCaption())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedInsert("Communication Method Label", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.Unknown, message.NoRowInserted())
	}

	return nil
}

func (cm *communicationMethodLabelRepository) DoUpdate(ctx context.Context, data *communicationmethodlabelmodel.CommunicationMethodLabel) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`UPDATE communication_method_label 
		SET caption=$4 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2 
			AND communication_method_label_code=$3`)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Communication Method Label", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetCommunicationMethodCode(), data.GetCommunicationMethodLabelCode(), data.GetCaption())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedUpdate("Communication Method Label", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Label"))
	}

	return nil
}

func (cm *communicationMethodLabelRepository) DoDelete(ctx context.Context, contactSystemCode string, communicationMethodCode string, communicationMethodLabelCode string) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM communication_method_label 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2 
			AND communication_method_label_code=$3`)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("Communication Method Label", err))
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode, communicationMethodLabelCode)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("Communication Method Label", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Communication Method Label"))
	}

	return nil
}

func (cm *communicationMethodLabelRepository) DoDeleteAll(ctx context.Context, contactSystemCode string, communicationMethodCode string) error {
	conn, err := cm.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`DELETE FROM communication_method_label 
		WHERE contact_system_code=$1 
			AND communication_method_code=$2`)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("All Communication Method Labels", err))
	}

	_, err = stmt.ExecContext(ctx, contactSystemCode, communicationMethodCode)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("All Communication Method Labels", err))
	}

	return nil
}
