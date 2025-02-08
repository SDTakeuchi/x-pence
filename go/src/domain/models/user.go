package models

import (
	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	name     string `validate:"required"`
	email    string `validate:"required,email"`
	password string `validate:"required"`
}
