package communicationmethodlabel

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodlabel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx  context.Context
	repo ICommunicationMethodLabelRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*communicationmethodlabel.CommunicationMethodLabel
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewCommunicationMethodLabelRepository(db)

	data = append(data, &communicationmethodlabel.CommunicationMethodLabel{
		ContactSystemCode:            "CNTSYS001",
		CommunicationMethodCode:      "EMAIL",
		CommunicationMethodLabelCode: "HOME",
		Caption:                      "Home",
	}, &communicationmethodlabel.CommunicationMethodLabel{
		ContactSystemCode:            "CNTSYS001",
		CommunicationMethodCode:      "MOBILE",
		CommunicationMethodLabelCode: "WORK",
		Caption:                      "Work",
	}, &communicationmethodlabel.CommunicationMethodLabel{
		ContactSystemCode:            "CNTSYS001",
		CommunicationMethodCode:      "FAX",
		CommunicationMethodLabelCode: "SCHOOL",
		Caption:                      "School",
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodLabelRepository(t *testing.T) {
	t.Run("DoRead Communication Method Label", doRead(ctx))

	t.Run("DoReadAll Communication Method Label", doReadAll(ctx))

	t.Run("DoSave Communication Method Label", doSave(ctx))

	t.Run("DoDelete Communication Method Label", doDelete(ctx))

	t.Run("DoDeleteAll Communication Method Label", doDeleteAll(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoRead fail", doReadFailCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoRead unexisting", doReadUnexistingCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoRead row error", doReadRowErrorCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoRead existing", doReadExistingCommunicationMethodLabel(ctx, data[0]))
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoReadAll fail", doReadAllFailCommunicationMethodLabels(ctx, data[0]))

		t.Run("DoReadAll unexisting", doReadAllUnexistingCommunicationMethodLabels(ctx, data[0]))

		t.Run("DoReadAll row error", doReadAllRowErrorCommunicationMethodLabels(ctx, data[0]))

		t.Run("DoReadAll existing", doReadAllExistingCommunicationMethodLabels(ctx, data[0]))
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoSave new fail", doSaveNewFailCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoSave new", doSaveNewCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoSave existing fail", doSaveExistingFailCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoSave existing", doSaveExistingCommunicationMethodLabel(ctx, data[0]))
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDelete fail", doDeleteFailCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoDelete unexisting", doDeleteUnexistingCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoDelete existing", doDeleteExistingCommunicationMethodLabel(ctx, data[0]))
	}
}

func doDeleteAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("DoDeleteAll fail", doDeleteAllFailCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoDeleteAll unexisting", doDeleteAllUnexistingCommunicationMethodLabel(ctx, data[0]))

		t.Run("DoDeleteAll existing", doDeleteAllExistingCommunicationMethodLabel(ctx, data[0]))
	}
}

func doReadFailCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnError(fmt.Errorf("DoRead communication method label failed"))

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
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
			t.Errorf("Expect communication method label is nil")
		}
	}
}

func doReadUnexistingCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
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
			t.Errorf("Expect communication method label is nil")
		}
	}
}

func doReadRowErrorCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).
			RowError(0, fmt.Errorf("DoRead row error"))

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
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
			t.Errorf("Expect communication method label is nil")
		}
	}
}

func doReadExistingCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption())

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnRows(rows)

		res, err := repo.DoRead(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
		if err != nil {
			t.Errorf("Failed to read communication method label: %v", err)
		}

		if res == nil {
			t.Errorf("Expect communication method label is not nil")
		}

		if res.GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res.GetContactSystemCode())
		}

		if res.GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res.GetCommunicationMethodCode())
		}

		if res.GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), res.GetCommunicationMethodLabelCode())
		}

		if res.GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), res.GetCaption())
		}
	}
}

func doReadAllFailCommunicationMethodLabels(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("DoReadAll communication method label failed"))

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doReadAllUnexistingCommunicationMethodLabels(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"})

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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
			t.Errorf("Expect communication method labels is not nil")
		}

		if len(res) != 0 {
			t.Errorf("Expect no communication method labels retrieved")
		}
	}
}

