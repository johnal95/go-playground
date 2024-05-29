package model

type Example struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

var currentId = 0

func NewExample(message string) Example {
	currentId += 1
	return Example{
		Id:      currentId,
		Message: message,
	}
}
