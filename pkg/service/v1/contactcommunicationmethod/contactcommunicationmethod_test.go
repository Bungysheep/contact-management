package contactcommunicationmethod

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethod"
)

var (
	ctx context.Context
	svc contactcommunicationmethod.ContactCommunicationMethodServiceServer
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	svc = NewContactCommunicationMethodService()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactCommunicationMethodService(t *testing.T) {
	t.Run("DoRead Contact Communication Method", doRead(ctx))

	t.Run("DoReadAll Contact Communication Method", doRead(ctx))

	t.Run("DoSave Contact Communication Method", doRead(ctx))

	t.Run("DoDelete Contact Communication Method", doRead(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoRead(ctx, &contactcommunicationmethod.DoReadRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoReadAll(ctx, &contactcommunicationmethod.DoReadAllRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoSave(ctx, &contactcommunicationmethod.DoSaveRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoDelete(ctx, &contactcommunicationmethod.DoDeleteRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}
