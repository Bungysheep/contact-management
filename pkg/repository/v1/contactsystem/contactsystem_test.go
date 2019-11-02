package contactsystem

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/contactsystem"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	repo IContactSystemRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contactsystem.ContactSystem
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewContactSystemRepository(db)

	data = append(data, &contactsystem.ContactSystem{
		ContactSystemCode: "CNTSYS001",
		Description:       "Contact System 1",
		Details:           "Contact System 1",
		Status:            "A",
	}, &contactsystem.ContactSystem{
		ContactSystemCode: "CNTSYS002",
		Description:       "Contact System 2",
		Details:           "Contact System 2",
		Status:            "A",
	}, &contactsystem.ContactSystem{
		ContactSystemCode: "CNTSYS003",
		Description:       "Contact System 3",
		Details:           "Contact System 3",
		Status:            "A",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactSystemRepository(t *testing.T) {
	t.Run("DoRead Contact System", doRead(ctx))

	t.Run("DoReadAll Contact System", doReadAll(ctx))

	t.Run("DoSave Contact System", doSave(ctx))

	t.Run("DoDelete Contact System", doDelete(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailContactSystem(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingContactSystem(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorContactSystem(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingContactSystem(ctx, data[0]))
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail", doReadAllFailContactSystems(ctx, data[0]))

		t.Run("DoReadAll unexisting", doReadAllUnexistingContactSystems(ctx, data[0]))

		t.Run("DoReadAll row error", doReadAllRowErrorContactSystems(ctx, data[0]))

		t.Run("DoReadAll existing", doReadAllExistingContactSystems(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailContactSystem(ctx, data[0]))

		t.Run("DoSave new", doSaveNewContactSystem(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailContactSystem(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingContactSystem(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailContactSystem(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingContactSystem(ctx, data[0]))

		t.Run("DoDelete fail any Communication Method reference", doDeleteFailAnyCommunicationMethodReference(ctx, data[0]))

		t.Run("DoDelete any Communication Method reference", doDeleteAnyCommunicationMethodReference(ctx, data[0]))

		t.Run("DoDelete fail any Contact reference", doDeleteFailAnyContactReference(ctx, data[0]))

		t.Run("DoDelete any Contact reference", doDeleteAnyContactReference(ctx, data[0]))

		t.Run("DoDelete existing", doDeleteExistingContactSystem(ctx, data[0]))
	}
}

func doReadFailContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnError(fmt.Errorf("DoRead contact system failed"))

		res, err := repo.DoRead(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect contact system is nil")
		}
	}
}

func doReadUnexistingContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect contact system is nil")
		}
	}
}

func doReadRowErrorContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoRead row error"))

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect contact system is nil")
		}
	}
}

func doReadExistingContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode())
		if err != nil {
			t.Errorf("Failed to read contact system: %v", err)
		}

		if res == nil {
			t.Errorf("Expect contact system is not nil")
		}

		if res.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res.GetContactSystemCode())
		}

		if res.GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), res.GetDescription())
		}

		if res.GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), res.GetDetails())
		}

		if res.GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), res.GetStatus())
		}
	}
}

func doReadAllFailContactSystems(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WillReturnError(fmt.Errorf("DoReadAll contact system failed"))

		res, err := repo.DoReadAll(ctx)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadAllUnexistingContactSystems(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if res == nil {
			t.Errorf("Expect contact systems is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect no contact systems retrieved")
		}
	}
}

func doReadAllRowErrorContactSystems(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetDescription(), data[0].GetDetails(), data[0].GetStatus(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetDescription(), data[1].GetDetails(), data[1].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetDescription(), data[2].GetDetails(), data[2].GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect response is nil")
		}
	}
}

func doReadAllExistingContactSystems(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "description", "details", "status", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetDescription(), data[0].GetDetails(), data[0].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetDescription(), data[1].GetDetails(), data[1].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetDescription(), data[2].GetDetails(), data[2].GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, description, details, status, created_at, modified_at, vers FROM contact_system").ExpectQuery()
		expQuery.WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx)
		if err != nil {
			t.Errorf("Failed to read all contact systems: %v", err)
		}

		if res == nil {
			t.Errorf("Expect contact systems is not nil")
		}

		if len(res) < 3 {
			t.Errorf("Expect there are contact systems retrieved")
		}

		if res[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res[0].GetContactSystemCode())
		}

		if res[0].GetDescription() != input.GetDescription() {
			t.Errorf("Expect description %s, but got %s", input.GetDescription(), res[0].GetDescription())
		}

		if res[0].GetDetails() != input.GetDetails() {
			t.Errorf("Expect details %s, but got %s", input.GetDetails(), res[0].GetDetails())
		}

		if res[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), res[0].GetStatus())
		}
	}
}

func doSaveNewFailContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_system").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow, tmNow).WillReturnError(fmt.Errorf("DoInsert contact system failed"))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doSaveNewContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_system").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact system: %v", err)
		}
	}
}

func doSaveExistingFailContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_system").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow).WillReturnError(fmt.Errorf("DoUpdate contact system failed"))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doSaveExistingContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_system").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact system: %v", err)
		}
	}
}

func doDeleteFailContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		expCommMethodQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expContactQuery := mock.ExpectPrepare("SELECT 1 FROM contact").ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expQuery := mock.ExpectPrepare("DELETE FROM contact_system").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnError(fmt.Errorf("Delete contact system failed"))

		err := repo.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a Unknown error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteUnexistingContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		expCommMethodQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expContactQuery := mock.ExpectPrepare("SELECT 1 FROM contact").ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expQuery := mock.ExpectPrepare("DELETE FROM contact_system").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteAnyCommunicationMethodReference(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow("1")

		expQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		err := repo.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteFailAnyCommunicationMethodReference(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow("1").RowError(0, fmt.Errorf("AnyReference communication method row error"))

		expQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		err := repo.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteFailAnyContactReference(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow("1").RowError(0, fmt.Errorf("AnyReference contact row error"))

		expCommMethodQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expContactQuery := mock.ExpectPrepare("SELECT 1 FROM contact").ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		err := repo.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteAnyContactReference(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow("1")

		expCommMethodQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expContactQuery := mock.ExpectPrepare("SELECT 1 FROM contact").ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		err := repo.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.Unknown {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		} else {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doDeleteExistingContactSystem(ctx context.Context, input *contactsystem.ContactSystem) func(t *testing.T) {
	return func(t *testing.T) {
		expCommMethodQuery := mock.ExpectPrepare("SELECT 1 FROM communication_method").ExpectQuery()
		expCommMethodQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expContactQuery := mock.ExpectPrepare("SELECT 1 FROM contact").ExpectQuery()
		expContactQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		expQuery := mock.ExpectPrepare("DELETE FROM contact_system").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode())
		if err != nil {
			t.Errorf("Failed to delete contact system: %v", err)
		}
	}
}
