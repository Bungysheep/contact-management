package message

import "fmt"

// Unimplemented message
func Unimplemented(entityName string) string {
	return fmt.Sprintf("%s has not been implemented.", entityName)
}

// FailedConnectToDatabase message
func FailedConnectToDatabase(err error) string {
	return fmt.Sprintf("Failed connect to database: %v", err)
}

// FailedPrepareRead message
func FailedPrepareRead(entityName string, err error) string {
	return fmt.Sprintf("Failed to prepare read %s: %v", entityName, err)
}

// FailedRead message
func FailedRead(name string, err error) string {
	return fmt.Sprintf("Failed to read %s: %v", name, err)
}

// FailedRetrieveRow message
func FailedRetrieveRow(name string, err error) string {
	return fmt.Sprintf("Failed to retrieve row of %s: %v", name, err)
}

// FailedRetrieveValues message
func FailedRetrieveValues(name string, err error) string {
	return fmt.Sprintf("Failed to retrieve row of %s: %v", name, err)
}

// DoesNotExist message
func DoesNotExist(name string) string {
	return fmt.Sprintf("%s does not exist", name)
}

// FailedPrepareInsert message
func FailedPrepareInsert(name string, err error) string {
	return fmt.Sprintf("Failed to prepare insert %s: %v", name, err)
}

// FailedInsert message
func FailedInsert(name string, err error) string {
	return fmt.Sprintf("Failed to insert %s: %v", name, err)
}

// NoRowInserted message
func NoRowInserted() string {
	return fmt.Sprintf("No row inseted")
}

// FailedPrepareUpdate message
func FailedPrepareUpdate(name string, err error) string {
	return fmt.Sprintf("Failed to prepare update %s: %v", name, err)
}

// FailedUpdate message
func FailedUpdate(name string, err error) string {
	return fmt.Sprintf("Failed to update %s: %v", name, err)
}

// FailedPrepareDelete message
func FailedPrepareDelete(name string, err error) string {
	return fmt.Sprintf("Failed to prepare delete %s: %v", name, err)
}

// FailedDelete message
func FailedDelete(name string, err error) string {
	return fmt.Sprintf("Failed to delete %s: %v", name, err)
}

// FailedRetrieveRowDeleted message
func FailedRetrieveRowDeleted(err error) string {
	return fmt.Sprintf("Failed to retrieve number of row deleted: %v", err)
}

// FailedDeleteAsReferenceExist message
func FailedDeleteAsReferenceExist(name string) string {
	return fmt.Sprintf("Failed to delete as there is any reference to %s", name)
}

// UnableDeleteDefault message
func UnableDeleteDefault(name string) string {
	return fmt.Sprintf("Unable to delete default %s", name)
}
