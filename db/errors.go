package db

import "fmt"

type InvalidDatabaseTypeError struct {
	suppliedType string
}

func (e InvalidDatabaseTypeError) Error() string {
	return fmt.Sprintf("invalid database type: %s", e.suppliedType)
}
