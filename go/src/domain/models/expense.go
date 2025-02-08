package models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	id             uuid.UUID
	title          string `validate:"required"`
	amount         int
	purchasedAt    time.Time `validate:"required"`
	settledAt      *time.Time
	expenseTypeID  uuid.UUID `validate:"required"`
	groupID        uuid.UUID `validate:"required"`
	purchaseUserID uuid.UUID `validate:"required"`
}
