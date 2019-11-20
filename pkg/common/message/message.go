package message

import (
	"github.com/bungysheep/contact-management/pkg/common/constant/messagetype"
	"github.com/bungysheep/contact-management/pkg/models/v1/message"
)

// DoesNotExist message
func DoesNotExist(name string) message.IMessage {
	return message.NewMessage("CNT0001", messagetype.Error, "%s does not exist", name)
}

// Unimplemented message
func Unimplemented(entityName string) message.IMessage {
	return message.NewMessage("CNT0002", messagetype.Error, "%s has not been implemented.", entityName)
}

// FailedConnectToDatabase message
func FailedConnectToDatabase(err error) message.IMessage {
	return message.NewMessage("CNT0003", messagetype.Error, "Failed connect to database: %v", err)
}

// FailedPrepareRead message
func FailedPrepareRead(entityName string, err error) message.IMessage {
	return message.NewMessage("CNT0004", messagetype.Error, "Failed to prepare read %s: %v", entityName, err)
}

// FailedRead message
func FailedRead(name string, err error) message.IMessage {
	return message.NewMessage("CNT0005", messagetype.Error, "Failed to read %s: %v", name, err)
}

// FailedRetrieveRow message
func FailedRetrieveRow(name string, err error) message.IMessage {
	return message.NewMessage("CNT0006", messagetype.Error, "Failed to retrieve row of %s: %v", name, err)
}

// FailedRetrieveValues message
func FailedRetrieveValues(name string, err error) message.IMessage {
	return message.NewMessage("CNT0007", messagetype.Error, "Failed to retrieve row of %s: %v", name, err)
}

// FailedPrepareInsert message
func FailedPrepareInsert(name string, err error) message.IMessage {
	return message.NewMessage("CNT0008", messagetype.Error, "Failed to prepare insert %s: %v", name, err)
}

// FailedInsert message
func FailedInsert(name string, err error) message.IMessage {
	return message.NewMessage("CNT0009", messagetype.Error, "Failed to insert %s: %v", name, err)
}

// NoRowInserted message
func NoRowInserted() message.IMessage {
	return message.NewMessage("CNT0010", messagetype.Error, "No row inseted")
}

// FailedPrepareUpdate message
func FailedPrepareUpdate(name string, err error) message.IMessage {
	return message.NewMessage("CNT0011", messagetype.Error, "Failed to prepare update %s: %v", name, err)
}

// FailedUpdate message
func FailedUpdate(name string, err error) message.IMessage {
	return message.NewMessage("CNT0012", messagetype.Error, "Failed to update %s: %v", name, err)
}

// FailedPrepareDelete message
func FailedPrepareDelete(name string, err error) message.IMessage {
	return message.NewMessage("CNT0013", messagetype.Error, "Failed to prepare delete %s: %v", name, err)
}

// FailedDelete message
func FailedDelete(name string, err error) message.IMessage {
	return message.NewMessage("CNT0014", messagetype.Error, "Failed to delete %s: %v", name, err)
}

// FailedRetrieveRowDeleted message
func FailedRetrieveRowDeleted(err error) message.IMessage {
	return message.NewMessage("CNT0015", messagetype.Error, "Failed to retrieve number of row deleted: %v", err)
}

// FailedDeleteAsReferenceExist message
func FailedDeleteAsReferenceExist(name string) message.IMessage {
	return message.NewMessage("CNT0016", messagetype.Error, "Failed to delete as there is any reference to %s", name)
}

// UnableDeleteDefault message
func UnableDeleteDefault(name string) message.IMessage {
	return message.NewMessage("CNT0017", messagetype.Error, "Unable to delete default %s", name)
}

// MustBeSpecified message
func MustBeSpecified(name string) message.IMessage {
	return message.NewMessage("CNT0018", messagetype.Error, "%s must be specified", name)
}

// CannotMoreThanNChars message
func CannotMoreThanNChars(name string, nbr string) message.IMessage {
	return message.NewMessage("CNT0019", messagetype.Error, "%s can not more than %s chars", name, nbr)
}

// NotValid message
func NotValid(name string, value string) message.IMessage {
	return message.NewMessage("CNT0020", messagetype.Error, "%s '%s' is not valid", name, value)
}
