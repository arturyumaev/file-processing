package models

import (
	"github.com/google/uuid"
)

type FileInfoStatus string

const (
	FileInfoStatusRecieved   FileInfoStatus = "recieved"
	FileInfoStatusInQueue    FileInfoStatus = "in_queue"
	FileInfoStatusProcessing FileInfoStatus = "processing"
	FileInfoStatusDone       FileInfoStatus = "done"
	FileInfoStatusError      FileInfoStatus = "error"
)

type FileInfo struct {
	Id        uuid.UUID      `json:"uuid"`
	Hash      string         `json:"hash"`
	Status    FileInfoStatus `json:"status"`
	TimeStamp int64          `json:"timestamp"`
}
