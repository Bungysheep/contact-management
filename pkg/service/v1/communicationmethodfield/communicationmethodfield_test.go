package communicationmethodfield

import (
	"context"
	"os"
	"testing"

	"github.com/bungysheep/contact-management/pkg/api/v1/communicationmethodfield"
)

var (
	ctx context.Context
	svc communicationmethodfield.CommunicationMethodFieldServiceServer
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	svc = NewCommunicationMethodFieldService()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCommunicationMethodFieldService(t *testing.T) {
	t.Run("DoRead Communication Method Field", doRead(ctx))

	t.Run("DoReadAll Communication Method Field", doRead(ctx))

	t.Run("DoSave Communication Method Field", doRead(ctx))

	t.Run("DoDelete Communication Method Field", doRead(ctx))
}

func doRead(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoRead(ctx, &communicationmethodfield.DoReadRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doReadAll(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoReadAll(ctx, &communicationmethodfield.DoReadAllRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doSave(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoSave(ctx, &communicationmethodfield.DoSaveRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}

func doDelete(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := svc.DoDelete(ctx, &communicationmethodfield.DoDeleteRequest{})
		if err == nil {
			t.Errorf("Should be failed due to unimplemented.")
		}
	}
}
