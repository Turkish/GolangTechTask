package domain


type Vote struct {
	Uuid string
	VoteableUuid string
	AnswerIndex  int64
}
