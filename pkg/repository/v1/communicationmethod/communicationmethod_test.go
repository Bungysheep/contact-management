package communicationmethod

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bungysheep/contact-management/pkg/api/v1/audit"
	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	repo ICommunicationMethodRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*communicationmethod.CommunicationMethod
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewCommunicationMethodRepository(db)

	data = append(data, &communicationmethod.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "EMAIL",
		Description:             "Email",
		Details:                 "Email",
		Status:                  "A",
		FormatField:             "[EMAIL_ADDRESS]",
	}, &communicationmethod.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "MOBILE",
		Description:             "Mobile",
		Details:                 "Mobile",
		Status:                  "A",
		FormatField:             "[MOBILE_NO]",
	}, &communicationmethod.CommunicationMethod{
		ContactSystemCode:       "CNTSYS001",
		CommunicationMethodCode: "FAX",
		Description:             "Fax",
		Details:                 "Fax",
		Status:                  "A",
		FormatField:             "[FAX_NO]",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodRepository(t *testing.T) {
	t.Run("DoRead Communication Method", doRead(ctx))

	t.Run("DoReadAll Communication Method", doReadAll(ctx))

	t.Run("DoSave Communication Method", doSave(ctx))

	t.Run("DoDelete Communication Method", doDelete(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailCommunicationMethod(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingCommunicationMethod(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorCommunicationMethod(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingCommunicationMethod(ctx, data[0]))
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail", doReadAllFailCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll unexisting", doReadAllUnexistingCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll row error", doReadAllRowErrorCommunicationMethods(ctx, data[0]))

		t.Run("DoReadAll existing", doReadAllExistingCommunicationMethods(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailCommunicationMethod(ctx, data[0]))

		t.Run("DoSave new", doSaveNewCommunicationMethod(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailCommunicationMethod(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingCommunicationMethod(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailCommunicationMethod(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingCommunicationMethod(ctx, data[0]))

		t.Run("DoDelete existing", doDeleteExistingCommunicationMethod(ctx, data[0]))
	}
}

func doReadFailCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("DoRead communication method failed"))

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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
			t.Errorf("Expect communication method is nil")
		}
	}
}

func doReadUnexistingCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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
			t.Errorf("Expect communication method is nil")
		}
	}
}

func doReadRowErrorCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoRead row error"))

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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
			t.Errorf("Expect communication method is nil")
		}
	}
}

func doReadExistingCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Errorf("Failed to read communication method: %v", err)
		}

		if res == nil {
			t.Errorf("Expect communication method is not nil")
		}

		if res.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res.GetContactSystemCode())
		}

		if res.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res.GetCommunicationMethodCode())
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

		if res.GetFormatField() != input.GetFormatField() {
			t.Errorf("Expect format field %s, but got %s", input.GetFormatField(), res.GetFormatField())
		}
	}
}

func doReadAllFailCommunicationMethods(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode()).WillReturnError(fmt.Errorf("DoReadAll communication method failed"))

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

func doReadAllUnexistingCommunicationMethods(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
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

func doReadAllRowErrorCommunicationMethods(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetDescription(), data[0].GetDetails(), data[0].GetStatus(), data[0].GetFormatField(), tmNow, tmNow, 1).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetDescription(), data[1].GetDetails(), data[1].GetStatus(), data[1].GetFormatField(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetDescription(), data[2].GetDetails(), data[2].GetStatus(), data[2].GetFormatField(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
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

func doReadAllExistingCommunicationMethods(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "description", "details", "status", "format_field", "created_at", "modified_at", "vers"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetDescription(), data[0].GetDetails(), data[0].GetStatus(), data[0].GetFormatField(), tmNow, tmNow, 1).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetDescription(), data[1].GetDetails(), data[1].GetStatus(), data[1].GetFormatField(), tmNow, tmNow, 1).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetDescription(), data[2].GetDetails(), data[2].GetStatus(), data[2].GetFormatField(), tmNow, tmNow, 1)

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, description, details, status, format_field, created_at, modified_at, vers FROM communication_method").ExpectQuery()
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

		if res[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res[0].GetCommunicationMethodCode())
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

		if res[0].GetFormatField() != input.GetFormatField() {
			t.Errorf("Expect format field %s, but got %s", input.GetFormatField(), res[0].GetFormatField())
		}
	}
}

func doSaveNewFailCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow, tmNow).WillReturnError(fmt.Errorf("DoInsert communication method failed"))

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

func doSaveNewCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       1,
		}

		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow, tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Errorf("Failed to save communication method: %v", err)
		}
	}
}

func doSaveExistingFailCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow).WillReturnError(fmt.Errorf("DoUpdate communication method failed"))

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

func doSaveExistingCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)
		tmstpNow, _ := ptypes.TimestampProto(tmNow)

		input.Audit = &audit.Audit{
			CreatedAt:  tmstpNow,
			ModifiedAt: tmstpNow,
			Vers:       2,
		}

		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetDescription(), input.GetDetails(), input.GetStatus(), input.GetFormatField(), tmNow).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Errorf("Failed to save communication method: %v", err)
		}
	}
}

func doDeleteFailCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("Delete communication method failed"))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doDeleteUnexistingCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doDeleteExistingCommunicationMethod(ctx context.Context, input *communicationmethod.CommunicationMethod) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Errorf("Failed to delete communication method: %v", err)
		}
	}
}
