package contact

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/contact"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	repo IContactRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*contact.Contact
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewContactRepository(db)

	data = append(data, &contact.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactId:         1,
		FirstName:         "James",
		LastName:          "Embongbulan",
		Status:            "A",
	}, &contact.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactId:         2,
		FirstName:         "Rindi",
		LastName:          "Allorerung",
		Status:            "A",
	}, &contact.Contact{
		ContactSystemCode: "CNTSYS001",
		ContactId:         3,
		FirstName:         "Marvel",
		LastName:          "Embongbulan",
		Status:            "A",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactRepository(t *testing.T) {
	t.Run("DoRead Contact", doRead(ctx))

	t.Run("DoReadAll Contact", doReadAll(ctx))

	t.Run("DoSave Contact", doSave(ctx))

	t.Run("DoDelete Contact", doDelete(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailContact(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingContact(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorContact(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingContact(ctx, data[0]))
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail", doReadAllFailContacts(ctx, data[0]))

		t.Run("DoReadAll unexisting", doReadAllUnexistingContacts(ctx, data[0]))

		t.Run("DoReadAll row error", doReadAllRowErrorContacts(ctx, data[0]))

		t.Run("DoReadAll existing", doReadAllExistingContacts(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailContact(ctx, data[0]))

		t.Run("DoSave new", doSaveNewContact(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailContact(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingContact(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailContact(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingContact(ctx, data[0]))

		t.Run("DoDelete existing", doDeleteExistingContact(ctx, data[0]))
	}
}

func doReadFailContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnError(fmt.Errorf("DoRead contact failed"))

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId())
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
			t.Errorf("Expect contact is nil")
		}
	}
}

func doReadUnexistingContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId())
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
			t.Errorf("Expect contact is nil")
		}
	}
}

func doReadRowErrorContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactId(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoRead row error"))

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId())
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
			t.Errorf("Expect contact is nil")
		}
	}
}

func doReadExistingContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetContactId(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetContactId())
		if err != nil {
			t.Errorf("Failed to read contact: %v", err)
		}

		if res == nil {
			t.Errorf("Expect contact is not nil")
		}

		if res.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res.GetContactSystemCode())
		}

		if res.GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), res.GetContactId())
		}

		if res.GetFirstName() != input.GetFirstName() {
			t.Errorf("Expect firstname %s, but got %s", input.GetFirstName(), res.GetFirstName())
		}

		if res.GetLastName() != input.GetLastName() {
			t.Errorf("Expect lastname %s, but got %s", input.GetLastName(), res.GetLastName())
		}

		if res.GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), res.GetStatus())
		}
	}
}

func doReadAllFailContacts(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnError(fmt.Errorf("DoReadAll contact failed"))

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode())
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

func doReadAllUnexistingContacts(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode())
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
			t.Errorf("Expect communication methods is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect no communication methods retrieved")
		}
	}
}

func doReadAllRowErrorContacts(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactId(), data[0].GetFirstName(), data[0].GetLastName(), data[0].GetStatus(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactId(), data[1].GetFirstName(), data[1].GetLastName(), data[1].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactId(), data[2].GetFirstName(), data[2].GetLastName(), data[2].GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode())
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

func doReadAllExistingContacts(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "contact_id", "first_name", "last_name", "status", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetContactId(), data[0].GetFirstName(), data[0].GetLastName(), data[0].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetContactId(), data[1].GetFirstName(), data[1].GetLastName(), data[1].GetStatus(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetContactId(), data[2].GetFirstName(), data[2].GetLastName(), data[2].GetStatus(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, contact_id, first_name, last_name, status, created_at, modified_at, vers FROM contact").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode())
		if err != nil {
			t.Errorf("Failed to read all communication methods: %v", err)
		}

		if res == nil {
			t.Errorf("Expect communication methods is not nil")
		}

		if len(res) < 3 {
			t.Errorf("Expect there are communication methods retrieved")
		}

		if res[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res[0].GetContactSystemCode())
		}

		if res[0].GetContactId() != input.GetContactId() {
			t.Errorf("Expect contact id %d, but got %d", input.GetContactId(), res[0].GetContactId())
		}

		if res[0].GetFirstName() != input.GetFirstName() {
			t.Errorf("Expect firstname %s, but got %s", input.GetFirstName(), res[0].GetFirstName())
		}

		if res[0].GetLastName() != input.GetLastName() {
			t.Errorf("Expect lastname %s, but got %s", input.GetLastName(), res[0].GetLastName())
		}

		if res[0].GetStatus() != input.GetStatus() {
			t.Errorf("Expect status %s, but got %s", input.GetStatus(), res[0].GetStatus())
		}
	}
}

func doSaveNewFailContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow, tmNow).WillReturnError(fmt.Errorf("DoInsert contact failed"))

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

func doSaveNewContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO contact").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact: %v", err)
		}
	}
}

func doSaveExistingFailContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow).WillReturnError(fmt.Errorf("DoUpdate contact failed"))

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

func doSaveExistingContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE contact").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId(), input.GetFirstName(), input.GetLastName(), input.GetStatus(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Errorf("Failed to save contact: %v", err)
		}
	}
}

func doDeleteFailContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnError(fmt.Errorf("Delete contact failed"))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId())
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

func doDeleteUnexistingContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId())
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

func doDeleteExistingContact(ctx context.Context, input *contact.Contact) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM contact").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetContactId()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetContactId())
		if err != nil {
			t.Errorf("Failed to delete contact: %v", err)
		}
	}
}
