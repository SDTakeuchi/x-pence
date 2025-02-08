package models

import "github.com/google/uuid"

type Group struct {
	id           uuid.UUID     `validate:"required"`
	expenseTypes []ExpenseType `validate:"required"`
	userAID      uuid.UUID     `validate:"required"`
	userBID      uuid.UUID     `validate:"required"`
}
