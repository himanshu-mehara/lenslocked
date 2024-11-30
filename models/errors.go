package models

import "errors"

var (
	ErrNotFound = errors.New("models: no resource could be found")
	ErrEmailTaken = errors.New("models: email address is already in use")
)
