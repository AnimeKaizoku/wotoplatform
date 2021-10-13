package interfaces

type OngoingBattleBase interface {
	IsTooLate() bool
	IsLastRound() bool
	IsTooLateOrTooSoon() bool
	GetBattleSession() string
}
