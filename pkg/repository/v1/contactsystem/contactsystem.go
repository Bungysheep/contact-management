package contactsystem

import (
	"context"
	"database/sql"

	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	communicationmethodrepository "github.com/bungysheep/contact-management/pkg/repository/v1/communicationmethod"
	contactrepository "github.com/bungysheep/contact-management/pkg/repository/v1/contact"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IContactSystemRepository - Contact System repository interface
type IContactSystemRepository interface {
	DoRead(context.Context, string) (*contactsystem.ContactSystem, error)
	DoReadAll(context.Context) ([]*contactsystem.ContactSystem, error)
	DoInsert(context.Context, *contactsystem.ContactSystem) error
	DoUpdate(context.Context, *contactsystem.ContactSystem) error
	DoDelete(context.Context, string) error
}

type contactSystemRepository struct {
	db *sql.DB
}

// NewContactSystemRepository - Contact System repository implementation
func NewContactSystemRepository(db *sql.DB) IContactSystemRepository {
	return &contactSystemRepository{db: db}
}

func (cntsys *contactSystemRepository) DoRead(ctx context.Context, contactSystemCode string) (*contactsystem.ContactSystem, error) {
	result := contactsystem.NewContactSystem()

	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system WHERE contact_system_code=$1")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact System", err))
	}

	rows, err := stmt.QueryContext(ctx, contactSystemCode)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRead("Contact System", err))
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact System", err))
		}
		return nil, status.Errorf(codes.NotFound, message.DoesNotExist("Contact System"))
	}

	if err := rows.Scan(
		&result.ContactSystemCode,
		&result.Description,
		&result.Details,
		&result.Status,
		&result.GetAudit().CreatedAt,
		&result.GetAudit().ModifiedAt,
		&result.GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact System", err))
	}

	return result, nil
}

func (cntsys *contactSystemRepository) DoReadAll(ctx context.Context) ([]*contactsystem.ContactSystem, error) {
	result := make([]*contactsystem.ContactSystem, 0)

	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system")
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact System", err))
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return result, status.Errorf(codes.Unknown, message.FailedRead("Contact System", err))
	}
	defer rows.Close()

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return result, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact System", err))
			}
			if len(result) == 0 {
				return result, status.Errorf(codes.NotFound, message.DoesNotExist("Contact System"))
			}
			break
		}

		contactSystem := contactsystem.NewContactSystem()
		if err := rows.Scan(
			&contactSystem.ContactSystemCode,
			&contactSystem.Description,
			&contactSystem.Details,
			&contactSystem.Status,
			&contactSystem.GetAudit().CreatedAt,
			&contactSystem.GetAudit().ModifiedAt,
			&contactSystem.GetAudit().Vers); err != nil {
			return result, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact System", err))
		}

		result = append(result, contactSystem)
	}

	return result, nil
}

func (cntsys *contactSystemRepository) DoInsert(ctx context.Context, data *contactsystem.ContactSystem) error {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "INSERT INTO contact_system (contact_system_code, description, details, status, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, 1)")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareInsert("Contact System", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetAudit().GetCreatedAt(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedInsert("Contact System", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.Unknown, message.NoRowInserted())
	}

	return nil
}

func (cntsys *contactSystemRepository) DoUpdate(ctx context.Context, data *contactsystem.ContactSystem) error {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "UPDATE contact_system SET description=$2, details=$3, status=$4, modified_at=$5, vers=vers+1 WHERE contact_system_code=$1")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Contact System", err))
	}

	result, err := stmt.ExecContext(ctx, data.GetContactSystemCode(), data.GetDescription(), data.GetDetails(), data.GetStatus(), data.GetAudit().GetModifiedAt())
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedUpdate("Contact System", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Contact System"))
	}

	return nil
}

func (cntsys *contactSystemRepository) DoDelete(ctx context.Context, contactSystemCode string) error {
	if err := cntsys.anyReferences(ctx, contactSystemCode); err != nil {
		return err
	}

	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "DELETE FROM contact_system WHERE contact_system_code=$1")
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedPrepareDelete("Contact System", err))
	}

	result, err := stmt.ExecContext(ctx, contactSystemCode)
	if err != nil {
		return status.Errorf(codes.Unknown, message.FailedDelete("Contact System", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return status.Errorf(codes.NotFound, message.DoesNotExist("Contact System"))
	}

	return nil
}

func (cntsys *contactSystemRepository) anyReferences(ctx context.Context, contactSystemCode string) error {
	// Check if any Communication Method references
	cm := communicationmethodrepository.NewCommunicationMethodRepository(cntsys.db)
	anyRef, err := cm.AnyReference(ctx, contactSystemCode)
	if err != nil {
		return err
	} else if anyRef {
		return status.Errorf(codes.Unknown, message.FailedDeleteAsReferenceExist("Communication Method"))
	}

	// Check if any Contact references
	cnt := contactrepository.NewContactRepository(cntsys.db)
	anyRef, err = cnt.AnyReference(ctx, contactSystemCode)
	if err != nil {
		return err
	} else if anyRef {
		return status.Errorf(codes.Unknown, message.FailedDeleteAsReferenceExist("Contact"))
	}

	return nil
}
