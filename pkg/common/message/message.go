package message

import "fmt"

// Unimplemented message
func Unimplemented(entityName string) string {
	return fmt.Sprintf("%s has not been implemented.", entityName)
}
