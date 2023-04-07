package file_info

type FileInfoStatus string

const (
	StatusRecieved   FileInfoStatus = "recieved"
	StatusInQueue    FileInfoStatus = "in_queue"
	StatusProcessing FileInfoStatus = "processing"
	StatusDone       FileInfoStatus = "done"
	StatusError      FileInfoStatus = "error"
)

type FileInfo struct {
	Id        int64          `json:"id" db:"id"`
	Filename  string         `json:"filename_hash" db:"filename"`
	Status    FileInfoStatus `json:"status" db:"status"`
	TimeStamp string         `json:"timestamp" db:"timestamp"`
}
