package models

import "github.com/google/uuid"

type ExpenseType struct {
	id                      uuid.UUID
	name                    string
	shareCalcLogic          ShareCalcLogic
	shareCalcLogicParameter int
}
