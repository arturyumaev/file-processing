package models

import (
	"github.com/google/uuid"
)

type Status string

const (
	StatusRecieved   Status = "recieved"
	StatusInQueue    Status = "in_queue"
	StatusProcessing Status = "processing"
	StatusDone       Status = "done"
	StatusError      Status = "error"
)

type FileInfo struct {
	Id        uuid.UUID `json:"uuid"`
	Hash      string    `json:"hash"`
	Status    Status    `json:"status"`
	TimeStamp int64     `json:"timestamp"`
}
