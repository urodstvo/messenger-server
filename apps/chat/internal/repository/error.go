package repository

import "fmt"

func NewRepositoryError(err error) error {
	return fmt.Errorf("repository error: %v", err.Error())
}
