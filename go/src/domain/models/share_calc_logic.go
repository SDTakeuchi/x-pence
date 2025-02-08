package models

import "github.com/google/uuid"

type ShareCalcLogic struct {
	name  string
	logic shareCalcLogic
}

func (l *ShareCalcLogic) Calculate(purchaseAmount int, userA, userB User, logicParameter int) int {
	return l.logic.calculate(purchaseAmount, userA, userB, logicParameter)
}

type shareCalcLogic interface {
	calculate(purchaseAmount int, userA, userB User, logicParameter int) int
}

func payUpTo(purchaseAmount int, userA, userB User, userAMaxShare int) map[uuid.UUID]int {
	if purchaseAmount <= userAMaxShare {
		return map[uuid.UUID]int{
			userA.id: purchaseAmount,
			userB.id: 0,
		}
	}

	return map[uuid.UUID]int{
		userA.id: userAMaxShare,
		userB.id: purchaseAmount - userAMaxShare,
	}
}

func byPercentage(purchaseAmount int, userA, userB User, userAPercentage int) map[uuid.UUID]int {
	userAAmount := purchaseAmount * userAPercentage / 100
	userBAmount := purchaseAmount - userAAmount

	return map[uuid.UUID]int{
		userA.id: userAAmount,
		userB.id: userBAmount,
	}
}

var (
	ShareCalcLogicPayUpTo = ShareCalcLogic{
		name:  "pay_up_to",
		logic: payUpTo,
	}
	ShareCalcLogicByPercentage = ShareCalcLogic{
		name:  "by_percentage",
		logic: byPercentage,
	}
)
