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
	Id           uuid.UUID      `json:"id" db:"id"`
	FilenameHash string         `json:"filename_hash" db:"filename_hash"`
	Status       FileInfoStatus `json:"status" db:"status"`
	TimeStamp    string         `json:"timestamp" db:"timestamp"`
}
