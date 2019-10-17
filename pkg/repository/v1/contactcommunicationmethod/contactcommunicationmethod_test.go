package contactcommunicationmethod

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	repo IContactCommunicationMethodRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contactcommunicationmethod.ContactCommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewContactCommunicationMethodRepository(db)

	data = append(data, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:            "CNTSYS001",
		ContactId:                    1,
		ContactCommunicationMethodId: 1,
		CommunicationMethodCode:      "EMAIL",
	}, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:            "CNTSYS001",
		ContactId:                    1,
		ContactCommunicationMethodId: 2,
		CommunicationMethodCode:      "MOBILE",
	}, &contactcommunicationmethod.ContactCommunicationMethod{
		ContactSystemCode:            "CNTSYS001",
		ContactId:                    1,
		ContactCommunicationMethodId: 3,
		CommunicationMethodCode:      "FAX",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactCommunicationMethodRepository(t *testing.T) {
	t.Run("DoRead Contact Communication Method", doRead(ctx))

	t.Run("DoReadAll Contact Communication Method", doReadAll(ctx))

	t.Run("DoSave Contact Communication Method", doSave(ctx))

	t.Run("DoDelete Contact Communication Method", doDelete(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingContactCommunicationMethod(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorContactCommunicationMethod(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingContactCommunicationMethod(ctx, data[0]))
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail", doReadAllFailContactCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll unexisting", doReadAllUnexistingContactCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll row error", doReadAllRowErrorContactCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll existing", doReadAllExistingContactCommunicationMethods(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoSave new", doSaveNewContactCommunicationMethod(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingContactCommunicationMethod(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailContactCommunicationMethod(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingContactCommunicationMethod(ctx, data[0]))

		t.Run("DoDelete existing", doDeleteExistingContactCommunicationMethod(ctx, data[0]))
	}
}

func doReadFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).WillReturnError(fmt.Errorf("DoRead contact communication method failed"))

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId())
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
			t.Errorf("Expect contact communication method is nil")
		}
	}
}

func doReadUnexistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId())
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
			t.Errorf("Expect contact communication method is nil")
		}
	}
}

func doReadRowErrorContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId(), input.GetCommunicationMethodCode(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoRead row error"))

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId())
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
			t.Errorf("Expect contact communication method is nil")
		}
	}
}

func doReadExistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId(), input.GetCommunicationMethodCode(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId())
		if err != nil {
			t.Errorf("Failed to read contact communication method: %v", err)
		}

		if res == nil {
			t.Errorf("Expect contact communication method is not nil")
		}

		if res.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact contact system code %s, but got %s", input.GetContactSystemCode(), res.GetContactSystemCode())
		}

		if res.GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), res.GetContactId())
		}

		if res.GetContactCommunicationMethodId() != input.GetContactCommunicationMethodId() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodId(), res.GetContactCommunicationMethodId())
		}

		if res.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res.GetCommunicationMethodCode())
		}
	}
}

func doReadAllFailContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnError(fmt.Errorf("DoReadAll contact communication method failed"))

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactId())
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

func doReadAllUnexistingContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactId())
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
			t.Errorf("Expect contact communication methods is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect no contact communication methods retrieved")
		}
	}
}

func doReadAllRowErrorContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactId(), data[0].GetContactCommunicationMethodId(), data[0].GetCommunicationMethodCode(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactId(), data[1].GetContactCommunicationMethodId(), data[1].GetCommunicationMethodCode(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactId(), data[2].GetContactCommunicationMethodId(), data[2].GetCommunicationMethodCode(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactId())
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

func doReadAllExistingContactCommunicationMethods(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "contact_communication_method_id", "communication_method_code", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactId(), data[0].GetContactCommunicationMethodId(), data[0].GetCommunicationMethodCode(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactId(), data[1].GetContactCommunicationMethodId(), data[1].GetCommunicationMethodCode(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactId(), data[2].GetContactCommunicationMethodId(), data[2].GetCommunicationMethodCode(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, contact_communication_method_id, communication_method_code, created_at, modified_at, vers FROM contact_communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetContactId())
		if err != nil {
			t.Errorf("Failed to read all contact communication methods: %v", err)
		}

		if res == nil {
			t.Errorf("Expect contact communication methods is not nil")
		}

		if len(res) < 3 {
			t.Errorf("Expect there are contact communication methods retrieved")
		}

		if res[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res[0].GetContactSystemCode())
		}

		if res[0].GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), res[0].GetContactId())
		}

		if res[0].GetContactCommunicationMethodId() != input.GetContactCommunicationMethodId() {
			t.Errorf("Expect contact communication method id %d, but got %d", input.GetContactCommunicationMethodId(), res[0].GetContactCommunicationMethodId())
		}

		if res[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res[0].GetCommunicationMethodCode())
		}
	}
}

func doSaveNewFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId(), input.GetCommunicationMethodCode(), tmNow, tmNow).WillReturnError(fmt.Errorf("DoInsert contact communication method failed"))

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

func doSaveNewContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact_communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId(), input.GetCommunicationMethodCode(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact communication method: %v", err)
		}
	}
}

func doSaveExistingFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId(), input.GetCommunicationMethodCode(), tmNow).WillReturnError(fmt.Errorf("DoUpdate contact communication method failed"))

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

func doSaveExistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact_communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId(), input.GetCommunicationMethodCode(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact communication method: %v", err)
		}
	}
}

func doDeleteFailContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).WillReturnError(fmt.Errorf("Delete contact communication method failed"))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId())
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

func doDeleteUnexistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId())
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

func doDeleteExistingContactCommunicationMethod(ctx context.Context, input *contactcommunicationmethod.ContactCommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact_communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId(), input.GetContactCommunicationMethodId())
		if err != nil {
			t.Errorf("Failed to delete contact communication method: %v", err)
		}
	}
}
