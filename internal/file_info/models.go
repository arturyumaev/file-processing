package file_info

import "github.com/google/uuid"

type FileInfoStatus string

const (
	StatusRecieved   FileInfoStatus = "recieved"
	StatusInQueue    FileInfoStatus = "in_queue"
	StatusProcessing FileInfoStatus = "processing"
	StatusDone       FileInfoStatus = "done"
	StatusError      FileInfoStatus = "error"
)

type FileInfo struct {
	Id           uuid.UUID      `json:"id" db:"id"`
	FilenameHash string         `json:"filename_hash" db:"filename_hash"`
	Status       FileInfoStatus `json:"status" db:"status"`
	TimeStamp    string         `json:"timestamp" db:"timestamp"`
}