func doReadAllRowErrorCommunicationMethodLabels(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetCommunicationMethodLabelCode(), data[0].GetCaption()).
			RowError(0, fmt.Errorf("DoReadAll row error")).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetCommunicationMethodLabelCode(), data[1].GetCaption()).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetCommunicationMethodLabelCode(), data[2].GetCaption())

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doReadAllExistingCommunicationMethodLabels(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"contact_system_code", "communication_method_code", "communication_method_label_code", "caption"}).
			AddRow(data[0].GetContactSystemCode(), data[0].GetCommunicationMethodCode(), data[0].GetCommunicationMethodLabelCode(), data[0].GetCaption()).
			AddRow(data[1].GetContactSystemCode(), data[1].GetCommunicationMethodCode(), data[1].GetCommunicationMethodLabelCode(), data[1].GetCaption()).
			AddRow(data[2].GetContactSystemCode(), data[2].GetCommunicationMethodCode(), data[2].GetCommunicationMethodLabelCode(), data[2].GetCaption())

		expQuery := mock.ExpectPrepare("SELECT contact_system_code, communication_method_code, communication_method_label_code, caption FROM communication_method_label").ExpectQuery()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnRows(rows)

		res, err := repo.DoReadAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Errorf("Failed to read all communication method labels: %v", err)
		}

		if res == nil {
			t.Errorf("Expect communication method labels is not nil")
		}

		if len(res) < 3 {
			t.Errorf("Expect there are communication method labels retrieved")
		}

		if res[0].GetContactSystemCode() != input.GetContactSystemCode() {
			t.Errorf("Expect contact system code %s, but got %s", input.GetContactSystemCode(), res[0].GetContactSystemCode())
		}

		if res[0].GetCommunicationMethodCode() != input.GetCommunicationMethodCode() {
			t.Errorf("Expect communication method code %s, but got %s", input.GetCommunicationMethodCode(), res[0].GetCommunicationMethodCode())
		}

		if res[0].GetCommunicationMethodLabelCode() != input.GetCommunicationMethodLabelCode() {
			t.Errorf("Expect communication method label code %s, but got %s", input.GetCommunicationMethodLabelCode(), res[0].GetCommunicationMethodLabelCode())
		}

		if res[0].GetCaption() != input.GetCaption() {
			t.Errorf("Expect caption %s, but got %s", input.GetCaption(), res[0].GetCaption())
		}
	}
}

func doSaveNewFailCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method_label").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).WillReturnError(fmt.Errorf("DoInsert communication method label failed"))

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

func doSaveNewCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expInsQuery := mock.ExpectPrepare("INSERT INTO communication_method_label").ExpectExec()
		expInsQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoInsert(ctx, input)
		if err != nil {
			t.Errorf("Failed to save communication method label: %v", err)
		}
	}
}

func doSaveExistingFailCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_label").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).WillReturnError(fmt.Errorf("DoUpdate communication method label failed"))

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

func doSaveExistingCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expUpdQuery := mock.ExpectPrepare("UPDATE communication_method_label").ExpectExec()
		expUpdQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode(), input.GetCaption()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoUpdate(ctx, input)
		if err != nil {
			t.Errorf("Failed to save communication method label: %v", err)
		}
	}
}

func doDeleteFailCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnError(fmt.Errorf("Delete communication method label failed"))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
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

func doDeleteUnexistingCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
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

func doDeleteExistingCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDelete(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode(), input.GetCommunicationMethodLabelCode())
		if err != nil {
			t.Errorf("Failed to delete communication method label: %v", err)
		}
	}
}

func doDeleteAllFailCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnError(fmt.Errorf("Delete all communication method labels failed"))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
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

func doDeleteAllUnexistingCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				if s.Code() != codes.NotFound {
					t.Fatalf("Expect a NotFound error, but got %s", s.Code())
				}
			}
		}
	}
}

func doDeleteAllExistingCommunicationMethodLabel(ctx context.Context, input *communicationmethodlabel.CommunicationMethodLabel) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare("DELETE FROM communication_method_label").ExpectExec()
		expQuery.WithArgs(input.GetContactSystemCode(), input.GetCommunicationMethodCode()).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DoDeleteAll(ctx, input.GetContactSystemCode(), input.GetCommunicationMethodCode())
		if err != nil {
			t.Errorf("Failed to delete all communication method labels: %v", err)
		}
	}
}
