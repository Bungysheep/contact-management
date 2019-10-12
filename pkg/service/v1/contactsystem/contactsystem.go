package contactsystem

import (
	"context"
	"database/sql"
	"time"

	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	"github.com/bungysheep/contact-management/pkg/common/message"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contactSystemService struct {
	db *sql.DB
}

// NewContactSystemService - Contact System service implementation
func NewContactSystemService(db *sql.DB) contactsystem.ContactSystemServiceServer {
	return &contactSystemService{db: db}
}

func (cntsys *contactSystemService) DoRead(ctx context.Context, req *contactsystem.DoReadRequest) (*contactsystem.DoReadResponse, error) {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system WHERE contact_system_code=$1")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact System", err))
	}

	rows, err := stmt.QueryContext(ctx, req.GetContactSystemCode())
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

	var createdAt, modifiedAt time.Time
	resp := &contactsystem.DoReadResponse{ContactSystem: &contactsystem.ContactSystem{Audit: &audit.Audit{}}}

	if err := rows.Scan(
		&resp.GetContactSystem().ContactSystemCode,
		&resp.GetContactSystem().Description,
		&resp.GetContactSystem().Details,
		&resp.GetContactSystem().Status,
		&createdAt,
		&modifiedAt,
		&resp.GetContactSystem().GetAudit().Vers); err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact System", err))
	}

	resp.GetContactSystem().GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
	resp.GetContactSystem().GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

	return resp, nil
}

func (cntsys *contactSystemService) DoReadAll(ctx context.Context, req *contactsystem.DoReadAllRequest) (*contactsystem.DoReadAllResponse, error) {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedPrepareRead("Contact System", err))
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedRead("Contact System", err))
	}
	defer rows.Close()

	var createdAt, modifiedAt time.Time
	contactSystems := []*contactsystem.ContactSystem{}

	for {
		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return nil, status.Errorf(codes.Unknown, message.FailedRetrieveRow("Contact System", err))
			}
			break
		}

		contactSystem := &contactsystem.ContactSystem{Audit: &audit.Audit{}}
		if err := rows.Scan(
			&contactSystem.ContactSystemCode,
			&contactSystem.Description,
			&contactSystem.Details,
			&contactSystem.Status,
			&createdAt,
			&modifiedAt,
			&contactSystem.GetAudit().Vers); err != nil {
			return nil, status.Errorf(codes.Unknown, message.FailedRetrieveValues("Contact System", err))
		}

		contactSystem.GetAudit().CreatedAt, _ = ptypes.TimestampProto(createdAt)
		contactSystem.GetAudit().ModifiedAt, _ = ptypes.TimestampProto(modifiedAt)

		contactSystems = append(contactSystems, contactSystem)
	}

	return &contactsystem.DoReadAllResponse{ContactSystems: contactSystems}, nil
}

func (cntsys *contactSystemService) DoSave(ctx context.Context, req *contactsystem.DoSaveRequest) (*contactsystem.DoSaveResponse, error) {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	res, err := doUpdate(conn)(ctx, req)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return doInsert(conn)(ctx, req)
			}
		}
	}

	return res, err
}

func (cntsys *contactSystemService) DoDelete(ctx context.Context, req *contactsystem.DoDeleteRequest) (*contactsystem.DoDeleteResponse, error) {
	conn, err := cntsys.db.Conn(ctx)
	if err != nil {
		return &contactsystem.DoDeleteResponse{Result: false}, status.Errorf(codes.Unknown, message.FailedConnectToDatabase(err))
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, "DELETE FROM contact_system WHERE contact_system_code=$1")
	if err != nil {
		return &contactsystem.DoDeleteResponse{Result: false}, status.Errorf(codes.Unknown, message.FailedPrepareDelete("Contact System", err))
	}

	result, err := stmt.ExecContext(ctx, req.GetContactSystemCode())
	if err != nil {
		return &contactsystem.DoDeleteResponse{Result: false}, status.Errorf(codes.Unknown, message.FailedDelete("Contact System", err))
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return &contactsystem.DoDeleteResponse{Result: false}, status.Errorf(codes.NotFound, message.DoesNotExist("Contact System"))
	}

	return &contactsystem.DoDeleteResponse{Result: true}, nil
}

func doInsert(conn *sql.Conn) func(ctx context.Context, req *contactsystem.DoSaveRequest) (*contactsystem.DoSaveResponse, error) {
	return func(ctx context.Context, req *contactsystem.DoSaveRequest) (*contactsystem.DoSaveResponse, error) {
		createdAt, _ := ptypes.Timestamp(req.GetContactSystem().GetAudit().GetCreatedAt())
		modifiedAt, _ := ptypes.Timestamp(req.GetContactSystem().GetAudit().GetModifiedAt())

		stmt, err := conn.PrepareContext(ctx, "INSERT INTO contact_system (contact_system_code, description, details, status, created_at, modified_at, vers) VALUES ($1, $2, $3, $4, $5, $6, 1)")
		if err != nil {
			return &contactsystem.DoSaveResponse{Result: false}, status.Errorf(codes.Unknown, message.FailedPrepareInsert("Contact System", err))
		}

		result, err := stmt.ExecContext(ctx, req.ContactSystem.GetContactSystemCode(), req.ContactSystem.GetDescription(), req.ContactSystem.GetDetails(), req.ContactSystem.GetStatus(), createdAt, modifiedAt)
		if err != nil {
			return &contactsystem.DoSaveResponse{Result: false}, status.Errorf(codes.Unknown, message.FailedInsert("Contact System", err))
		}

		rows, _ := result.RowsAffected()
		if rows == 0 {
			return &contactsystem.DoSaveResponse{Result: false}, status.Errorf(codes.Unknown, message.NoRowInserted())
		}

		return &contactsystem.DoSaveResponse{Result: true}, nil
	}
}

func doUpdate(conn *sql.Conn) func(ctx context.Context, req *contactsystem.DoSaveRequest) (*contactsystem.DoSaveResponse, error) {
	return func(ctx context.Context, req *contactsystem.DoSaveRequest) (*contactsystem.DoSaveResponse, error) {
		modifiedAt, _ := ptypes.Timestamp(req.GetContactSystem().GetAudit().GetModifiedAt())

		stmt, err := conn.PrepareContext(ctx, "UPDATE contact_system SET description=$2, details=$3, status=$4, modified_at=$5, vers=vers+1 WHERE contact_system_code=$1")
		if err != nil {
			return &contactsystem.DoSaveResponse{Result: false}, status.Errorf(codes.Unknown, message.FailedPrepareUpdate("Contact System", err))
		}

		result, err := stmt.ExecContext(ctx, req.ContactSystem.GetContactSystemCode(), req.ContactSystem.GetDescription(), req.ContactSystem.GetDetails(), req.ContactSystem.GetStatus(), modifiedAt)
		if err != nil {
			return &contactsystem.DoSaveResponse{Result: false}, status.Errorf(codes.Unknown, message.FailedUpdate("Contact System", err))
		}

		rows, _ := result.RowsAffected()
		if rows == 0 {
			return &contactsystem.DoSaveResponse{Result: false}, status.Errorf(codes.NotFound, message.DoesNotExist("Contact System"))
		}

		return &contactsystem.DoSaveResponse{Result: true}, nil
	}
}
